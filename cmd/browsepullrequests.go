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
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		base, err := hlb.NewCmdBase()
		utils.PanicIfErrorExist(err)
		sw := finder.NewServiceWrapper(base)

		filterablePrs, err := sw.GetFilterablePullRequests()
		utils.PanicIfErrorExist(err)

		var filterableStrings []finder.FilterableStringer
		for _, fis := range filterablePrs {
			filterableStrings = append(filterableStrings, finder.FilterableStringer(fis))
		}

		res, err := finder.Filter(filterableStrings)
		utils.PanicIfErrorExist(err)
		pr := res.(*finder.FilterablePullRequest)
		url := pr.GetHTMLURL()
		open.Run(url)
		return
	},
}

func init() {
	browseCmd.AddCommand(browsepullrequestsCmd)
}
