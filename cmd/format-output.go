package cmd

import (
	"fmt"

	"github.com/sxmbaka/lacuna/web"
)

func printFormat(wordlist []web.Word) {
	for i, word := range wordlist {
		fmt.Printf("%4d: %s\n", i+1, word.Word)
		// fmt.Printf("Score: %d\n", word.Score)
		// we don't need to print tags
		// fmt.Printf("Tags: %#v\n", word.Tags)
	}
}
