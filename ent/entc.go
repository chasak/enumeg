// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ariga/ogent"
	"github.com/ogen-go/ogen"
)

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(entoas.Spec(spec))
	// oas, err := entoas.NewExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	if err := entc.Generate("../schema",
		&gen.Config{
			Target:  "../entgen",
			Package: "example.com/enumeg/entgen",
		},
		entc.Extensions(ogent, oas)); err != nil {
		// entc.Extensions(oas)); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
