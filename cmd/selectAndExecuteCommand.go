package cmd

import (
	"fmt"
	"github.com/yammerjp/mdrun/executor"
	"github.com/yammerjp/mdrun/markdown"
	"github.com/yammerjp/mdrun/selector"
	"os"
)

func selectAndExecuteCommand(dryRun bool, targetFilePath string) {
	commandBlocks, err := markdown.ExtractFile(targetFilePath)
	if err != nil {
		fmt.Println("Error occured")
		fmt.Println(err)
		os.Exit(1)
	}
	selectedCmdStr := selector.Select(commandBlocks).CommandOneLineString()

	if dryRun {
		fmt.Println(selectedCmdStr)
	} else {
		executor.Execute(selectedCmdStr)
	}
}
