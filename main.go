package main

import (
	"github.com/th-lange/glox/base"
	"github.com/th-lange/glox/cmd"
	"path"
	"runtime"
)

var HomeDir string


func main() {

	_, executionPath, _, ok := runtime.Caller(0)
	if !ok{
		panic("No information about execution")
	}
	base.HomeDir = path.Dir(executionPath)
	cmd.Execute()
}

