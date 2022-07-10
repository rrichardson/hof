package gen

import (
	"fmt"
	"sort"
	"strings"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/hof/lib/templates"
)

func (R *Runtime) AsModule() error {
	FP := R.Flagpole
	name := FP.AsModule
	var content string

	if R.Verbosity > 0 {
		fmt.Println("modularizing", name)
		fmt.Println(strings.Repeat("-", 23))
	}

	// parse template flags
	tcfgs  := []AdhocTemplateConfig{}
	tfiles := make([]string,0)
	for _, tf := range R.Flagpole.Template {
		cfg, err := parseTemplateFlag(tf)
		if err != nil {
			return err
		}
		tcfgs  = append(tcfgs, cfg)
		tfiles = append(tfiles, cfg.Filepath)

		if R.Verbosity > 0 {
			fmt.Printf("%#v\n", cfg)
		}
	}

	// top-level fields that would have been accessible
	ins := []string{}
	// but not anything with this attribute
	filters := map[string]bool{
		"gen": true,
		"hof": true,
	}

	// get top-level CUE value as a struct
	S, err := R.CueRuntime.CueValue.Struct()
	if err != nil {
		return err
	}

	// Loop through all top level fields
	// They must be regular by design
	iter := S.Fields()
	for iter.Next() {

		// what we will add if not filtered
		label := iter.Label()

		// let's possibly filster
		value := iter.Value()
		attrs := value.Attributes(cue.ValueAttr)

		filtered := false
		// find top-level with gen attr
		for _, A := range attrs {
			// does it have "@gen()"
			if _, ok := filters[A.Name()]; ok {
				filtered = true
			}
		}

		if !filtered {
			ins = append(ins, label)
		}
	}

	// get generator names that were loaded by -G
	gens := []string{}
	for label, _ := range R.Generators {
		if label == "AdhocGen" {
			continue
		}
		gens = append(gens, label)
	}
	sort.Strings(gens)

	// construct template input data
	data := map[string]interface{}{
		"Name": name,
		"Inputs": ins,
		"Configs": tcfgs,
		"Templates": tfiles,
		"Partials": FP.Partial,
		"Generators": gens,
		"Diff3": FP.Diff3,
		"WatchGlobs": FP.WatchGlobs,
		"WatchXcue": FP.WatchXcue,
	}

	// render template
	ft, err := templates.CreateFromString("as-module", asModuleTemplate, nil)
	if err != nil {
		return err
	}
	bs, err := ft.Render(data)
	if err != nil {
		return err
	}

	// format content (or bs if bytes are better)
	content = string(bs)

	// stdout or write file
	if name == "-" {
		fmt.Println(content)
	} else {
		// TODO write to file
		fmt.Println(content)

		// also write mod file

		// render template
		ft, err := templates.CreateFromString("cue-mods", cuemodsTemplate, nil)
		if err != nil {
			return err
		}
		bs, err := ft.Render(data)
		if err != nil {
			return err
		}

		// format content (or bs if bytes are better)
		content = string(bs)

		// TODO write to 'cue.mods'
		// fmt.Println(content)


		// fmt.Println(finalMsg + name + "\n")
	}

	if R.Verbosity > 0 {
		fmt.Println(strings.Repeat("-", 23))
	}

	return nil
}

const finalMsg = `
Now run
  hof mod vendor cue
  hof gen -G `

const asModuleTemplate = `
package {{ .Name }}

import (
	"github.com/hofstadter-io/hof/schema/gen"
)

// This is example usage of your generator
{{ camelT .Name }}Example: #{{ camelT .Name }}Generator & {
	@gen({{ .Name }})

	// inputs to the generator
	{{ range .Inputs -}}
	"{{.}}": {{.}},
	{{ else }}
	// you almost certainly need to
	// manually add input data here
	// "data": data,
	{{- end }}

	// other settings
	Diff3: {{ .Diff3 }}	
	Outdir: "./"
	
	{{ if .WatchGlobs }}
	// File globs to watch and trigger regen when changed
	// Normally, a user would set this to their designs / datamodel
	WatchGlobs: [ {{ range .WatchGlobs }}"{{.}}", {{ end }} ]
	{{ end }}
	{{ if .WatchXcue }}
	// This is really only useful for module authors
	WatchXcue:  [...string] | *[ {{ range .WatchXcue  }}"{{.}}", {{ end }} ]
	{{ end }}

	// required by examples inside the same module
	// your users do not set or see this field
	PackageName: ""
}


// This is your reusable generator module
//
#{{ camelT .Name }}Generator: gen.#Generator & {

	//
	// user input fields
	//

	// this is the interface for this generator module
	// typically you enforce schema(s) here
	{{ range .Inputs -}}
	{{.}}: _
	{{ else }}
	// you almost certainly need to
	// manually add input fields & schemas here
	// "data": _
	// "input": #Input
	{{- end }}

	//
	// Internal Fields
	//

	// This is the global input data the templates will see
	// You can reshape and transform the user inputs
	// While we put it under internal, you can expose In
	In: {
		// if you want to user your input data
		// add top-level fields from your
		// CUE entrypoints here, adjusting as needed
		// Since you made this a module for others,
		// it won't output until this field is filled

		{{ range .Inputs -}}
		"{{.}}": {{ . }}
		{{ else }}
		// you almost certainly need to
		// manually add input fields & schemas here
		// "data": _
		// "input": #Input
		{{- end }}

		...
	}

	// required for hof CUE modules to work
	// your users do not set or see this field
	PackageName: string | *"hof.io/{{ .Name }}"

	{{ if .Templates -}}
	// Templates: [{Globs: ["./templates/**/*"], TrimPrefix: "./templates/"}]
	Templates: [ { Globs: [ {{ range .Templates }}"{{.}}", {{ end }} ] } ]
	{{ else }}
	// Templates: [{Globs: ["./templates/**/*"], TrimPrefix: "./templates/"}]
	Templates: []
	{{ end }}
	{{ if .Partials -}}
	// Partials: [#Templates & {Globs: ["./partials/**/*"], TrimPrefix: "./partials/"}]
	Partials:  [ { Globs: [ {{ range .Partials  }}"{{.}}", {{ end }} ] } ]
	{{ else }}
	// Partials: [#Templates & {Globs: ["./partials/**/*"], TrimPrefix: "./partials/"}]
	Partials: []
	{{ end }}

	{{ if .Generators -}}
	// these should be in the same CUE package
	// or you may have to manually import as needed
	Generators: {
		{{ range .Generators -}}
		"{{.}}": {{.}},
		{{- end }}
	}
	{{ end }}

	// The final list of files for hof to generate
	Out: [...gen.#File] & [
		{{ range $i, $cfg := .Configs -}}
		{{ if $cfg.Repeated }}for _, t in t_{{ $i }} { t }{{ else }}t_{{ $i }}{{end}},
		{{ end }}
	]

	// These are the -T mappings
	{{ range $i, $cfg := .Configs -}}
	t_{{ $i }}: {{ if not .Repeated }}{
		TemplatePath: "{{ .Filepath }}"
		Filepath:     "{{ .Outpath }}"
	}{{ else }}[{
		TemplatePath: "{{ .Filepath }}"
		Filepath:     "{{ .Outpath }}"
	}]{{ end }}
	{{ end }}

	// so your users can build on this
	...
}
`

const cuemodsTemplate = `
module hof.io/{{ .Name }}

cue v0.4.3

require (
	github.com/hofstadter-io/hof v0.6.3
)
`