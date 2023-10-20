package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		fmt.Println("Blue: Acquiring lock1")
		lock1.Lock()
		fmt.Println("Blue: Acquiring lock2")
		lock2.Lock()
		fmt.Println("Blue: Both locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: Locks Released")
	}
}

// at some point cause we tent to unlock the first one and then we are waiting for the two to be unlocked
// it waits forever
func redRobot() {
	for {
		fmt.Println("Red: Acquiring lock2")
		lock2.Lock()
		fmt.Println("Red: Acquiring lock1")
		lock1.Lock()
		fmt.Println("Red: Both locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Red: Locks Released")
	}
}

func main() {
	go redRobot()
	go blueRobot()
	time.Sleep(20 * time.Second)

	// they never hit done
	fmt.Println("Done")
}
