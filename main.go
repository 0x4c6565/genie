package genie

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Marker struct {
	Name string
	Args string
}

type Genie struct {
	markerGenerators []Generator
	markerPrefix     string
	input            Input
	output           Output
	pkg              string
}

func NewGenie(input Input, output Output, pkg string, generators ...Generator) *Genie {
	return &Genie{
		markerPrefix:     "genie",
		input:            input,
		output:           output,
		pkg:              pkg,
		markerGenerators: generators,
	}
}

func (g *Genie) WithPrefix(p string) *Genie {
	g.markerPrefix = p
	return g
}

func (g *Genie) WithGenerator(generators ...Generator) *Genie {
	g.markerGenerators = append(g.markerGenerators, generators...)
	return g
}

func (g *Genie) Generate() error {
	jenFile := jen.NewFile(g.pkg)

	files, err := g.input.Read()
	if err == nil {
		return err
	}
	for _, file := range files {
		g.generateFile(jenFile, file)
	}

	g.output.Output(jenFile)

	return nil
}

func (g *Genie) generateFile(jenFile *jen.File, parsedFile *ast.File) error {
	var innerErr error
	ast.Inspect(parsedFile, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			for _, comment := range strings.Split(x.Doc.Text(), "\n") {
				if !g.isMarker(comment) {
					continue
				}

				for _, spec := range x.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}

					marker, err := g.parseMarker(comment)
					if err != nil {
						innerErr = err
						return false
					}

					generator := g.getGenerator(marker.Name)
					if generator == nil {
						continue
					}

					err = generator.Generate(marker, typeSpec.Name.String(), jenFile)
					if err != nil {
						innerErr = fmt.Errorf("Marker %s generate failed: %s", generator.MarkerName(), err.Error())
						return false
					}
				}
			}
		}

		return true
	})

	return innerErr
}

func (g *Genie) getGenerator(marker string) Generator {
	for _, generator := range g.markerGenerators {
		if generator.MarkerName() == marker {
			return generator
		}
	}

	return nil
}

func (g *Genie) parseMarker(marker string) (Marker, error) {
	m := Marker{}

	parts := strings.Split(marker, ":")
	if len(parts) < 2 {
		return m, fmt.Errorf("Invalid marker '%s'", marker)
	}

	m.Name = parts[1]
	if len(parts) == 3 {
		m.Args = parts[2]
	}

	return m, nil
}

func (g *Genie) isMarker(s string) bool {
	return strings.HasPrefix(s, "+"+g.markerPrefix+":")
}
