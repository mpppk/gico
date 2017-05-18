package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mpppk/gico/finder"
	"github.com/skratchdot/open-golang/open"
	"github.com/mpppk/gico/utils"
	"github.com/mpppk/hlb/hlb"
)

// browsepullrequestsCmd represents the browsepullrequests command
var browsepullrequestsCmd = &cobra.Command{
	Use:   "pullrequests",
	Short: "browse pulrequests",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		base, err := hlb.NewCmdBase()
		utils.PanicIfErrorExist(err)
		sw := finder.NewServiceWrapper(base)

		filterablePrs, err := sw.GetFilterablePullRequests()
		utils.PanicIfErrorExist(err)

		pr, err := finder.FilterPullRequests(filterablePrs)
		url := pr.GetHTMLURL()
		open.Run(url)
		return
	},
}

func init() {
	browseCmd.AddCommand(browsepullrequestsCmd)
}
