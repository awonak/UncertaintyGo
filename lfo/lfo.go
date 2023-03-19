package main

import (
	"log"
	"math"

	uncertainty "github.com/awonak/UncertaintyGo"
)

type LFO struct {
	output  *uncertainty.Output
	rate    int
	value   int
	voltage float32
	rising  bool
}

func NewLFO(output *uncertainty.Output, rate int) *LFO {
	if 0 > rate || rate > math.MaxInt16 {
		log.Fatalf("rate must be between 0 and 32768: %v", rate)

	}
	return &LFO{output, rate, 0, 0, true}
}

// Next will advance the LFO value within a 0 to 32768 according its rate and
// cv input resulting in a 0v to 5v cv output.
func (lfo *LFO) Next(nudge int) {
	// Increase or decrease the current LFO value.
	if lfo.rising {
		lfo.value += lfo.rate + nudge
	} else {
		lfo.value -= lfo.rate + nudge
	}

	// If the LFO value has exceeded the 0 to 32768 boundary, clamp and flip
	// the rising bit.
	if lfo.rising && lfo.value >= math.MaxInt16 {
		lfo.value = math.MaxInt16
		lfo.rising = false
	} else if !lfo.rising && lfo.value <= 0 {
		lfo.value = 0
		lfo.rising = true
	}

	// Calculate and set the output voltage.
	lfo.voltage = uncertainty.MaxVoltage * (float32(lfo.value) / float32(math.MaxInt16))
	lfo.output.Voltage(lfo.voltage)
}
