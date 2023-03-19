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
	MaxReadVoltage float32 = 5.0

	// Min and Max voltage range available to be used by the cv output.
	MinVoltage float32 = 0.0
	MaxVoltage float32 = 5.0

	// The default PWM frequncy is 100khz (1 second in nanoseconds / 100k).
	// This results in a period of 10,000ns per cycle.
	defaultPeriod uint64 = 1e9 / 100_000
)

var (
	// The array of 8 configured cv outputs
	Outputs [8]*Output

	// Package private variable for the cv input peripherial.
	cvInput machine.ADC
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
	o.PWM.Set(o.ch, o.PWM.Top())
}

// Low will set the current output to a low voltage of roughly 0v.
func (o *Output) Low() {
	o.PWM.Set(o.ch, 0)
}

// Voltage sets the current output voltage within a range of 0.0 to 5.0.
func (o *Output) Voltage(v float32) {
	v = clamp(v, MinVoltage, MaxVoltage)
	cv := (v / MaxVoltage) * float32(o.PWM.Top())
	o.PWM.Set(o.ch, uint32(cv))
}

func NewOutput(pin machine.Pin, pwm PWM) *Output {
	// Configure the PWM with the default period.
	err := pwm.Configure(machine.PWMConfig{
		Period: defaultPeriod,
	})
	if err != nil {
		log.Fatalf("pwm Configure(%v) error: %v", pin, err.Error())
	}

	ch, err := pwm.Channel(pin)
	if err != nil {
		log.Fatalf("pwm Channel(%v) error: %v", pin, err.Error())
	}

	return &Output{pin, pwm, ch}
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
func ReadVoltage() float32 {
	read := Read()
	return MaxReadVoltage * (float32(read-MinCalibratedRead) / float32(MaxCalibratedRead-MinCalibratedRead))
}

type clampable interface {
	~uint8 | ~uint16 | ~int | ~float32
}

func clamp[V clampable](value, low, high V) V {
	if value >= high {
		return high
	}
	if value <= low {
		return low
	}
	return value
}

func init() {
	// Initialize the cv input GPIO as an analog input.
	machine.InitADC()
	cvInput = machine.ADC{Pin: CVInput}
	cvInput.Configure(machine.ADCConfig{})

	// Create 8 configured outputs with Pin and PWM configurations per output.
	Outputs = [8]*Output{
		NewOutput(CV1, machine.PWM5), // GPIO27 peripherals: PWM5 channel B
		NewOutput(CV2, machine.PWM6), // GPIO28 peripherals: PWM6 channel A
		NewOutput(CV3, machine.PWM6), // GPIO29 peripherals: PWM6 channel B
		NewOutput(CV4, machine.PWM0), // GPIO0  peripherals: PWM0 channel A
		NewOutput(CV5, machine.PWM1), // GPIO3  peripherals: PWM1 channel B
		NewOutput(CV6, machine.PWM2), // GPIO4  peripherals: PWM2 channel A
		NewOutput(CV7, machine.PWM1), // GPIO2  peripherals: PWM1 channel A
		NewOutput(CV8, machine.PWM0), // GPIO1  peripherals: PWM0 channel B
	}
}
