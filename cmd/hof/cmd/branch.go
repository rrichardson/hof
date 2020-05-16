package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"
)

var branchLong = `List, create, or delete branches`

func BranchRun(args []string) (err error) {

	return err
}

var BranchCmd = &cobra.Command{

	Use: "branch",

	Short: "List, create, or delete branches",

	Long: branchLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = BranchRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
