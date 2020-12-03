package main_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/builder/personbuilder"
)

func Test_builder(t *testing.T) {
	builder := personbuilder.New()
	p := builder.Build()
	if p.Name != "" {
		t.Error("The name should be empty")
	}

	p1 := builder.WithName("Pepe").Build()
	if p1.Name != "Pepe" {
		t.Error("The name should be Pepe")
	}

	p2 := builder.WithAge(12).Build()
	if p2.Name != "Pepe" && p2.Age != 12 {
		t.Error("The name should be Pepe and the age should be 12")
	}

	p3 := builder.WithAddress("Av Siempre Viva 123").Build()
	if p3.Name != "Pepe" && p3.Age != 12 && p3.Address != "Av Siempre Viva 123" {
		t.Error("The name should be Pepe and the age should be 12 and address Av Siempre Viva 123")
	}
}
