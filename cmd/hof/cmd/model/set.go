package cmdmodel

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/cmd/model/set"

	"github.com/hofstadter-io/hof/cmd/hof/ga"
)

var setLong = `create, view, migrate, and understand your data model sets.`

var SetCmd = &cobra.Command{

	Use: "set",

	Short: "create, view, migrate, and understand your data model sets.",

	Long: setLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},
}

func init() {
	hf := SetCmd.HelpFunc()
	f := func(cmd *cobra.Command, args []string) {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/help", "<omit>", 0)
		hf(cmd, args)
	}
	SetCmd.SetHelpFunc(f)
	SetCmd.AddCommand(cmdset.CreateCmd)
	SetCmd.AddCommand(cmdset.ViewCmd)
	SetCmd.AddCommand(cmdset.ListCmd)
	SetCmd.AddCommand(cmdset.StatusCmd)
	SetCmd.AddCommand(cmdset.GraphCmd)
	SetCmd.AddCommand(cmdset.DiffCmd)
	SetCmd.AddCommand(cmdset.MigrateCmd)
	SetCmd.AddCommand(cmdset.TestCmd)
	SetCmd.AddCommand(cmdset.DeleteCmd)
}