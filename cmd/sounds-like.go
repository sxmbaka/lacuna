package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sxmbaka/lacuna/web"
)

var soundsLikeCmd = &cobra.Command{
	Use:   "sl [word] [flags]...",
	Short: "Find words that sound like the given word",
	Long:  "Sounds like command\nconstraint: require that the results are pronounced similarly to this string of characters. (If the string of characters doesn't have a known pronunciation, the system will make its best guess using a text-to-phonemes algorithm).\n\nWildcard characters:\n(do not work for means-like)\n- *: matches any number of characters\n- ?: matches exactly one character\n\nOutput Format:\n[index] [word] [score]\nIndex: the index of the word in the list\nWord: the word that matches the query\nScore: the score field has no interpretable meaning, other than as a way to rank the results\n\n",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) < 1 {
			cmd.Println("Please provide a word to search for with the sl command.")
			cmd.Println("Example: lacuna sl [word] [flags]...")
			cmd.Println("Run 'lacuna sl --help' for more information.")
			return
		}
		sl.Active = true
		sl.Arg = args[0]
		ml.Active = false
		sp.Active = false
		max, err := cmd.Flags().GetInt("max")
		cobra.CheckErr(err)
		show, err := cmd.Flags().GetBool("show-query")
		cobra.CheckErr(err)

		wordlist := web.GetData(&ml, &sl, &sp, max, show)
		cmd.Println("Sounds like:", sl.Arg)
		printFormat(wordlist, cmd.OutOrStdout())
	},
}

func init() {
	rootCmd.AddCommand(soundsLikeCmd)
	soundsLikeCmd.Flags().IntP("max", "x", 10, "maximum number of results to return")
	soundsLikeCmd.Flags().Bool("show-query", false, "show the query url that was used to fetch the results")
}
