package agregation

import main "github.com/rlgino/hundred-days-of-code/patterns/bridge/interfaces"

// agregation is the implementation of agregate
type agregation struct{}

// New substraction constructor
func New() main.Operation {
	return &agregation{}
}

// Calculate is the implementation of the Agregation
func (a agregation) Calculate(n1, n2 int) int {
	return n1 + n2
}
