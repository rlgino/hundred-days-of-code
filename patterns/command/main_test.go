package main_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/command/task"
)

func Test_commandSuccess(t *testing.T) {
	t1 := task.New(task.Clear, nil)

	if t1 == nil {
		t.Error("t1 should be not nil")
	}

	t1.Execute()
}
