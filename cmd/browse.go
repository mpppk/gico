package cmd

import (
	//"context"

	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
	"github.com/mpppk/gico/finder"
	"github.com/mpppk/gico/git"
	"github.com/mpppk/gico/utils"
	"github.com/spf13/viper"
	"strings"
	"github.com/mpppk/gico/etc"
	"context"
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

		var config etc.Config
		err = viper.Unmarshal(&config)
		utils.PanicIfErrorExist(err)

		if issueFlag {
			for _, host := range config.Hosts {
				if strings.Contains(originRemote.Service, host.HostType) {
					issue, err := finder.SelectIssueInteractive(ctx, host.HostType, host.OAuthToken, originRemote)
					utils.PanicIfErrorExist(err)
					open.Run(issue.GetHTMLURL())
					return
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
	browseCmd.Flags().BoolVarP(&issueFlag, "issue", "i", false, "browse issue")
}
