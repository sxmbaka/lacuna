package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sxmbaka/lacuna/web"
)

var meansLikeCmd = &cobra.Command{
	Use:   "ml [word] [flags]...",
	Short: "Find words that are more like the given word",
	Long:  "Means like command\nconstraint: require that the results have a meaning related to this string value, which can be any word or sequence of words.\nThis is effectively the reverse dictionary feature of OneLook (https://www.onelook.com/reverse-dictionary.shtml).\n\nWildcard characters:\n(do not work for means-like)\n- *: matches any number of characters\n- ?: matches exactly one character\n\nOutput Format:\n[index] [word] [score]\nIndex: the index of the word in the list\nWord: the word that matches the query\nScore: the score field has no interpretable meaning, other than as a way to rank the results\n\n",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) < 1 {
			cmd.Println("Please provide a word to search for with the ml command.")
			cmd.Println("Example: lacuna ml [word] [flags]...")
			cmd.Println("Run 'lacuna ml --help' for more information.")
			return
		}
		ml.Active = true
		ml.Arg = args[0]
		sl.Active = false
		sp.Active = false
		max, err := cmd.Flags().GetInt("max")
		cobra.CheckErr(err)
		show, err := cmd.Flags().GetBool("show-query")
		cobra.CheckErr(err)

		wordlist := web.GetData(&ml, &sl, &sp, max, show)
		cmd.Println("More like:", ml.Arg)
		printFormat(wordlist, cmd.OutOrStdout())
	},
}

func init() {
	rootCmd.AddCommand(meansLikeCmd)
	meansLikeCmd.Flags().IntP("max", "x", 10, "maximum number of results to return")
	meansLikeCmd.Flags().Bool("show-query", false, "show the query url that was used to fetch the results")
}
