package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	matrixSize = 250
)

var (
	matrixA      = [matrixSize][matrixSize]int{}
	matrixB      = [matrixSize][matrixSize]int{}
	result       = [matrixSize][matrixSize]int{}
	workStart    = NewBarrier(matrixSize + 1)
	workComplete = NewBarrier(matrixSize + 1)
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

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workOutRow(row int) {
	for {
		workStart.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
		workComplete.Wait()
	}
}

func main() {
	fmt.Println("Working...")
	for row := 0; row < matrixSize; row++ {
		go workOutRow(row)
	}

	start := time.Now()
	for i := 0; i < 100; i++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		workStart.Wait()
		workComplete.Wait()
	}
	elapsed := time.Since(start)

	fmt.Println("Done")
	fmt.Println("Processing took %s\n", elapsed)
}
