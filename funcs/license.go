package funcs

import (
	"fmt"
	"io"
	"license/cmn"
	"license/define"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var (
	Stdout = log.New(os.Stdout, "", 0)
	Stderr = log.New(os.Stderr, "", 0)
)

//go:generate go-bindata -o templates.go .templates/
//license: the license type you want
//output: the output file address
//name: your license personal name
//year: your license expire year
func Plic(license string) {
	name := cmn.GName()
	year := cmn.GYear()
	output := define.Args.FO
	data, err := Asset(fmt.Sprintf(".templates/%s.tmpl", license))
	if err != nil {
		Stderr.Printf("unknown license %q\nrun \"license -list\" for list of available licenses", license)
		os.Exit(2)
	}

	t, err := template.New("license").Parse(string(data))
	if err != nil {
		Stderr.Printf("internal: failed to parse license template for %s", license)
		os.Exit(1)
	}

	var outFile io.Writer = os.Stdout
	if output != "" {
		f, err := os.Create(filepath.Clean(output))
		if err != nil {
			Stderr.Printf("failed to create file %s: %s", output, err)
			os.Exit(1)
		}
		outFile = f
	}

	if err := t.Execute(outFile, struct {
		Name string
		Year string
	}{name, year}); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
