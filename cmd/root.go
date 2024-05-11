package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sxmbaka/lacuna/web"
)

var rootCmd = &cobra.Command{
	Use:   "lacuna",
	Short: "@todo: short desc of root cmd",
	Long:  "@todo: verbose dec of root cmd",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		ml.Active = cmd.Flags().Changed("more-like")
		if ml.Active {
			ml.Arg, err = cmd.Flags().GetString("more-like")
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

		web.GetData(&ml, &sl, &sp, max)
		fmt.Println("@todo: lacuna root")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringP("more-like", "m", "", "@todo: more-like flag desc")
	rootCmd.Flags().StringP("sounds-like", "s", "", "@todo: sounds-like flag desc")
	rootCmd.Flags().StringP("spelled-like", "p", "", "@todo: spelled-like flag desc")
	rootCmd.Flags().IntP("max", "x", 5, "@todo: max flag desc")
}
