package bublesort

import (
	"fmt"
	"sort"

	"github.com/rlgino/hundred-days-of-code/patterns/strategy/sortstrategy"
)

type bubleSort struct{}

// New constructor of buble sort
func New() sortstrategy.Sortstrategy {
	return &bubleSort{}
}

func (bubleSort) Sort(arr []int) []int {
	fmt.Println("Buble sort!")
	// it should be the buble sort but it isn't mind
	sort.Ints(arr)
	return arr
}
