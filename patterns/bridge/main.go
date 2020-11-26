package main

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/bridge/agregation"
	"github.com/rlgino/hundred-days-of-code/patterns/bridge/bridge"
	"github.com/rlgino/hundred-days-of-code/patterns/bridge/substraction"
)

func main() {
	n1 := 1
	n2 := 2

	agregationImpl := bridge.New(agregation.New())
	result := agregationImpl.Calculate(n1, n2)
	fmt.Println(fmt.Sprintf("The agregation resultmof %d and %d is: %d", n1, n2, result))

	substractionImpl := bridge.New(substraction.New())
	result = substractionImpl.Calculate(n1, n2)
	fmt.Println(fmt.Sprintf("The substraction result of %d and %d is: %d", n1, n2, result))
}
