package main

import (
	"log"
	"time"

	"tinygo.org/x/drivers/tone"

	uncertainty "github.com/awonak/UncertaintyGo"
)

// Enable to print serial monitoring log messages.
const Debug = false

var (
	// An array of VCOs bound to CV Outputs.
	vcos [3]VCO

	// Configurable variables for VCOs. Values set in configure.go.
	scales [3]Scale
	roots  [3]tone.Note
)

func main() {
	if Debug {
		// Provide a brief pause to allow time to start up the serial monitor to capture errors.
		log.Print("START...")
		time.Sleep(time.Second * 5)
		log.Print("Ready...")
	}

	// Initialize a collection of VCOs bound to a CV Output and provide the
	// configured scale and root note for each VCO.
	for i, output := range []uncertainty.Outputer{
		uncertainty.Outputs[1],
		uncertainty.Outputs[3],
		uncertainty.Outputs[5],
	} {
		vcos[i] = NewVCO(output.PWM(), output.Pin(), scales[i], roots[i])
	}

	// Main program loop.
	for {
		voltage := uncertainty.ReadVoltage()
		for _, vco := range vcos {
			newNote := vco.NoteFromVoltage(voltage)
			vco.SendNote(newNote)
		}

		if Debug {
			read := uncertainty.ReadCV()
			volts := uncertainty.ReadVoltage()
			note := vcos[0].NoteFromVoltage(voltage)
			log.Printf("readCV: %d\tvoltage: %f\tnote: %v\n", read, volts, note)
		}

		time.Sleep(2 * time.Millisecond)
	}
}
