package main

import (
	"log"
	"machine"
	"math"
	"time"
)

const (
	// GPIO mapping to Uncertainty panel.
	CVInput = machine.ADC0
	CV1     = machine.GPIO27
	CV2     = machine.GPIO28
	CV3     = machine.GPIO29
	CV4     = machine.GPIO0
	CV5     = machine.GPIO3
	CV6     = machine.GPIO4
	CV7     = machine.GPIO2
	CV8     = machine.GPIO1

	// Enable to print serial monitoring log messages.
	Debug = false
)

func main() {
	// Initialize the cv input GPIO as an analog input.
	machine.InitADC()
	cvInput := machine.ADC{Pin: CVInput}
	cvInput.Configure(machine.ADCConfig{})

	// Create an array of our cv outputs and configure for output.
	cvOutputs := [8]machine.Pin{CV1, CV2, CV3, CV4, CV5, CV6, CV7, CV8}
	for _, cv := range cvOutputs {
		cv.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	for {
		// Read the cv input clipped to a 0-5v range.
		read := int(cvInput.Get()) - math.MaxInt16
		if read < 0 {
			read = 0
		}

		// Read the cv input and scale down voltage to select one of the 8 cv ouputs to activate.
		activeGate := int(read / 4096)

		// Iterate over all of the cv outputs and set the "active" cv high, otherwise set the cv low.
		for i, cv := range cvOutputs {
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
