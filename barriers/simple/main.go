package main

import (
	"sync"
	"time"
)

type Barrier struct {
	total int
	count int
	mutex *sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(size int) *Barrier {
	lockToUse := &sync.Mutex{}
	condToUse := sync.NewCond(lockToUse)
	return &Barrier{size, size, lockToUse, condToUse}
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count -= 1
	if b.count == 0 {
		b.count = b.total
		// when they meet it broadcast
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.mutex.Unlock()
}

func waitOnBarrier(name string, timeToSleep int, barrier *Barrier) {
	for {
		println(name, "running")
		time.Sleep(time.Duration(timeToSleep) * time.Second)
		println(name, "is waiting on barrier")
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	// red run slower than blue
	go waitOnBarrier("red", 4, barrier)
	go waitOnBarrier("blue", 10, barrier)

	// time to run the script
	time.Sleep(time.Duration(100) * time.Second)
}
