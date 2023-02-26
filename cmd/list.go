package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yammerjp/mdrun/markdown"
	"log"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show commands that defined in the target markdown file (default: README.md)",
	Long:  `The "mdrun list" subcommand produces standard output of the commands defined in the target markdown file, one line for each code block. This is the same as each line that is ambiguously searched when the "mdrun" command is executed.`,
	Run: func(cmd *cobra.Command, args []string) {
		commandBlocks, err := markdown.ExtractFile(targetFilePath)
		if err != nil {
			log.Fatal(err)
		}
		for _, commandBlock := range commandBlocks {
			fmt.Println(commandBlock.CommandOneLineString() + " #" + commandBlock.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
