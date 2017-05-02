package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
	"github.com/mpppk/gico/finder"
	"github.com/mpppk/gico/git"
	"github.com/mpppk/gico/utils"
)

var issueFlag bool

// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse repo",
	Long:  `browse repo`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		originRemote, err := git.GetOriginRemote()
		utils.PanicIfErrorExist(err)

		if issueFlag {
			issue, err := finder.SelectIssueInteractive(ctx, os.Getenv("GICO_GITHUB_TOKEN"), originRemote)
			utils.PanicIfErrorExist(err)

			open.Run(issue.GetHTMLURL())
		}
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
	browseCmd.Flags().BoolVarP(&issueFlag, "issue", "i", false, "browse issue")
}
