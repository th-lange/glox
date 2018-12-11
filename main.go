package main

import (
	"os"

	"github.com/th-lange/glox/interpreter"
)

var debug bool

func main() {

	//TODO: Move to COBRA
	debug = false

	intpr := interpreter.Init(debug)
	args := os.Args
	if len(args) == 1 {
		intpr.RunPrompt()
	} else {
		intpr.RunFiles(args[1:]...)
	}

}
