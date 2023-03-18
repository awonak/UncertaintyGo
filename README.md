# Uncertainty

[![Go Reference](https://pkg.go.dev/badge/github.com/awonak/UncertaintyGo.svg)](https://pkg.go.dev/github.com/awonak/UncertaintyGo)

Firmware and scripts written for the [Uncertainty](https://oamodular.org/discount/AWONAK?redirect=%2Fproducts%2Funcertainty) eurorack module using [TinyGo](https://tinygo.org/).

> **Note**
> The Uncertainty package is under active develoment and does not guarantee backwards compatibity across commits.

### Scripts

⚡ **[Voltage Gates](voltage_gates/main.go)**

This one was taken directly from the main Uncertainty repo and rewritten in TinyGo. https://github.com/oamodular/uncertainty#coding-the-code

Demo video: [https://youtu.be/PLs5O3ZkTm0](https://youtu.be/PLs5O3ZkTm0)

🎹 **[3 Voice Quantized Digital VCO](vco/)**

The script is using the [tone package](https://pkg.go.dev/tinygo.org/x/drivers/tone) to create 3 configurable digital square wave oscillators. Each oscillator can be quantized to a melodic scale and can set a root note for stacking chords or creating sub oscillators.

Demo video: [https://youtu.be/f9nFkzrO6-8](https://youtu.be/f9nFkzrO6-8)

# Getting started

Install Go

[https://go.dev/doc/install](https://go.dev/doc/install)

Install TinyGo

[https://tinygo.org/getting-started/install](https://tinygo.org/getting-started/install)

Install the TinyGo VSCode plugin

[https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo](https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo)

## Build the example

From the root directory of the project, use the `tinygo flash` command while Uncertainty USB is connected to compile and flash the script.

```shell
tinygo flash --target xiao-rp2040 ./voltage_gates
```

## Serial printing

When your Uncertainty is connected via USB you can view printed serial output in VSCode's Serial Monitoring tab or from the terminal using a tool like `minicom`.

For example, a line in your code like:

```go
log.Printf("CV Input: %2.2f\n", ReadVoltage())
```

You can launch minicom to view the printed output:

```shell
minicom -b 115200 -o -D /dev/ttyACM0
```

## VSCode build task

Add the TinyGo flash command as your default build task:

```plaintext
Ctrl + Shift + P > Tasks: Configure Default Build Task
```

Use the following example task configuration to set tinygo flash as your default build command, which will build and flash the package of the current open file:

```json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "tinygo flash",
            "type": "shell",
            "command": "tinygo flash --target xiao-rp2040 ${fileDirname}",
            "group": {
                "kind": "build",
                "isDefault": true
            },
        }
    ]
}
```

Now you can build and flash your project using `Ctrl + Shift + B` or search for the command:

```plaintext
Ctrl + Shift + P > Tasks: Run Build Task
```
