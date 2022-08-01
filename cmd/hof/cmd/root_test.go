package cmd_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/script/runtime"

	"github.com/hofstadter-io/hof/cmd/hof/cmd"
)

func init() {
	// ensure our root command is setup
	cmd.RootInit()
}

func TestScriptRootCliTests(t *testing.T) {
	// setup some directories
	workdir := ".workdir/cli/root"
	yagu.Mkdir(workdir)

	runtime.Run(t, runtime.Params{
		Setup: func(env *runtime.Env) error {
			// add any environment variables for your tests here

			return nil
		},
		Funcs: map[string]func(ts *runtime.Script, args []string) error{
			"__hof": cmd.CallTS,
		},
		Dir:         "hls/cli/root",
		WorkdirRoot: workdir,
	})
}