package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rlgino/hundred-days-of-code/patterns/strategy/bublesort"
	"github.com/rlgino/hundred-days-of-code/patterns/strategy/context"
	"github.com/rlgino/hundred-days-of-code/patterns/strategy/quicksort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choice your sort algorithm-> ")
	fmt.Println("1- buble ")
	fmt.Println("2- quick ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	arr := []int{3, 2, 1}
	var ctx context.Context

	if text == "1" {
		ctx = context.NewContext(arr, bublesort.New())
	} else if text == "2" {
		ctx = context.NewContext(arr, quicksort.New())
	} else {
		return
	}

	newArr := ctx.Sort()
	fmt.Println(newArr)
}
