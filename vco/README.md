# 3 voice quantized digital vco

The script is using the [tone package](https://pkg.go.dev/tinygo.org/x/drivers/tone) to create 3 configurable digital square wave oscillators. Each oscillator can be quantized to a melodic scale and can set a root note for stacking chords or creating sub oscillators.

Demo video: [https://youtu.be/f9nFkzrO6-8](https://youtu.be/f9nFkzrO6-8)

|Uncertainty| Function |
| ------ | ------ |
|CV input|analog cv input 0v to 5v to control the vco pitch.|
|Gate 1|-|
|Gate 2|vco 1|
|Gate 3|-|
|Gate 4|vco 2|
|Gate 5|-|
|Gate 6|vco 3|
|Gate 7|-|
|Gate 8|-|

## Getting Started

See the [Getting Started](/README.md#getting-started) first if you have not yet done so.

## Configuration

Change up the scale and root note for each vco by editing the [configure.go](configure.go) file. The list of available quantizer scales is [here](scale.go#L19) and the note references is [here](https://pkg.go.dev/tinygo.org/x/drivers/tone#Note).

For example:

```go
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
```

## Building & Flashing

From the root directory of the project, use the `tinygo flash` command while Uncertainty USB is connected to compile the script and copy it to your Uncertainty.

```shell
tinygo flash --target xiao-rp2040 ./vco
```
