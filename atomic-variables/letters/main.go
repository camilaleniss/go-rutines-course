package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

var lock = sync.Mutex{}

func countLetters(url string, frequency *[26]int32, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("unable to get url")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("unable to read response body")
	}

	for i := 0; i <= 20; i++ {
		for _, b := range body {
			c := strings.ToLower(string(b))
			//lock.Lock() for the example using simple mutex
			index := strings.Index(allLetters, c)
			if index >= 0 {
				//*frequency[c] += 1 for the example using simple mutex
				atomic.AddInt32(&frequency[index], 1)
			}
			//lock.Unlock() for the example using simple mutex
		}
	}
	wg.Done()
}

func main() {
	var frequency [26]int32
	wg := sync.WaitGroup{}

	// To process the requests
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s\n", elapsed)

	// To print the response
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}
