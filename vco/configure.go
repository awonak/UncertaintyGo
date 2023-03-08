package main

import "tinygo.org/x/drivers/tone"

func init() {

	// Define the scale for each VCO.
	scales = [3]Scale{
		Major,
		MajorTriad,
		Octave,
	}

	// Define the root note for each VCO.
	roots = [3]tone.Note{
		tone.C2,
		tone.E2,
		tone.C1,
	}
}
