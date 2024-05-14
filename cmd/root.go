package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sxmbaka/lacuna/web"
)

var rootCmd = &cobra.Command{
	Use:   "lacuna",
	Short: "@todo: short desc of root cmd",
  Long:  "lacuna is a word-finding query engine for developers. You can use it in your apps to find words that match a given set of constraints and that are likely in a given context. You can specify a wide variety of constraints on meaning, spelling, sound, and vocabulary in your queries, in any combination.\nSee lacuna examples for some usage examples.\nTwo ways to use the tool:\n\n1. Root command: when you need to use all the features in combination use the root command and use the proper flags for structuring your query.\ne.g., lacuna -m light -p bun\n\n2. Specific commands: when you need to use only a single feature use the command associated with it. You CANNOT use other features when using a specific feature command\ne.g., lacuna ml justice\n\nWildcard characters:\n(do not work for means-like)\n- *: matches any number of characters\n- ?: matches exactly one character\n\nOutput Format:\n[index] [word] [score]\nIndex: the index of the word in the list\nWord: the word that matches the query\nScore: the score field has no interpretable meaning, other than as a way to rank the results\n\n",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		ml.Active = cmd.Flags().Changed("means-like")
		if ml.Active {
			ml.Arg, err = cmd.Flags().GetString("means-like")
			cobra.CheckErr(err)
		}

		sl.Active = cmd.Flags().Changed("sounds-like")
		if sl.Active {
			sl.Arg, err = cmd.Flags().GetString("sounds-like")
			cobra.CheckErr(err)
		}

		sp.Active = cmd.Flags().Changed("spelled-like")
		if sp.Active {
			sp.Arg, err = cmd.Flags().GetString("spelled-like")
			cobra.CheckErr(err)
		}
		max, err := cmd.Flags().GetInt("max")
		cobra.CheckErr(err)
		show, err := cmd.Flags().GetBool("show-query")
		cobra.CheckErr(err)

		if ml.Active || sl.Active || sp.Active { // if any of the flags are active
			wordlist := web.GetData(&ml, &sl, &sp, max, show)

			printFormat(wordlist, cmd.OutOrStdout())
		} else {
			cmd.Help()
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringP("means-like", "m", "", "when set, find words that are more like the given word")
	rootCmd.Flags().StringP("sounds-like", "s", "", "when set, find words that sound like the given word")
	rootCmd.Flags().StringP("spelled-like", "p", "", "when set, find words that are spelled like the given word")
	rootCmd.Flags().IntP("max", "x", 10, "maximum number of results to return")
	rootCmd.Flags().Bool("show-query", false, "show the query url that was used to fetch the results")
}
