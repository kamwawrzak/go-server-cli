package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// application version
const AppVer = "1.0"

// function printing application version
func DisplayVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("Version: %s \n", AppVer)

}

// versionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current version of the application.",
	Long:  "Print current version of the application.",
	Run:   DisplayVersion,
}

// init command
func init() {
	rootCmd.AddCommand(VersionCmd)
}
