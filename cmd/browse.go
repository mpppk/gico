package cmd

import (
	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
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
