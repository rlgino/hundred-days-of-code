package main

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/adapter/adapter"
)

func main() {
	savingBox := adapter.New(0)
	savingBox.AddARS(89)
	savingBox.AddARS(89)

	fmt.Println(savingBox.GetAmount())
}
