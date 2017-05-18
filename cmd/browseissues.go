package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mpppk/hlb/hlb"
	"github.com/mpppk/gico/utils"
	"github.com/mpppk/gico/finder"
	"github.com/skratchdot/open-golang/open"
)

// browseissuesCmd represents the browseissues command
var browseissuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		base, err := hlb.NewCmdBase()
		utils.PanicIfErrorExist(err)
		sw := finder.NewServiceWrapper(base)
		filterableIssues, err := sw.GetFilterableIssues()
		utils.PanicIfErrorExist(err)

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
	},
}

func init() {
	browseCmd.AddCommand(browseissuesCmd)
}
