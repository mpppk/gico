package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mpppk/gico/gico"
	"context"
	"github.com/skratchdot/open-golang/open"
	"os"
)

// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse repo",
	Long: `browse repo`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client := gico.GetGitHubClient(context.Background(), os.Getenv("GICO_GITHUB_TOKEN"))
		issues, err := gico.GetIssues(ctx, client, "mpppk", "skylarq", nil)
		gico.PanicIfErrorExist(err)

		var issueTitles []string
		for _, issue := range issues {
			issueTitles = append(issueTitles, issue.GetTitle())
		}

		selectedIssueTitle, err := gico.PipeToPeco(issueTitles)
		gico.PanicIfErrorExist(err)

		issue, err := gico.FindIssue(issues, selectedIssueTitle)
		gico.PanicIfErrorExist(err)

		open.Run(issue.GetHTMLURL())
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
}
