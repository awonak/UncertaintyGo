package uncertainty // import uncertainty "github.com/awonak/UncertaintyGo"

import (
	"log"
	"machine"
	"math"
)

const (
	// GPIO mapping to Uncertainty panel.
	CVInput = machine.ADC0
	CV1     = machine.GPIO27
	CV2     = machine.GPIO28
	CV3     = machine.GPIO29
	CV4     = machine.GPIO0
	CV5     = machine.GPIO3
	CV6     = machine.GPIO4
	CV7     = machine.GPIO2
	CV8     = machine.GPIO1

	// Number of times to read analog input for an average reading.
	ReadSamples = 500

	// Calibrated average min read uint16 voltage within a 0-5v range.
	MinCalibratedRead = 415

	// Calibrated average max read uint16 voltage within a 0-5v range.
	MaxCalibratedRead = 29582

	// Upper limit of voltage read by the cv input.
	MaxReadVoltage float64 = 5
)

var (
	// The array of 8 configured cv outputs
	Outputs [8]*Output

	// Create package global variables for the cv input and outputs.
	cvInput    machine.ADC
	cvOutputs  [8]machine.Pin
	pwmOutputs [8]PWM
)

// PWM is the interface necessary for configuring a cv output for PWM.
type PWM interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Top() uint32
	Set(channel uint8, value uint32)
	SetPeriod(period uint64) error
}

// Output represents a single cv output.
type Output struct {
	Pin machine.Pin
	PWM PWM
	ch  uint8
}

// High will set the current output to a high voltage of roughly 5v.
func (o *Output) High() {
	o.Pin.High()
}

// Low will set the current output to a low voltage of roughly 0v.
func (o *Output) Low() {
	o.Pin.Low()
}

func newOutput(pin machine.Pin, pwm PWM) *Output {
	ch, err := pwm.Channel(pin)
	if err != nil {
		log.Fatal("pwm Channel error: ", err.Error())
	}

	return &Output{
		Pin: pin,
		PWM: pwm,
		ch:  ch,
	}
}

// Read will return the cv input scaled to 0v-5v as an int with 0 for 0v and 32768 for 5v.
func Read() int {
	var sum int
	for i := 0; i < ReadSamples; i++ {
		read := int(cvInput.Get()) - math.MaxInt16
		if read < 0 {
			read = 0
		}
		sum += read
	}
	return sum / ReadSamples
}

// ReadVoltage will return the cv input scaled to 0v-5v as a float with 0.0 for 0v and 5.0 for 5v.
func ReadVoltage() float64 {
	read := Read()
	return MaxReadVoltage * (float64(read-MinCalibratedRead) / float64(MaxCalibratedRead-MinCalibratedRead))
}

func init() {
	// Initialize the cv input GPIO as an analog input.
	machine.InitADC()
	cvInput = machine.ADC{Pin: CVInput}
	cvInput.Configure(machine.ADCConfig{})

	// Create an array of our cv outputs and configure for output.
	cvOutputs = [8]machine.Pin{CV1, CV2, CV3, CV4, CV5, CV6, CV7, CV8}

	// Configure each of the cv outputs with their PWM peripherial channel.
	//
	// Note: PWM pins on the same peripherial will overwrite eachother.
	// For example, cv out 2 (GPIO28) and cv out 3 (GPIO29) both use PWM6,
	// so whenever you set the frequency of one, the other will update to
	// that same frequency too.
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

	// Create our 8 configured outputs with Pin and PWM configurations per output.
	for i, cv := range cvOutputs {
		Outputs[i] = newOutput(cv, pwmOutputs[i])
		Outputs[i].Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}
