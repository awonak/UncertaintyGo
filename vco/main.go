package main

import (
	"log"
	"time"
)

// Enable to print serial monitoring log messages.
const Debug = false

func main() {
	if Debug {
		// Provide a brief pause to allow time to start up the serial monitor to capture errors.
		log.Print("START...")
		time.Sleep(time.Second * 5)
		log.Print("Ready...")
	}

	var (
		// Initialize a collection of PWM VCOs bound to a scale.
		vco1 = NewVCO(pwmOutputs[1], cvOutputs[1], MajorPentatonic)
		vco2 = NewVCO(pwmOutputs[3], cvOutputs[3], MajorTriad)
		vco3 = NewVCO(pwmOutputs[5], cvOutputs[5], Octave)
		vcos = []VCO{vco1, vco2, vco3}
	)

	// Main program loop.
	for {
		newNote := NoteFromVoltage(ReadVoltage())
		for _, vco := range vcos {
			vco.SendNote(newNote)
		}

		if Debug {
			log.Printf("readCV: %d\tvoltage: %f\tnote: %v\n", ReadCV(), ReadVoltage(), newNote)
			time.Sleep(time.Millisecond * 10)
		}
	}
}
