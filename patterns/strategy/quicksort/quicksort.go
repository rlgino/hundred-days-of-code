package quicksort

import (
	"fmt"
	"sort"

	"github.com/rlgino/hundred-days-of-code/patterns/strategy/sortstrategy"
)

type quickSort struct{}

// New constructor of buble sort
func New() sortstrategy.Sortstrategy {
	return &quickSort{}
}

func (quickSort) Sort(arr []int) []int {
	fmt.Println("Quick sort!")
	// it should be the quick sort but it isn't mind
	sort.Ints(arr)
	return arr
}
