package launch

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
)

const (
	defaultExt = ".html"
)

type Renderer *ViewRenderer

type partialsProvider struct {
	ext  string
	root string
}

func (p partialsProvider) Get(name string) (string, error) {
	return name, nil
}

type ViewRenderer struct {
	root             string
	ext              string
	partialsProvider partialsProvider
	loaded           bool
	reload           bool
	templates        map[string]*mustache.Template
}

func NewViewRenderer(root string) *ViewRenderer {
	r := &ViewRenderer{
		root: root,
		ext:  ".html",
		partialsProvider: partialsProvider{
			ext:  defaultExt,
			root: root,
		},
		reload: true,
	}
	if err := r.load(); err != nil {
		log.Fatal(err)
	}
	return r
}

func (r *ViewRenderer) Render(w io.Writer, name string, data any) error {
	if !r.loaded || r.reload {
		if r.loaded {
			r.loaded = true
		}
		if err := r.load(); err != nil {
			return err
		}
	}
	tmpl, ok := r.templates[name]
	if !ok {
		return fmt.Errorf("view <%s> not found", name)
	}
	return tmpl.FRender(w, data)
}

func (r *ViewRenderer) load() error {
	r.templates = make(map[string]*mustache.Template)
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info == nil || info.IsDir() {
			return nil
		}
		if len(r.ext) >= len(path) || path[len(path)-len(r.ext):] != r.ext {
			return nil
		}
		rel, err := filepath.Rel(r.root, path)
		if err != nil {
			return err
		}
		name := filepath.ToSlash(rel)
		name = strings.TrimSuffix(name, r.ext)
		buf, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var tmpl *mustache.Template
		tmpl, err = mustache.ParseStringPartials(string(buf), r.partialsProvider)
		if err != nil {
			return fmt.Errorf("template view error %s", err.Error())
		}
		r.templates[name] = tmpl
		return err
	}
	r.loaded = true
	return filepath.Walk(r.root, walkFn)
}
