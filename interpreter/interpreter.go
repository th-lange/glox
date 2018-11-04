package interpreter

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/th-lange/glox/statusCodes"

	"github.com/th-lange/glox/scanner"
)

type Interpreter struct {
	Scnr         scanner.Scanner
	IgnoreErrors bool
}

func Init() Interpreter {
	return Interpreter{
		Scnr:         scanner.Scanner{},
		IgnoreErrors: false,
	}
}

func (intp *Interpreter) BreakOnError(isTrue bool) {
	intp.IgnoreErrors = isTrue
}

func (intp *Interpreter) run(lines string) {
	intp.runScanner(lines)
}

func (intp *Interpreter) runScanner(lines string) {
	intp.Scnr.Scan(lines)
	if intp.Scnr.HadError {
		for _, err := range intp.Scnr.Errors {
			fmt.Println(err.Error())
		}
		if !intp.IgnoreErrors {
			os.Exit(statusCodes.EXIT_DATA_ERROR)
		}
	}
}

func (intp *Interpreter) RunPrompt() {
	intp.IgnoreErrors = true
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		intp.run(text)
	}
}

func (intp *Interpreter) RunFiles(files ...string) {
	for _, item := range files {
		intp.runFile(item)
	}
}

func (intp *Interpreter) runFile(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("HadError! Could not read file: ", file)
	}
	fmt.Println("-------------------------------------------------------------------------------------------------------")
	fmt.Println("-- Interpreting:", file)
	fmt.Println("-------------------------------------------------------------------------------------------------------")

	intp.run(string(data))
}
