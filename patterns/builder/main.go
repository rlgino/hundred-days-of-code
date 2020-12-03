package main

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/builder/personbuilder"
)

func main() {
	builder := personbuilder.New()
	builder.WithName("Pepe")
	p := builder.Build()
	fmt.Println(p.String())
}
