package genie

import (
	"fmt"
	"io/ioutil"

	"github.com/dave/jennifer/jen"
)

type Output interface {
	Output(j *jen.File) error
}

type FileOutput struct {
	DestinationPath string
}

func NewFileOutput(destinationPath string) *FileOutput {
	return &FileOutput{
		DestinationPath: destinationPath,
	}
}

func (f *FileOutput) Output(j *jen.File) error {
	return ioutil.WriteFile(f.DestinationPath, []byte(fmt.Sprintf("%#v", j)), 0644)
}
