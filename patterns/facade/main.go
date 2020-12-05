package main

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/facade/entities"
	"github.com/rlgino/hundred-days-of-code/patterns/facade/facade"
)

func main() {
	s := entities.Sale{
		ID:           12,
		ProductsList: []string{"TV", "Computer", "Table"},
	}

	handler := facade.New()
	err := handler.Finish(s)
	if err != nil {
		panic(fmt.Sprintf("Oh no! The sale  finishing has failed: %s", err.Error()))
	}
}
