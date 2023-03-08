package main

import (
	"tinygo.org/x/drivers/tone"
)

const (
	// The first midi note number. 24 = C1
	MinNoteNum = 24

	// The max midi note number. 108 = C8
	MaxNoteNum = 108
)

var (
	// A collection of common melodic scales.
	// Scales are represented as the semitone steps that should played in that scale.
	// For example the numbers 1, 5, 8 would represent a Major Triad (C, E, G).
	Chromatic       = NewScale([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	Major           = NewScale([]int{1, 3, 5, 6, 8, 10, 12})
	Minor           = NewScale([]int{1, 3, 4, 6, 8, 9, 11})
	Blues           = NewScale([]int{1, 4, 6, 7, 8, 11})
	MajorPentatonic = NewScale([]int{1, 3, 5, 8, 10})
	MinorPentatonic = NewScale([]int{1, 4, 6, 8, 11})
	MajorTriad      = NewScale([]int{1, 5, 8})
	MinorTriad      = NewScale([]int{1, 4, 8})
	Octave          = NewScale([]int{1})
)

// Scale is a slice of ints representing which of the 12 steps in an octave should be played.
//
// For example the numbers 1, 5, 8 would represent a Major Triad (C, E, G).
type Scale []tone.Note

// NewScale takes in a slice of ints in the range of 1..12 in length and creates a slice containing all the notes from those steps from C1 to C7.
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
