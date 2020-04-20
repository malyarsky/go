package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", path.Base(os.Args[0]), "dir_to_list")
		return
	}
	pd := os.Args[1]
	pda, _ := filepath.Abs(pd)
	fmt.Println(pda)
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
			fmt.Printf(" %s %#o %dbyte(s)\n", f.Name(), f.Mode().Perm(), f.Size())
			wg.Done()
		}
		entry(pda, f)
	}
	wg.Wait()
}
