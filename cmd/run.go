package cmd

import (
	"fmt"
	"path/filepath"
	"serverapp/webserver"

	"github.com/spf13/cobra"
)

// function running getting file path and running server
func RunServer(cmd *cobra.Command, args []string) {
	f, _ := cmd.Flags().GetString("file")
	fileExtension := filepath.Ext(f)
	fmt.Println(fileExtension)
	if f == "" {
		fmt.Println("You need to pass file path.")

	} else if fileExtension != ".html" {
		fmt.Println("Only .html files are allowed.")
	} else {
		fmt.Println("Server starting...")
		webserver.StartServer(f)
	}
}

// runCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Command starts HTTP server serving passed HTML file",
	Long:  `To start HTTP server use command run --file <file path>`,
	Run:   RunServer,
}

// init command
func init() {
	rootCmd.AddCommand(RunCmd)
	RunCmd.Flags().StringP("file", "f", "", "Set file path.")
}
