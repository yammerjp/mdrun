package selector

import (
	"fmt"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/yammerjp/mdrun/markdown"
)

func Select(commandBlocks []markdown.CommandBlock) markdown.CommandBlock {
	idx, err := fuzzyfinder.Find(
		commandBlocks,
		func(i int) string {
			cb := commandBlocks[i]
			return cb.CommandOneLineString() + " # " + cb.Title
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			return fmt.Sprintf(
				"%s\n---\n\n```\n%s```",
				commandBlocks[i].Title,
				commandBlocks[i].Command,
			)
		}),
	)
	if err != nil {
		fmt.Println("Error occured")
		fmt.Println(err)
		os.Exit(1)
	}
	return commandBlocks[idx]
}
