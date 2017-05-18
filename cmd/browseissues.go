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
	Short: "browse issues",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		base, err := hlb.NewCmdBase()
		utils.PanicIfErrorExist(err)
		sw := finder.NewServiceWrapper(base)
		filterableIssues, err := sw.GetFilterableIssues()
		utils.PanicIfErrorExist(err)

		issue, err := finder.FilterIssues(filterableIssues)
		utils.PanicIfErrorExist(err)
		url := issue.GetHTMLURL()
		open.Run(url)
		return
	},
}

func init() {
	browseCmd.AddCommand(browseissuesCmd)
}
