# Genie

Genie is a tool for generating Go code, inspired by [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder).

At a high-level, Genie uses `Generators` which act upon annotated Go types, using `markers` (see components below)

## Components

**Marker**

A marker consists of a prefix (`genie`), name (`somemarker`) and optional arguments (`someargs=abcdef`)

Example:

```go
// +genie:somemarker:someargs=abcdef
type SomeStruct struct {
	SomeProperty string
}
```

**Generator**

Generators define a marker name upon which they act, and are registered in Genie. Generators utilise [jennifer](https://github.com/dave/jennifer) to generate code

Example:

```go
type SomeGenerator struct {
	*genie.BaseGenerator
}

func NewSomeMarkerGenerator() *SomeGenerator {
	return &SomeGenerator{
		BaseGenerator: genie.NewBaseGenerator("somemarker"),
	}
}

func (g *SomeGenerator) Generate(marker genie.Marker, typeName string, j *jen.File) error {
	j.Type().Id("Generated" + typeName).Struct(
		jen.Id("GeneratedProperty1").Float32(),
	)
	return nil
}
```

**Input**

Inputs implement the `Input` interface, which are responsible for parsing Go code and returning a slice of `ast.File`. Genie provides two basic implementations:

* `FileInput`: Reads and parses code from a file
* `DirectoryInput`: Reads and parses code from a directory

**Output**

Outputs implement the `Output` interface, which are responsible for outputting `jen.File`, typically to files. Genie provides one basic implementation:

* `FileOutput`: Outputs generated code to a file

## Example

Using the components above, we'll create a package to be invoked by go:generate to generate code from struct markers. In this example, we're going to generate an API response struct for an annotated model

**model.go**

This file will contain our annotated model struct and go:generate comment:

```go
//go:generate go run generate/model_apiresponse.go
package main

// +genie:apiresponse
type Person struct {
	Name string
}
```

**generate/model_apiresponse.go**

This file will contain a single generator to act upon the `apiresponse` marker, and a `main` func to invoke Genie with this generator (note, Genie can act upon multiple generators at a time)

```go
package main

import (
	"log"

	"github.com/0x4c6565/genie"
	"github.com/dave/jennifer/jen"
)

type APIResponseGenerator struct {
	*genie.BaseGenerator
}

func NewAPIResponseGenerator() *APIResponseGenerator {
	return &APIResponseGenerator{
		BaseGenerator: genie.NewBaseGenerator("apiresponse"),
	}
}

func (g *APIResponseGenerator) Generate(marker genie.Marker, typeName string, j *jen.File) error {
	j.Type().Id(typeName + "APIResponse").Struct(
		jen.Id("Data").Id(typeName),
	)
	return nil
}

func main() {
	err := genie.NewGenie(
		genie.NewFileInput("model.go"),
		genie.NewFileOutput("model_apiresponse_generated.go"),
		"main",
		NewAPIResponseGenerator(),
	).Generate()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Finished")
}
```

We can then proceed to run the generator:

> go generate ./...

This will render the following into `model_apiresponse_generated.go`:

```go
package main

type PersonAPIResponse struct {
	Data Person
}
```
