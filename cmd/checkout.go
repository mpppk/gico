package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/mpppk/gico/finder"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkout interactive",
	Long: `checkout interactive`,
	Run: func(cmd *cobra.Command, args []string) {
		err := finder.SwitchBranchInteractive()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(checkoutCmd)
}
