package main

import (
	uncertainty "github.com/awonak/UncertaintyGo"
)

var (
	// rates define the amount each of the 8 LFO outputs should advance by.
	rates = [8]int{
		1 << 9,
		1 << 8,
		1 << 7,
		1 << 6,
		1 << 5,
		1 << 4,
		1 << 3,
		1 << 2,
	}

	// The collection of LFO state machines for each output.
	lfos [8]*LFO
)

func main() {
	// Initialize each LFO state machine.
	for i, rate := range rates {
		lfos[i] = NewLFO(uncertainty.Outputs[i], rate)
	}

	// Main loop.
	for {
		for i, lfo := range lfos {
			// Capture the cv input to increase the LFO speed.
			nudge := (uncertainty.Read() / 4)
			// Calculate next voltage value for this LFO and set the cv output.
			lfo.Next(nudge / (i + 1))
		}
	}
}
