package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errors := make(chan error, 1)
	go func() {
		defer close(paths)
		errors <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
			}
			return nil
		})
	}()

	return paths, errors
}

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

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

func MD5All(root string) (map[string][md5.Size]byte, error) {
	fmt.Printf("Root dir: %s\n", root)
	done := make(chan struct{})
	defer close(done)

	paths, errors := walkFiles(done, root)

	c := make(chan result)
	numOfWorkers := 10
	var wg sync.WaitGroup
	wg.Add(numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return m, nil
}

func main() {
	m, err := MD5All(os.Args[1])

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
		fmt.Printf("%x, %s\n", m[path], path)
	}
}
