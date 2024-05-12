package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sxmbaka/lacuna/web"
)

var morelikeCmd = &cobra.Command{
	Use:   "ml [flags]",
	Short: "@todo: short desc of more-like cmd",
	Long:  "@todo: verbose desc of more-like cmd",
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

		wordlist := web.GetData(&ml, &sl, &sp, max)
		printFormat(wordlist)
		fmt.Println("@todo: lacuna more-like")
	},
}

func init() {
	rootCmd.AddCommand(morelikeCmd)
	// rootCmd.PersistentFlags().Lookup("more-like").Hidden = true
	// rootCmd.PersistentFlags().Lookup("sounds-like").Hidden = true
	// rootCmd.PersistentFlags().Lookup("spelled-like").Hidden = true
	morelikeCmd.Flags().IntP("max", "x", 5, "@todo: max flag desc")
}
