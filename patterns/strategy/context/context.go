package context

import "github.com/rlgino/hundred-days-of-code/patterns/strategy/sortstrategy"

// Context of the solution
type Context struct {
	arr      []int
	strategy sortstrategy.Sortstrategy
}

// NewContext constructor
func NewContext(arr []int, strategy sortstrategy.Sortstrategy) Context {
	return Context{
		arr:      arr,
		strategy: strategy,
	}
}

// Sort it's the method executed for the context
func (c Context) Sort() []int {
	return c.strategy.Sort(c.arr)
}
