package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2D struct {
	x int
	y int
}

// we are supposing we have 8 nucles
const numberOfThreads int = 8

var (
	// regexp to match the points of the polygon
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)
	// to wait for them all to finish
	waitGroup = sync.WaitGroup{}
)

// using the shoelace algorithm
func findArea(inputChannel chan string) {
	for pointsStr := range inputChannel {
		var points []Point2D
		for _, p := range r.FindAllStringSubmatch(pointsStr, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{x, y})
		}

		area := 0.0
		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}
		fmt.Println(math.Abs(area) / 2.0)
	}
	waitGroup.Done()
}

func main() {
	dat, err := ioutil.ReadFile("./channels/thread-pool/polygons.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	text := string(dat)

	inputChannel := make(chan string, 1000)
	for i := 0; i < numberOfThreads; i++ {
		go findArea(inputChannel)
	}
	// between the 8 threads all the messages will be procesed
	waitGroup.Add(numberOfThreads)
	start := time.Now()

	for _, line := range strings.Split(text, "\n") {
		// sent the messages to the channel where the threads are hearing
		inputChannel <- line
	}
	// finish to add
	close(inputChannel)

	waitGroup.Wait()

	elapsed := time.Since(start)

	fmt.Printf("Processing took %s \n", elapsed)
}
