package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version of geasy",
	Long:  "All software has versions. This is geasy's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Geasy Static Web Server Engine v0.0.1 ")
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
