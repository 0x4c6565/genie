package genie

import "github.com/dave/jennifer/jen"

type Generator interface {
	MarkerName() string
	Generate(marker Marker, typeName string, j *jen.File) error
}

type BaseGenerator struct {
	markerName string
}

func NewBaseGenerator(markerName string) *BaseGenerator {
	return &BaseGenerator{
		markerName: markerName,
	}
}

func (g *BaseGenerator) MarkerName() string {
	return g.markerName
}
