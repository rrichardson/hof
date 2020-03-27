package schema


// Definition for a generator
HofGen :: {
  // Base directory for the output
  Outdir: string | *"./"

  // "Global" input, merged with out replacing onto the files
  In: { ... } | * {...}

  // The list fo files for hof to generate
  Out: [...GenFile] | *[...]

  // Subgenerators for composition
  Generators: [...HofGen] | {...}

  //
  // For file based generators
  //

  // Used for indexing into the vendor directory...
  PackageName: string | * ""

  // Base directory of entrypoint templates to load
  TemplatesDir: string | * ""
  // Base directory of partial templatess to load
  PartialsDir: string | * ""
  // Filepath globs for static files to load
  StaticGlobs: [...string] | * ""

  // Open for whatever else you may need
  //   often hidden fields are used
  ...
} 

// A file which should be generated by hof
GenFile :: {
  // The local input data
  In: { ... }

  // The full path to the output location
  Filename: string

  //
  // Template parameters
  //

  // System params
  TemplateSystem: *"text/template" | "mustache"

  // TODO delimiters
  // [double, triple delims] X [LHS, RHS] X [MAIN, SWAP, TEMP]

  // Swap delimiters, becuase the template systems use `{{` and `}}`
  //   and if you want to preserve those, we need three sets of delimiters
  Alt: bool | *false
  SwapDelims: bool
  SwapDelims: Alt

  // The template contents
  Template: string | *""

  // Relative name from TemplatesDir
  TemplateFile: string | *""

  // Open for whatever else you may need, not sure about this case though
  ...
}