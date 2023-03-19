package main

import "tinygo.org/x/drivers/tone"

func init() {

	// Define the scale for each VCO.
	scales = [voices]Scale{
		Major,
		MajorTriad,
		Octave,
	}

	// Define the root note for each VCO.
	roots = [voices]tone.Note{
		tone.C2,
		tone.E2,
		tone.C1,
	}
}
