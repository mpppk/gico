package cmd

import (
	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
	"github.com/mpppk/gico/utils"
	"github.com/mpppk/hlb/hlb"
	"github.com/mpppk/gico/finder"
)

// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse repo",
	Long:  `browse repo`,
	Run: func(cmd *cobra.Command, args []string) {
		base, err := hlb.NewCmdBase()
		utils.PanicIfErrorExist(err)
		sw := finder.NewServiceWrapper(base)

		var links []finder.FilterableLink
		repoUrl, err := sw.GetRepositoryURL()
		utils.PanicIfErrorExist(err)
		issuesUrl, err := sw.GetIssuesURL()
		utils.PanicIfErrorExist(err)
		pullsUrl, err := sw.GetPullRequestsURL()
		utils.PanicIfErrorExist(err)

		links = append(links,
			&finder.FilterableUrl{Url: repoUrl, FilterStr: "*repo"},
			&finder.FilterableUrl{Url: issuesUrl, FilterStr: "#issues"},
			&finder.FilterableUrl{Url: pullsUrl, FilterStr: "!pullrequests"},
		)

		issues, err := sw.GetFilterableIssues()
		utils.PanicIfErrorExist(err)
		for _, issue := range issues {
			links = append(links, issue)
		}

		pulls, err := sw.GetFilterablePullRequests()
		utils.PanicIfErrorExist(err)
		for _, pull := range pulls {
			links = append(links, pull)
		}

		link, err := finder.FilterLinks(links)
		utils.PanicIfErrorExist(err)
		open.Run(link.GetHTMLURL())
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
}
