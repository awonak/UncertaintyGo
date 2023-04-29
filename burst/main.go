package main

import (
	"math"
	"sync"
	"time"

	uncertainty "github.com/awonak/UncertaintyGo"
)

type Gate rune

const (
	high Gate = iota
	low

	MaxWidth = 500 // in milliseconds
	MinWidth = 10
)

var (
	oldGate Gate
	gate    Gate
	level   int
)

func main() {
	// Use a waitgroup to ensure bursts are not retriggered before completing.
	var wg sync.WaitGroup

	// Main loop.
	gate, level = readInput()
	for {
		oldGate = gate
		gate, level = readInput()

		if gate == high && oldGate == low {
			width := MaxWidth - int(float32(level)/float32(math.MaxInt16)*MaxWidth)
			width = uncertainty.Clamp(width, MinWidth, MaxWidth)
			for i, cv := range uncertainty.Outputs {
				wg.Add(1)
				go func(cv *uncertainty.Output, count int, width time.Duration) {
					defer wg.Done()
					createBurst(cv, count, width)
				}(cv, i+1, time.Duration(width)*time.Millisecond)
			}
			wg.Wait()
		}
	}
}

func readInput() (Gate, int) {
	// Ensure a little time passes between reads.
	time.Sleep(10 * time.Millisecond)
	read := uncertainty.Read()
	// Check if input voltage > 1v.
	if read > 6553 {
		return high, read
	}
	return low, read
}

func createBurst(cv *uncertainty.Output, count int, width time.Duration) {
	for i := 0; i < count; i++ {
		cv.High()
		time.Sleep(width / 2)
		cv.Low()
		time.Sleep(width / 2)
	}
}
