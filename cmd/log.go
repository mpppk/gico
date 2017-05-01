package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/mpppk/gico/finder"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "show git log",
	Long:  `show git log`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := finder.GetLogHashInteractive()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(logCmd)
}
