package cmd

import (
	"fmt"

	"github.com/sxmbaka/lacuna/web"
)

func printFormat(wordlist []web.Word) {
	for _, word := range wordlist {
		fmt.Printf("Word: %s\n", word.Word)
		fmt.Printf("Score: %d\n", word.Score)
		fmt.Printf("Tags: %#v\n", word.Tags)
	}
}
