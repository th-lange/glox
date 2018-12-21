package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/th-lange/glox/interpreter"
	"os"
)

var Debug int8

var rootCmd = &cobra.Command{
	Use:   "glox",
	Short: "g-lox is a interpreter written in go",
	Run: func(cmd *cobra.Command, args []string) {

		intpr := interpreter.Init(Debug)
		if len(args) == 0 {
			intpr.RunPrompt()
		} else {
			intpr.RunFiles(args...)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().Int8VarP(&Debug, "debug", "d", 0, "Debugging level and verbosity")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}