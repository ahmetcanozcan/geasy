package cmd

import (
	"fmt"
	"os"

	"github.com/ahmetcanozcan/geasy/server"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run geasy on current working directory.",
	Long: `Geasy run web server on current woking directory
	from a configuration json file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Too much argument")
			os.Exit(1)
		}
		cwd, _ := os.Getwd()
		filePath := fmt.Sprintf("%s\\%s", cwd, args[0])
		server.Run(filePath)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
