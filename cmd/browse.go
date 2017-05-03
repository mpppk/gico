package cmd

import (
	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
	"github.com/mpppk/gico/finder"
	"github.com/mpppk/gico/git"
	"github.com/mpppk/gico/utils"
	"github.com/spf13/viper"
	"strings"
	"github.com/mpppk/gico/etc"
	"context"
	"github.com/mpppk/gico/project"
)

var issueFlag bool
var prFlag bool
// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse repo",
	Long:  `browse repo`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		var config etc.Config
		err := viper.Unmarshal(&config)
		utils.PanicIfErrorExist(err)

		originRemote, err := git.GetOriginRemote(config.Hosts)
		utils.PanicIfErrorExist(err)

		for _, host := range config.Hosts {
			if !strings.Contains(originRemote.HostType, host.HostType) {
				continue
			}

			if issueFlag {
				issue, err := finder.SelectIssueInteractive(ctx, host.HostType, host.OAuthToken, originRemote)
				utils.PanicIfErrorExist(err)
				open.Run(issue.GetHTMLURL())
			}else if prFlag {
				pr, err := finder.SelectPullRequestInteractive(ctx, host.HostType, host.OAuthToken, originRemote)
				utils.PanicIfErrorExist(err)
				open.Run(pr.GetHTMLURL())
			}else {
				repo, err := project.GetRepository(ctx, host.HostType, host.OAuthToken, originRemote.Owner, originRemote.RepoName)
				utils.PanicIfErrorExist(err)
				open.Run(repo.GetHTMLURL())
			}
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
	browseCmd.Flags().BoolVarP(&issueFlag, "issue", "i", false, "browse issue")
	browseCmd.Flags().BoolVarP(&prFlag, "pull-request", "p", false, "browse pull/merge request")
}
