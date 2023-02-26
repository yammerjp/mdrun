package selector

import (
	"fmt"
	"os"
	"github.com/yammerjp/mdrun/markdown"
	fizzyfinder "github.com/ktr0731/go-fuzzyfinder"

)

func Select(commandBlocks []markdown.CommandBlock) markdown.CommandBlock {
		idx, err := fizzyfinder.Find(
			commandBlocks,
			func(i int) string {
				cb := commandBlocks[i]
				return cb.CommandOneLineString() + " # " + cb.Title
			})
		if err != nil {
			fmt.Println("Error occured")
			fmt.Println(err)
			os.Exit(1)
		}
		return commandBlocks[idx]
}
