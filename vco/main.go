package main

import (
	"log"
	"time"

	"tinygo.org/x/drivers/tone"
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
		vcos = [3]VCO{
			NewVCO(pwmOutputs[1], cvOutputs[1], MajorPentatonic, tone.C2),
			NewVCO(pwmOutputs[3], cvOutputs[3], MajorTriad, tone.G2),
			NewVCO(pwmOutputs[5], cvOutputs[5], Octave, tone.C1),
		}
	)

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
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1 * time.Millisecond)
		}
	}
}
