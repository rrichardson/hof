package flow_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/script/runtime"
)

func TestAPIFlow(t *testing.T) {
	yagu.Mkdir(".workdir/tasks/api")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata/tasks/api",
		Glob:        "*.txt",
		WorkdirRoot: ".workdir/tasks/api",
	})
}

func TestGenFlow(t *testing.T) {
	yagu.Mkdir(".workdir/tasks/gen")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata/tasks/gen",
		Glob:        "*.txt",
		WorkdirRoot: ".workdir/tasks/gen",
	})
}

func TestKVFlow(t *testing.T) {
	yagu.Mkdir(".workdir/tasks/kv")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata/tasks/kv",
		Glob:        "*.txt",
		WorkdirRoot: ".workdir/tasks/kv",
	})
}

func TestOSFlow(t *testing.T) {
	yagu.Mkdir(".workdir/tasks/os")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata/tasks/os",
		Glob:        "*.txt",
		WorkdirRoot: ".workdir/tasks/os",
	})
}

func TestStFlow(t *testing.T) {
	yagu.Mkdir(".workdir/tasks/st")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata/tasks/st",
		Glob:        "*.txt",
		WorkdirRoot: ".workdir/tasks/st",
	})
}
