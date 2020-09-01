package genie

import (
	"fmt"
	"io/ioutil"

	"github.com/dave/jennifer/jen"
)

type Output interface {
	Write(j *jen.File) error
}

type FileOutput struct {
	DestinationPath string
}

func NewFileOutput(destinationPath string) *FileOutput {
	return &FileOutput{
		DestinationPath: destinationPath,
	}
}

func (f *FileOutput) Write(j *jen.File) error {
	return ioutil.WriteFile(f.DestinationPath, []byte(fmt.Sprintf("%#v", j)), 0644)
}
