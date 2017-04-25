// Code generated with goagen v2.0.0-wip, DO NOT EDIT.
//
// Code Generator
//
// Command:
// $ goagen

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/codeclysm/genangular"
	"goa.design/goa.v2/codegen"
	"goa.design/goa.v2/eval"
	_ "goa.design/goa.v2/examples/account/design"
)

func main() {
	var (
		out = flag.String("output", "", "")
	)
	{
		flag.Parse()
		if *out == "" {
			fail("missing output flag")
		}
	}
	if err := eval.Context.Errors; err != nil {
		fail(err.Error())
	}
	if err := eval.RunDSL(); err != nil {
		fail(err.Error())
	}

	var roots []eval.Root
	{
		rs, err := eval.Context.Roots()
		if err != nil {
			fail(err.Error())
		}
		roots = rs
	}

	var genfiles []codegen.File
	{
		fs, err := genangular.Generate(roots...)
		if err != nil {
			fail(err.Error())
		}
		genfiles = append(genfiles, fs...)

		// Delete previously generated directories
		dirs := make(map[string]bool)
		for _, f := range genfiles {
			dirs[filepath.Dir(filepath.Join("gen", f.OutputPath()))] = true
		}
		for d := range dirs {
			if _, err := os.Stat(d); err == nil {
				if err := os.RemoveAll(d); err != nil {
					fail(err.Error())
				}
			}
		}
	}

	var w *codegen.Writer
	{
		w = &codegen.Writer{
			Dir:   *out,
			Files: make(map[string]bool),
		}
	}
	for _, f := range genfiles {
		if err := w.Write("gen", f); err != nil {
			fail(err.Error())
		}
	}

	var outputs []string
	{
		outputs = make([]string, len(w.Files))
		cwd, err := os.Getwd()
		if err != nil {
			cwd = "."
		}
		i := 0
		for o := range w.Files {
			rel, err := filepath.Rel(cwd, o)
			if err != nil {
				rel = o
			}
			outputs[i] = rel
			i++
		}
	}
	sort.Strings(outputs)
	fmt.Println(strings.Join(outputs, "\n"))
}

func fail(msg string, vals ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, vals...)
	os.Exit(1)
}
