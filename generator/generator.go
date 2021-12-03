package generator

import (
	"strings"

	"github.com/go-generalize/go-easyparser"
	"github.com/go-generalize/go-easyparser/types"
	"golang.org/x/xerrors"
)

// Generator generates firestore CRUD functions
type Generator struct {
	dir   string
	types map[string]types.Type

	AppVersion string
}

func NewGenerator(dir string) (*Generator, error) {
	psr, err := easyparser.NewParser(dir, func(fo *easyparser.FilterOpt) bool {
		return fo.BasePackage
	})

	if err != nil {
		return nil, xerrors.Errorf("failed to initialize parser: %w", err)
	}
	psr.Replacer = replacer

	types, err := psr.Parse()

	if err != nil {
		return nil, xerrors.Errorf("failed to parse: %w", err)
	}

	return &Generator{
		dir:   dir,
		types: types,

		AppVersion: "devel",
	}, nil
}

// GenerateOption is a parameter to generate repository
type GenerateOption struct {
	OutputDir      string
	PackageName    string
	MockGenPath    string
	MockOutputPath string
	UseMetaField   bool
	Subcollection  bool
}

// NewDefaultGenerateOption returns a default GenerateOption
func NewDefaultGenerateOption() GenerateOption {
	return GenerateOption{
		OutputDir:      ".",
		MockGenPath:    "mockgen",
		MockOutputPath: "mock/mock_{{ .GeneratedFileName }}/mock_{{ .GeneratedFileName }}.go",
		UseMetaField:   true,
	}
}

func (g *Generator) Generate(structName string, opt GenerateOption) error {
	var typ *types.Object
	for k, v := range g.types {
		split := strings.Split(k, ".")
		t := split[len(split)-1]

		if t == structName {
			t, ok := v.(*types.Object)

			if !ok {
				return xerrors.Errorf("Only struct is allowed")
			}
			typ = t
		}
	}

	if typ == nil {
		return xerrors.Errorf("struct not found: %s", structName)
	}

	gen, err := newStructGenerator(typ, structName, g.AppVersion, opt)

	if err != nil {
		return xerrors.Errorf("failed to initialize generator: %w", err)
	}

	if err := gen.parseType(); err != nil {
		return xerrors.Errorf("failed to parse type: %w", err)
	}

	if err := gen.generate(); err != nil {
		return xerrors.Errorf("failed to generate files: %w", err)
	}

	return nil
}
