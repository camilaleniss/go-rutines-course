package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// two ways to apply syncronization
	lock   = sync.Mutex{}
	rwLock = sync.RWMutex{}
)

// what happen when we call this?
func RunAndWait() {
	go callLockTwice()
	time.Sleep(10 * time.Second)
}

// the second lock will block forever, print line never happens, app exit in 10 seconds
func callLockTwice() {
	lock.Lock()
	lock.Lock()
	fmt.Print("Hello there")
}

// what happen when we call this?
func StartThreadsA() {
	for i := 1; i <= 2; i++ {
		go oneTwoThreeA()
	}
	time.Sleep(1 * time.Second)
}

// since the lock is before the for loop, the response will be 1,2,3,1,2,3
func oneTwoThreeA() {
	lock.Lock()
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	lock.Unlock()
}

// what happen when we call this?
func StartThreadsB() {
	for i := 1; i <= 2; i++ {
		go oneTwoThreeB()
	}
	time.Sleep(1 * time.Second)
}

// this will not block the read operation, in fact they will print 1,2,3 in any order 2 times
func oneTwoThreeB() {
	rwLock.RLock()
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	rwLock.RLock()
}

func main() {
	// this is how you start a thread
	go StartThreadsA()

	time.Sleep(5 * time.Second)
}
