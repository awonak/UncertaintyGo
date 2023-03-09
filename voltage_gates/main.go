package main

import (
	"log"
	"time"

	uncertainty "github.com/awonak/UncertaintyGo"
)

// Enable to print serial monitoring log messages.
const Debug = false

func main() {
	// Main loop.
	for {
		// Read the cv input clipped to a 0-5v range.
		read := uncertainty.ReadCV()

		// Read the cv input and scale down voltage to select one of the 8 cv
		// ouputs to activate.
		activeGate := int(read / 4096)

		// Iterate over all of the cv outputs and set the "active" cv high,
		// otherwise set the cv low.
		for i, cv := range uncertainty.Outputs {
			if i == activeGate {
				cv.High()
			} else {
				cv.Low()
			}
		}

		if Debug {
			log.Printf("CVInput: %d\tactiveGate: %d\n\r", read, activeGate)
			time.Sleep(time.Millisecond * 10)
		}
	}
}
