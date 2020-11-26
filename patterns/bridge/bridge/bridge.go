package bridge

import (
	main "github.com/rlgino/hundred-days-of-code/patterns/bridge/interfaces"
)

// AbstractionBridge is the abstraction
type abstractionBridge struct {
	operation main.Operation
}

// New constructor
func New(operation main.Operation) main.Abstraction {
	return &abstractionBridge{
		operation: operation,
	}
}

func (abs *abstractionBridge) Calculate(n1, n2 int) int {
	return abs.operation.Calculate(n1, n2)
}
