package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches []string
	// it's another way to make syncronization
	// if creates a main thread and you'll add child threads on it
	// that thread will wait for the others to notify they are done
	waitgroup = sync.WaitGroup{}

	// lock to block the writing opperation in the matches array
	lock = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)

	files, _ := ioutil.ReadDir(root)

	// we are doing a recursive func over the files in each folder
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			// for the writing opperation
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			// add a child to the waitgroup
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}

	// notify finish for each directory search
	waitgroup.Done()
}

func main() {
	// here we add the main thread
	waitgroup.Add(1)
	go fileSearch("C:/tools", "README.md")

	// tell it to wait for the fileSearch to finish to print the matches
	waitgroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}
