package main

import (
	"log"
	"time"

	"tinygo.org/x/drivers/tone"
)

// Enable to print serial monitoring log messages.
const Debug = false

var (
	// User configurable variables for VCOs.
	scales [3]Scale
	roots  [3]tone.Note
	vcos   [3]VCO
)

func main() {
	if Debug {
		// Provide a brief pause to allow time to start up the serial monitor to capture errors.
		log.Print("START...")
		time.Sleep(time.Second * 5)
		log.Print("Ready...")
	}

	// Initialize a collection of PWM VCOs bound to a scale and root note starting at 0v.
	vcos = [3]VCO{
		NewVCO(pwmOutputs[1], cvOutputs[1], scales[0], roots[0]),
		NewVCO(pwmOutputs[3], cvOutputs[3], scales[1], roots[1]),
		NewVCO(pwmOutputs[5], cvOutputs[5], scales[2], roots[2]),
	}

	// Main program loop.
	for {
		voltage := ReadVoltage()
		for _, vco := range vcos {
			newNote := vco.NoteFromVoltage(voltage)
			vco.SendNote(newNote)
		}

		if Debug {
			note := vcos[0].NoteFromVoltage(voltage)
			log.Printf("readCV: %d\tvoltage: %f\tnote: %v\n", ReadCV(), ReadVoltage(), note)
		}

		time.Sleep(2 * time.Millisecond)
	}
}
