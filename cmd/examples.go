package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var examplesCmd = &cobra.Command{
	Use:                   "examples",
	Short:                 "Print examples of how to use the lacuna CLI",
	Long:                  `lacuna is a word-finding query engine for developers. You can use it in your apps to find words that match a given set of constraints and that are likely in a given context. You can specify a wide variety of constraints on meaning, spelling, sound, and vocabulary in your queries, in any combination.`,
	Args:                  cobra.NoArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		data := [][]string{
			{`words with a meaning similar to "ringing in the ears"`, `lacuna ml "ringing in the ears"`},
			{`words related to duck that start with the letter b`, `lacuna --means-like duck --spelled-like "b*"`},
			{`words that sound like "cry"`, `lacuna --sounds-like cry`},
			{`words that start with "a" and end with "e"`, `lacuna --spelled-like "a*e"`},
			{`words that start with t, end in k, and have two letters in between`, `lacuna -p "t??L"
lacuna sp "t??k"`},
		}

		// Create a new table
		table := tablewriter.NewWriter(cmd.OutOrStdout())

		// Set the column headers
		table.SetHeader([]string{"Description", "Command"})

		// Set the column width
		table.SetColWidth(40)
		table.SetRowLine(true)

		// Wrap text if it exceeds the column width
		table.SetAutoWrapText(true)

		// Add data to the table
		for _, row := range data {
			table.Append(row)
		}

		// Render the table
		table.Render()

		cmd.Println("Tip: If you face problems when giving arguments with spaces of special characters, try enclosing the argument in double quotes. Infact it is a good practice to always enclose the argument in double quotes.")
	},
}

func init() {
	rootCmd.AddCommand(examplesCmd)
}
