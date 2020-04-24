package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"
)

// Рекурсивный обходи указанного в качестве первого параметра каталог и вывод в коноль имени размеров (только для файлов)
// Если задан второй параметр (любое слов), обработка каждого подкаталоге делается в отдельной горутине, пущенной в параллель
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", path.Base(os.Args[0]), "dir_to_list [in_parallel]")
		return
	}
	pd := os.Args[1]
	pda, _ := filepath.Abs(pd)
	start := time.Now()
	if len(os.Args) == 2 {
		fmt.Println("Walking the dir: ", pda)
		err := filepath.Walk(pda,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.Mode().IsDir() {
					fmt.Println(path)
				} else {
					fmt.Println(" ", path, info.Size(), "byte(s)")
				}
				return nil
			})
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Walking the dir in parallel: ", pda)
		files, _ := ioutil.ReadDir(pda)
		var wg sync.WaitGroup
		for _, f := range files {
			var entry func(p string, f os.FileInfo)
			entry = func(p string, f os.FileInfo) {
				wg.Add(1)
				if f.Mode().IsDir() {
					pd := filepath.Join(p, f.Name())
					pda, _ := filepath.Abs(pd)
					fmt.Println(pda)
					fd, _ := ioutil.ReadDir(pd)
					for _, i := range fd {
						entry(pd, i)
					}
				}
				fmt.Println(" ", f.Name(), f.Size(), "byte(s)")
				wg.Done()
			}
			entry(pda, f)
		}
		wg.Wait()
	}
	fmt.Println("Listed for ", time.Since(start))
}
