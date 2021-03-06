package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/th-lange/glox/base"
	"github.com/th-lange/glox/util"
)

var packageName string

var generateAstCmd = &cobra.Command{
	Use:   "generateAst",
	Short: "Generates the Parser AST",
	Long: `This creates the AST files, needed by the parser.
It will setup the files in the "expressions" folder. Any previous files will be overwritten!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createAst called. Target: " + base.HomeDir)
		util.GenerateAst(base.HomeDir, packageName)
	},
}

func init() {
	generateAstCmd.Flags().StringVarP(&packageName, "target", "t", "expression", "Target Path of the ast")
	rootCmd.AddCommand(generateAstCmd)
}
