package cmd

import (
	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
	"github.com/mpppk/gico/utils"
	"github.com/mpppk/hlb/hlb"
	"github.com/mpppk/gico/finder"
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
		sw := finder.NewServiceWrapper(base)

		url, err := sw.GetRepositoryURL()
		utils.PanicIfErrorExist(err)
		open.Run(url)
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
}
