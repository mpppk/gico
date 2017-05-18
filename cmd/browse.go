package cmd

import (
	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
	"github.com/mpppk/gico/finder"
	"github.com/mpppk/gico/utils"
	"github.com/mpppk/hlb/hlb"
)

type FilterStringer interface {
	FilterString() string
}

var issueFlag bool
var prFlag bool
// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse repo",
	Long:  `browse repo`,
	Run: func(cmd *cobra.Command, args []string) {

		base, err := hlb.NewCmdBase()
		utils.PanicIfErrorExist(err)
		sw := hlb.ServiceWrapper{Base: base}

		if issueFlag {
			issues, err := sw.GetIssues()
			utils.PanicIfErrorExist(err)

			var filterableIssues []*finder.FilterableIssue
			for _, is := range issues {
				filterableIssues = append(filterableIssues, &finder.FilterableIssue{Issue: is})
			}

			var filterableStrings []finder.FilterableStringer
			for _, fis := range filterableIssues {
				filterableStrings = append(filterableStrings, finder.FilterableStringer(fis))
			}

			res, err := finder.Filter(filterableStrings)
			utils.PanicIfErrorExist(err)
			issue := res.(*finder.FilterableIssue)
			url := issue.GetHTMLURL()
			open.Run(url)
			return
		}

		if prFlag {
			prs, err := sw.GetPullRequests()
			utils.PanicIfErrorExist(err)

			var filterableIssues []*finder.FilterablePullRequest
			for _, is := range prs {
				filterableIssues = append(filterableIssues, &finder.FilterablePullRequest{PullRequest: is})
			}

			var filterableStrings []finder.FilterableStringer
			for _, fis := range filterableIssues {
				filterableStrings = append(filterableStrings, finder.FilterableStringer(fis))
			}

			res, err := finder.Filter(filterableStrings)
			utils.PanicIfErrorExist(err)
			pr := res.(*finder.FilterablePullRequest)
			url := pr.GetHTMLURL()
			open.Run(url)
			return
		}

		url, err := sw.GetRepositoryURL()
		utils.PanicIfErrorExist(err)
		open.Run(url)
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
	browseCmd.Flags().BoolVarP(&issueFlag, "issue", "i", false, "browse issue")
	browseCmd.Flags().BoolVarP(&prFlag, "pull-request", "p", false, "browse pull/merge request")
}
