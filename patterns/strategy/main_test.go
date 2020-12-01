package main_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/strategy/bublesort"
	"github.com/rlgino/hundred-days-of-code/patterns/strategy/context"
	"github.com/rlgino/hundred-days-of-code/patterns/strategy/quicksort"
	"github.com/rlgino/hundred-days-of-code/patterns/strategy/sortstrategy"
)

func Test_strategy(t *testing.T) {
	tests := []struct {
		name         string
		sortStrategy sortstrategy.Sortstrategy
	}{
		{
			name:         "Buble sort",
			sortStrategy: bublesort.New(),
		},
		{
			name:         "Quick sort",
			sortStrategy: quicksort.New(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := []int{3, 2, 1, 0}
			ctx := context.NewContext(arr, test.sortStrategy)
			result := ctx.Sort()

			for i, el := range result {
				if i != el {
					t.Error(el, " should be equal ", i)
				}
			}
		})
	}
}
