package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/mpppk/gico/gico"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkout interactive",
	Long: `checkout interactive`,
	Run: func(cmd *cobra.Command, args []string) {
		err := gico.SwitchBranch()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(checkoutCmd)
}
