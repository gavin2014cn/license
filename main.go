package main

import (
	"flag"
	def "license/define"
	lic "license/funcs"
	"os"
	"strings"
)

func main() {

	flag.StringVar(&def.Args.FN, "name", "", "name on license")
	flag.StringVar(&def.Args.FN, "n", "", "name on license")
	flag.StringVar(&def.Args.FY, "year", "", "year on license")
	flag.StringVar(&def.Args.FY, "y", "", "year on license")
	flag.StringVar(&def.Args.FO, "output", "", "path to output file")
	flag.StringVar(&def.Args.FO, "o", "", "path to output file")
	flag.BoolVar(&def.Args.FV, "version", false, "print version")
	flag.BoolVar(&def.Args.FH, "help", false, "print help")
	flag.BoolVar(&def.Args.FL, "list", false, "print available licenses")

	flag.Usage = func() {
		lic.Pusage()
		os.Exit(1)
	}
	flag.Parse()

	run()
}

func run() {
	if flag.NArg() != 1 && !(def.Args.FV || def.Args.FH || def.Args.FL) {
		lic.Pusage()
		os.Exit(1)
	}

	switch {
	case def.Args.FV:
		lic.Pver()
		os.Exit(0)

	case def.Args.FH:
		lic.Pusage()
		os.Exit(0)

	case def.Args.FL:
		lic.Plist()
		os.Exit(0)

	default:
		license := strings.ToLower(flag.Arg(0))
		lic.Plic(license) // internally calls os.Exit() on failure
	}
}
