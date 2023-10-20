package main

import (
	"sync/atomic"
	"time"
)

var (
	money int32 = 100
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, 10)
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		// this method works as a mutex
		atomic.AddInt32(&money, -10)
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}

func main() {
	// since we are moving primitives, let's show how to use atomic methods
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
