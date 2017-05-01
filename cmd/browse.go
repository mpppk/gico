package cmd

import (
	"context"
	"os"

	"github.com/mpppk/gico/gico"
	"github.com/spf13/cobra"
)

var issueFlag bool

// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse repo",
	Long:  `browse repo`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		originRemote, err := gico.GetOriginRemote()
		gico.PanicIfErrorExist(err)

		if issueFlag {
			gico.OpenIssuePageInteractive(ctx, os.Getenv("GICO_GITHUB_TOKEN"), originRemote)
		}
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
	browseCmd.Flags().BoolVarP(&issueFlag, "issue", "i", false, "browse issue")
}

