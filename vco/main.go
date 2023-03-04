package main

import (
	"log"
	"machine"
	"time"
)

// Enable to print serial monitoring log messages.
const Debug = true

var (
	// Initialize a collection of PWM VCOs bound to a scale.
	vco1 = NewVCO(pwmOutputs[2], cvOutputs[2], MajorPentatonic)
	vco2 = NewVCO(pwmOutputs[4], cvOutputs[4], MajorTriad)
	vco3 = NewVCO(pwmOutputs[6], cvOutputs[6], Octave)
	vcos = []VCO{vco1, vco2, vco3}
)

// TODO: turn these into constructors and move definiton to script.
func init() {
	// Initialize the cv input GPIO as an analog input.
	machine.InitADC()
	cvInput = machine.ADC{Pin: CVInput}
	cvInput.Configure(machine.ADCConfig{})

	// Create an array of our cv outputs and configure for output.
	cvOutputs = [8]machine.Pin{CV1, CV2, CV3, CV4, CV5, CV6, CV7, CV8}
	for _, cv := range cvOutputs {
		cv.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	// Configure each of the cv outputs with their PWM peripherial channel.
	pwmOutputs = [8]PWM{
		machine.PWM5, // GPIO27 peripherals: PWM5 channel B
		machine.PWM6, // GPIO28 peripherals: PWM6 channel A
		machine.PWM6, // GPIO29 peripherals: PWM6 channel B
		machine.PWM0, // GPIO0  peripherals: PWM0 channel A
		machine.PWM1, // GPIO3  peripherals: PWM1 channel B
		machine.PWM2, // GPIO4  peripherals: PWM2 channel A
		machine.PWM1, // GPIO2  peripherals: PWM1 channel A
		machine.PWM0, // GPIO1  peripherals: PWM0 channel B
	}
	for _, pwm := range pwmOutputs {
		err := pwm.Configure(machine.PWMConfig{
			Period: defaultPeriod,
		})
		if err != nil {
			log.Fatalf("pwm(%v) Configure error: %v", pwm, err.Error())
		}
	}
}

func main() {
	if Debug {
		// Provide a brief pause to allow time to start up the serial monitor to capture errors.
		log.Print("START...")
		time.Sleep(time.Second * 5)
		log.Print("Ready...")
	}

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
