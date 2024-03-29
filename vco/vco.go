package main

import (
	"log"

	"tinygo.org/x/drivers/tone"

	uncertainty "github.com/awonak/UncertaintyGo"
)

// NoteRange represents the number of notes allowed in a 5v range.
// For example, 60 notes (12 notes per octave * 5 octaves), starting at note
// number 24 (C1).
const NoteRange = 60

// VCO is a configured pwm cv output that can play the notes from the given
// Scale across 5 volts starting at the given root note when 0v present.
type VCO struct {
	speaker     tone.Speaker
	scale       Scale
	rootNote    int
	currentNote tone.Note
}

// NewVCO returns a constructed VCO for the given configuration parameters.
func NewVCO(output *uncertainty.Output, scale Scale, rootNote tone.Note) VCO {
	// Set PWM frequency very high for best audio rate quality.
	output.PWM.SetPeriod(uint64(1e9 / 500))

	speaker, err := tone.New(output.PWM, output.Pin)
	if err != nil {
		log.Fatalf("NewVCO(%v) error: %v", output, err.Error())
	}

	return VCO{
		speaker: speaker,
		scale:   scale,
		// Add octave offset to account for the inaccurate tone.Note values.
		rootNote:    int(rootNote + 12),
		currentNote: tone.Note(rootNote + 12),
	}
}

// SendNote will check if the received note is present in the current vco
// scale and update the vco frequency if so.
func (vco *VCO) SendNote(note tone.Note) {
	if note == vco.currentNote {
		return
	}

	// Check if new note is in the quantized scale. If so, set the note.
	for _, n := range vco.scale {
		if note == n {
			vco.speaker.SetNote(note)
			vco.currentNote = note
		}
	}
}

// NoteFromVoltage gets the note in scale corresponding to the current voltage.
//
// For example, 60 notes (12 notes per octave * 5 octaves), starting at note
// number 24 (C1).
func (vco *VCO) NoteFromVoltage(v float64) tone.Note {
	scaleNum := int(v / float64(uncertainty.MaxReadVoltage) * NoteRange)
	noteNum := scaleNum + vco.rootNote
	return tone.Note(noteNum)
}
