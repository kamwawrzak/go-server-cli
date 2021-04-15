package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// function priting help
func PrintHelp(cmd *cobra.Command, args []string) {
	runDesc := "run -file <filepath> - starts HTTP server serving file passed as argument \n"
	helpDesc := "help - prints available commands \n"
	versDesc := "version - prints current version of the application"
	fmt.Println("Available commands: \n" + runDesc + helpDesc + versDesc)
}

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print available commands",
	Long:  `Print description of commands run, help and version`,
	Run:   PrintHelp,
}

// init command
func init() {
	rootCmd.AddCommand(helpCmd)
}
