package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

// main функция нашей программы вызывает вспомогательную функцию MD5All,
// которая возвращает карту (map) имени пути к значению дайджеста, затем сортирует и печатает результаты:
func main() {
	start := time.Now()
	// Рассчитать MD5 сумму всех файлов
	// в указанном каталоге,
	// затем печатаем результаты,
	// отсортированные по имени пути.
	m, err := MD5All("/home/andrey/go/src")
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
	fmt.Println("For ", time.Since(start))
}

/*
// MD5All (последовательная реализация) читает все файлы в дереве файлов с корнем в root
// и возвращает карту пути к файлу к MD5 сумме
// содержимого файла. Если происходит сбой прохода
// по каталогу или сбой любой операции чтения,
// MD5All возвращает ошибку.
func MD5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		m[path] = md5.Sum(data)
		return nil
	})
if err != nil {
		return nil, err
	}
	return m, nil
}
*/

// Параллельное получение digest'а
// В parallel.go, мы разделили MD5All на двухступенчатый пайплайн.
// Первый этап, sumFiles, обходит дерево, получает digest каждого файла в новой go-процедуре
// и отправляет результаты по каналу с типом значения result:
type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

// sumFiles возвращает два канала: один для результатов и другой для ошибки, возвращаемой filepath.Walk.
// Функция walk запускает новую go-процедуру для обработки каждого обычного файла, а затем проверяет done.
// Если done закрыт, walk немедленно останавливается:
func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	// Для каждого обычного файла запускаем goroutine,
	// которая суммирует файл и отправляет
	// результат в c. Ошибки walk отправляются в errc.
	c := make(chan result)
	errc := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}:
				case <-done:
				}
				wg.Done()
			}()
			// Завершаем walk если done закрыт.
			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})
		// Walk вернулся,
		// поэтому все вызовы wg.Add завершены.
		// Начинаем goroutine для закрытия c,
		// как только все посылки сделаны.
		go func() {
			wg.Wait()
			close(c)
		}()
		// select не нужен здесь, поскольку errc буферизован.
		errc <- err
	}()
	return c, errc
}

// MD5All получает значения дайджеста от c.
// MD5All возвращается досрочно при ошибке, закрытие осуществляется с помощью defer:
func MD5All(root string) (map[string][md5.Size]byte, error) {
	// MD5All закрывает done канал при возврате;
	// это может быть сделано
	// до получения всех значений от c и errc.
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

/*
// Ограниченный параллелизм
// Реализация MD5All в parallel.go запускает новую goroutine для каждого файла. В каталоге со многими большими файлами это может аллоцировать больше памяти, чем доступно на машине.
// Мы можем ограничить эти аллокации, ограничив число файлов, читаемых параллельно. В bounded.go мы делаем это, создавая фиксированное количество go-процедур для чтения файлов. Наш пайплайн теперь имеет три этапа: пройтись по дереву, прочитать файлы и получить дайджесты файлов и собрать дайджесты.
// Первый этап, walkFiles, генерирует пути обычных файлов в дереве:
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
    paths := make(chan string)
    errc := make(chan error, 1)
    go func() {
        // Закрываем paths канал после возврата Walk.
        defer close(paths)
        // select не требуется для этой отправки,
        // поскольку errc буферизован.
        errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.Mode().IsRegular() {
                return nil
            }
            select {
            case paths <- path:
            case <-done:
                return errors.New("walk canceled")
            }
            return nil
        })
    }()
    return paths, errc
}

// Средняя стадия запускает фиксированное число digester go-процедур, которые получают имена файлов из путей
// и отправляют результаты по каналу c:
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
    for path := range paths {
        data, err := ioutil.ReadFile(path)
        select {
        case c <- result{path, md5.Sum(data), err}:
        case <-done:
            return
        }
    }
}

// В отличие от наших предыдущих примеров, digester не закрывает свой выходной канал,
// так как несколько go-процедур отправляют по общему каналу.
// Вместо этого код в MD5All организует закрытие канала после завершения работы всех digester:

// Запускаем фиксированное количество go-процедур для чтения и получения дайджеста файлов.
c := make(chan result)
var wg sync.WaitGroup
const numDigesters = 20
wg.Add(numDigesters)
for i := 0; i < numDigesters; i++ {
    go func() {
        digester(done, paths, c)
        wg.Done()
    }()
}
go func() {
    wg.Wait()
    close(c)
}()

// Вместо этого мы могли бы сделать так чтобы каждый digester создавал и возвращал свой собственный выходной канал,
// но тогда нам потребовались бы дополнительные go-процедуры для сдувания (fan-in) результатов.
//
// Последний этап получает все результаты от c, затем проверяет ошибку от errc.
// Эта проверка не может произойти раньше, так как до этого момента walkFiles может блокировать отправку значений вниз:
    m := make(map[string][md5.Size]byte)
    for r := range c {
        if r.err != nil {
            return nil, r.err
        }
        m[r.path] = r.sum
    }
    // Проверяем на пройзошел ли сбой Walk.
    if err := <-errc; err != nil {
        return nil, err
    }
    return m, nil
}

*/
