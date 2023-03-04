package main

import (
	"machine"
	"math"
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

	// Number of times to read analog input for an average reading.
	ReadSamples = 500

	// Calibrated average min read uint16 voltage within a 0-5v range.
	MinCalibratedRead = 415

	// Calibrated average max read uint16 voltage within a 0-5v range.
	MaxCalibratedRead = 29582

	// Upper limit of voltage read by the cv input.
	MaxReadVoltage float64 = 5
)

var (
	// Create package global variables for the cv input and outputs.
	cvInput    machine.ADC
	cvOutputs  [8]machine.Pin
	pwmOutputs [8]PWM

	// We need a rather high frequency to achieve a stable cv ouput, which means we need a rather low duty cycle period.
	// Set a period of 500ns.
	defaultPeriod uint64 = 1e9 / 500
)

// PWM is the interface necessary for configuring a cv output for PWM.
type PWM interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Top() uint32
	Set(channel uint8, value uint32)
	SetPeriod(period uint64) error
}

// ReadCV will return the cv input scaled to 0v-5v as an int with 0 for 0v and 32768 for 5v.
func ReadCV() int {
	var sum int
	for i := 0; i < ReadSamples; i++ {
		read := int(cvInput.Get()) - math.MaxInt16
		if read < 0 {
			read = 0
		}
		sum += read
	}
	return sum / ReadSamples
}

// ReadVoltage will return the cv input scaled to 0v-5v as a float with 0.0 for 0v and 5.0 for 5v.
func ReadVoltage() float64 {
	read := ReadCV()
	return MaxReadVoltage * (float64(read-MinCalibratedRead) / float64(MaxCalibratedRead-MinCalibratedRead))
}
