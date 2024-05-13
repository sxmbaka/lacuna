package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sxmbaka/lacuna/web"
)

var spelledLikeCmd = &cobra.Command{
	Use:   "sp [word] [flags]...",
	Short: "Find words that are spelled like the given word",
	Long:  "Spelled like\nConstraint: require that the results have a spelling related to this string value, which can be any word or sequence of words.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) < 1 {
			cmd.Println("Please provide a word to search for with the sp command.")
			cmd.Println("Example: lacuna sp [word] [flags]...")
			cmd.Println("Run 'lacuna sp --help' for more information.")
			return
		}
		sp.Active = true
		sp.Arg = args[0]
		ml.Active = false
		sl.Active = false
		max, err := cmd.Flags().GetInt("max")
		cobra.CheckErr(err)

		wordlist := web.GetData(&ml, &sl, &sp, max)
		cmd.Println("Spelled like:", sp.Arg)
		printFormat(wordlist)
	},
}

func init() {
	rootCmd.AddCommand(spelledLikeCmd)
	spelledLikeCmd.Flags().IntP("max", "x", 10, "Maximum number of results to return")
}
