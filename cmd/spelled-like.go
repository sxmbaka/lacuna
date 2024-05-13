package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sxmbaka/lacuna/web"
)

var spelledLikeCmd = &cobra.Command{
	Use:   "sp [word] [flags]...",
	Short: "Find words that are spelled like the given word",
	Long:  "Spelled like command\nConstraint: require that the results have a spelling related to this string value, which can be any word or sequence of words.\n\nWildcard characters:\n(do not work for means-like)\n- *: matches any number of characters\n- ?: matches exactly one character\n\nOutput Format:\n[index] [word] [score]\nIndex: the index of the word in the list\nWord: the word that matches the query\nScore: the score field has no interpretable meaning, other than as a way to rank the results\n\n",
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
		show, err := cmd.Flags().GetBool("show-query")
		cobra.CheckErr(err)

		wordlist := web.GetData(&ml, &sl, &sp, max, show)
		cmd.Println("Spelled like:", sp.Arg)
		printFormat(wordlist, cmd.OutOrStdout())
	},
}

func init() {
	rootCmd.AddCommand(spelledLikeCmd)
	spelledLikeCmd.Flags().IntP("max", "x", 10, "maximum number of results to return")
	spelledLikeCmd.Flags().Bool("show-query", false, "show the query url that was used to fetch the results")
}
