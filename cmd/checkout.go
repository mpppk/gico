package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/libgit2/git2go.v25"
	"github.com/mpppk/gico/gico"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkout interactive",
	Long: `checkout interactive`,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := "."
		r, err := git.OpenRepository(repoPath)
		if err != nil {
			fmt.Println(err)
		}

		err = gico.SwitchBranch(r)

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(checkoutCmd)
}
