package gocodegen

import (
	"bytes"
	"go/format"
	"os"
	"text/template"

	"golang.org/x/xerrors"
)

// GoCodeGenerator generates Go code and formats it
type GoCodeGenerator struct {
	tmpl *template.Template
}

// NewGoCodeGenerator returns a new GoCodeGenerator
func NewGoCodeGenerator(tmpl *template.Template) *GoCodeGenerator {
	return &GoCodeGenerator{
		tmpl: tmpl,
	}
}

// Generate executes the template and formats it
func (g *GoCodeGenerator) Generate(name string, param interface{}) ([]byte, error) {
	buf := bytes.Buffer{}

	if err := g.tmpl.ExecuteTemplate(&buf, name, param); err != nil {
		return nil, xerrors.Errorf("failed to execte template of %s: %w", name, err)
	}

	b, err := format.Source(buf.Bytes())

	if err != nil {
		return buf.Bytes(), xerrors.Errorf("failed to format source: %w", err)
	}

	return b, nil
}

// GenerateTo is similar to Generate, but the result is written into the file with the specified path
func (g *GoCodeGenerator) GenerateTo(name string, param interface{}, path string) error {
	b, err := g.Generate(name, param)

	if err != nil {
		_ = os.WriteFile(path, b, 0774)

		return xerrors.Errorf("failed to generate code: %w", err)
	}

	err = os.WriteFile(path, b, 0774)

	if err != nil {
		return xerrors.Errorf("failed to write into %s: %w", path, err)
	}

	return nil
}
