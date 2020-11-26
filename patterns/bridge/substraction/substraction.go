package substraction

import (
	main "github.com/rlgino/hundred-days-of-code/patterns/bridge/interfaces"
)

// substraction it's the substraction implementation
type substraction struct{}

// New substraction constructor
func New() main.Operation {
	return &substraction{}
}

// Calculate it's the calculate of the substraction implementation
func (r substraction) Calculate(n1, n2 int) int {
	return n1 - n2
}
