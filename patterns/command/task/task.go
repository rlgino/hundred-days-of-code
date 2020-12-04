package task

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/command/command"
)

// Name of different tasks
type Name string

const (
	// Clear task
	Clear Name = "clear"
	// Wash task
	Wash Name = "wash"
	// Dry task
	Dry Name = "dry"
)

type task struct {
	callback command.Command
	name     Name
}

// New task
func New(name Name, callback command.Command) command.Command {
	return &task{
		name:     name,
		callback: callback,
	}
}

func (agg task) Execute() {
	fmt.Println("Executing task", agg.name)
	if agg.callback != nil {
		agg.callback.Execute()
	}
}
