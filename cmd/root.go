package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sxmbaka/lacuna/web"
)

var rootCmd = &cobra.Command{
	Use:   "lacuna",
	Short: "@todo: short desc of root cmd",
	Long:  "lacuna is a word-finding query engine for developers. You can use it in your apps to find words that match a given set of constraints and that are likely in a given context. You can specify a wide variety of constraints on meaning, spelling, sound, and vocabulary in your queries, in any combination.",
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

		if ml.Active || sl.Active || sp.Active { // if any of the flags are active
			wordlist := web.GetData(&ml, &sl, &sp, max)
			printFormat(wordlist)
		} else {
			cmd.Help()
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringP("means-like", "m", "", "@todo: more-like flag desc")
	rootCmd.Flags().StringP("sounds-like", "s", "", "@todo: sounds-like flag desc")
	rootCmd.Flags().StringP("spelled-like", "p", "", "@todo: spelled-like flag desc")
	rootCmd.Flags().IntP("max", "x", 10, "@todo: max flag desc")
}
