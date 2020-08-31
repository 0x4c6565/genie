package genie

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Input interface {
	Read() ([]*ast.File, error)
}

type FileInput struct {
	path string
}

func (f *FileInput) Read() ([]*ast.File, error) {
	parsedFile, err := parser.ParseFile(token.NewFileSet(), f.path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return []*ast.File{parsedFile}, nil
}

type DirectoryInput struct {
	path string
}

func (f *DirectoryInput) Read() ([]*ast.File, error) {
	d, err := parser.ParseDir(token.NewFileSet(), f.path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var parsedFiles []*ast.File
	for _, p := range d {
		for _, parsedFile := range p.Files {
			parsedFiles = append(parsedFiles, parsedFile)
		}
	}

	return parsedFiles, nil
}
