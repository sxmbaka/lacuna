package cmd

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/sxmbaka/lacuna/web"
)

func printFormat(wordlist []web.Word, cmd io.Writer) {
	table := tablewriter.NewWriter(cmd)
	table.SetHeader([]string{"Index", "Word", "Score"})
	for i, word := range wordlist {
		table.Append([]string{fmt.Sprintf("%d", i+1), word.Word, fmt.Sprintf("%d", word.Score)})
	}
	table.Render()
}
