package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/th-lange/glox/base"
	"github.com/th-lange/glox/util"
)

var visitorName string
var visitorPackageName string

var generateVisitorCmd = &cobra.Command{
	Use:   "generateVisitor",
	Short: "Generates the visitor boilerplate code",
	Long: `This creates the boilerplate visitor code.
It will setup the files in the "visitor" folder. It will overwrite existing files!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateVisitor called. Target: " + base.HomeDir)
		util.GenerateVisitor(base.HomeDir, visitorPackageName, visitorName)
	},
}

func init() {
	generateVisitorCmd.Flags().StringVarP(&visitorName, "visitor", "v", "ExampleVisitor", "The Visitor Name (required).")
	generateVisitorCmd.Flags().StringVarP(&visitorPackageName, "target", "t", "visitor", "Target Path and Package of the Visitor")
	rootCmd.AddCommand(generateVisitorCmd)
}
