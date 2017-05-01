package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/mpppk/gico/finder"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "git branch",
	Long:  "git branch",
	Run: func(cmd *cobra.Command, args []string) {
		branchName, err := finder.GetBranchInteractive()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(branchName)
	},
}

func init() {
	RootCmd.AddCommand(branchCmd)
}
