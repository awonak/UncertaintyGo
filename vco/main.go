package main

import (
	"time"

	"tinygo.org/x/drivers/tone"

	uncertainty "github.com/awonak/UncertaintyGo"
)

var (
	// An array of VCOs bound to CV Outputs.
	vcos [3]VCO

	// Configurable variables for VCOs. Values set in configure.go.
	scales [3]Scale
	roots  [3]tone.Note
)

func main() {

	// Initialize a collection of VCOs bound to a CV Output and provide the
	// configured scale and root note for each VCO.
	for i, output := range []*uncertainty.Output{
		uncertainty.Outputs[1],
		uncertainty.Outputs[3],
		uncertainty.Outputs[5],
	} {
		vcos[i] = NewVCO(output, scales[i], roots[i])
	}

	// Main program loop.
	for {
		voltage := uncertainty.ReadVoltage()

		for _, vco := range vcos {
			newNote := vco.NoteFromVoltage(voltage)
			vco.SendNote(newNote)
		}

		time.Sleep(2 * time.Millisecond)
	}
}
