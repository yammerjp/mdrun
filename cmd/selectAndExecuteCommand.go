package cmd

import (
	"fmt"
	"github.com/yammerjp/mdrun/executor"
	"github.com/yammerjp/mdrun/markdown"
	"github.com/yammerjp/mdrun/selector"
	"log"
)

func selectAndExecuteCommand(dryRun bool, targetFilePath string) {
	commandBlocks, err := markdown.ExtractFile(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}

	selectedCmdStr := selector.Select(commandBlocks).CommandOneLineString()

	if dryRun {
		fmt.Println(selectedCmdStr)
	} else {
		executor.Execute(selectedCmdStr)
	}
}
