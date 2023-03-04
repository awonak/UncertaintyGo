package main

import (
	"tinygo.org/x/drivers/tone"
)

const (
	// The first midi note number for 0v. 36 = C2
	MinNoteNum = 36

	// The max midi note number from a range of 12 notes per octave * 5 octaves + root note 24.
	MaxNoteNum = 12*5 + MinNoteNum
)

var (
	// A collection of common melodic scales.
	Chromatic       = NewScale([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	Major           = NewScale([]int{1, 3, 5, 6, 8, 10, 12})
	Minor           = NewScale([]int{1, 3, 4, 6, 8, 9, 11})
	MajorPentatonic = NewScale([]int{1, 3, 5, 8, 10})
	MinorPentatonic = NewScale([]int{1, 4, 6, 8, 11})
	MajorTriad      = NewScale([]int{1, 5, 8})
	MinorTriad      = NewScale([]int{1, 4, 8})
	Octave          = NewScale([]int{1})
)

type Scale []tone.Note

func NewScale(steps []int) Scale {
	var (
		scale Scale
		step  int = 1
	)

	// If there are no steps provided in the param, there's no work to be done.
	if len(steps) == 0 {
		return scale
	}

	// Iterate over all midi note numbers in our range to determine which notes belong in this scale.
	for note := MinNoteNum; note < MaxNoteNum; note++ {

		// Check if the current note's step within an octave is present in the notes parameter.
		for _, s := range steps {
			if step == s {
				scale = append(scale, tone.Note(note))
				break
			}
		}

		// Increment the step index within an octave range, resetting at 12 steps.
		step += 1
		if step > 12 {
			step = 1
		}
	}

	return scale
}
