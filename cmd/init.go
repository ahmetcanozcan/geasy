package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initilize geasy file structure",
	Long: `Geasy run web server on current woking directory
	from a configuration json file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Too much argument")
			os.Exit(1)
		}
		exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err)
		}
		sourceDir := fmt.Sprintf("%s/resources/geasy", exeDir)
		destDir, _ := os.Getwd()

		copyFiles(sourceDir, destDir)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func copyFiles(sourceDir string, destDir string) error {
	files := []string{"serv.json", "test.js"}

	for _, file := range files {
		err := func() error {
			srcFile := fmt.Sprintf("%s/%s", sourceDir, file)
			destFile := fmt.Sprintf("%s/%s", destDir, file)

			input, err := ioutil.ReadFile(srcFile)
			if err != nil {
				return err
			}

			return ioutil.WriteFile(destFile, input, 0644)

		}()

		if err != nil {
			return err
		}
	}
	return nil
}
