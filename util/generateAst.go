package util

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	DefaultStructSpacing = 15
)

var astDefinition [][]string = [][]string{
	[]string{"Binary", "Expression", "Left", "Token", "Operator", "Expression", "Right"},
	[]string{"Grouping", "Expression", "Expr"},
	[]string{"Literal", "Object", "Value"},
	[]string{"Unary", "Token", "Operator", "Expression", "Right"},
}

func GenerateAst(homeDir, packageName string) {

	basePath := homeDir + string(os.PathSeparator) + packageName
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		os.Mkdir(basePath, os.ModePerm)
	}
	writeToFile(basePath+string(os.PathSeparator)+"Warning.md", generateWarining())

	for _, element := range astDefinition {

		baseName := element[0]
		file := basePath + string(os.PathSeparator) + strings.ToLower(baseName) + ".go"
		structDefinition := generateStruct(packageName, baseName, element[1:]...)
		writeToFile(file, structDefinition)

	}

}

func writeToFile(file string, content string) {
	err := ioutil.WriteFile(file, []byte(content), 0644)
	if err != nil {
		panic("Could not write file: " + file + "!\nWith content: \n" + content + "\nCowardly refusing to carry on! Exiting!\nError:" + err.Error())
	}
}

func generateStruct(packageName, baseName string, elements ...string) string {
	header := "package " + packageName + "\n\n\n"

	structDefinition := "type " + baseName + " struct {\n"

	for i := 0; i < len(elements); i += 2 {
		spacing := DefaultStructSpacing - len(elements[i+1])
		structDefinition += "    " + elements[i+1] + strings.Repeat(" ", spacing) + elements[i] + "\n"
	}

	return header + structDefinition + "}\n"
}

func generateWarining() string {
	return `# Warining

This folder is autogenerated.
Do not modify it's content! It will be overwritten!'
`
}
