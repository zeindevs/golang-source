package main

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig/v3"
)

var Tmpl *template.Template

func ParseTemplate() error {
	t := template.New("").Funcs(sprig.FuncMap())
	if err := filepath.Walk("templates", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			tmplBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			_, err = t.New(path).Funcs(sprig.FuncMap()).Parse(string(tmplBytes))
			if err != nil {
				return err
			}
		}
		return err
	}); err != nil {
		return err
	}
	Tmpl = t
	return nil
}
