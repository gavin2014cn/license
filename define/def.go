package define

type Arg struct {
	FN string
	FY string
	FO string
	FV bool
	FH bool
	FL bool
}

var Args Arg

const (
	NameEnv       = "LICENSE_FULL_NAME"
	VersionString = "v0.1"

	UsageString = `Usage: license [flags] [license-type]

Flags:
       -help     print help information
       -list     print list of available license types
   -n, -name     full name to use on license (default %q)
   -o, -output   path to output file (prints to stdout if unspecified)
       -version  print version
   -y, -year     year to use on license (default %q)

Examples:
  license mit
  license -name "Gavin" -year 2025 -o LICENSE bsd-3-clause
  license -o LICENSE mpl-2.0`
)
