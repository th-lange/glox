package main

import (
	"os"

	"github.com/th-lange/glox/interpreter"
)

func main() {
	intpr := interpreter.Init()
	args := os.Args
	if len(args) == 1 {
		intpr.RunPrompt()
	} else {
		intpr.RunFiles(args[1:]...)
	}

}
