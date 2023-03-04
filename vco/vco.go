package main

import (
	"log"
	"machine"

	"tinygo.org/x/drivers/tone"
)

type VCO struct {
	output      tone.Speaker
	scale       Scale
	currentNote tone.Note
}

func NewVCO(pwm tone.PWM, pin machine.Pin, scale Scale) VCO {
	output, err := tone.New(pwm, pin)
	if err != nil {
		log.Fatalf("NewVCO(%v) error: %v", pin, err.Error())
	}

	return VCO{
		output:      output,
		scale:       scale,
		currentNote: tone.Note(0),
	}
}

// SendNote will check if the received note is present in the current vco scale and update the vco frequency if so.
func (vco *VCO) SendNote(note tone.Note) {
	if note == vco.currentNote {
		return
	}

	// Check if new note is in the quantized scale. If so, set the note.
	for _, n := range vco.scale {
		if note == n {
			vco.output.SetNote(note)
			vco.currentNote = note
		}
	}
}

// NoteFromVoltage gets the midi note number from a range of notes.
//
// For example, 60 notes (12 notes per octave * 5 octaves), starting at note number 24 (C1).
func NoteFromVoltage(v float64) tone.Note {
	noteNum := int(v/MaxReadVoltage*MaxNoteNum) + MinNoteNum
	return tone.Note(noteNum)
}
