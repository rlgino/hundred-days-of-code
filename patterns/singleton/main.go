package main

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/singleton/database"
)

func main() {
	instance := database.GetInstance()
	instance.Select(0)
	fmt.Println("Connection ID: ", instance.GetID())

	newInstance := database.GetInstance()
	newInstance.Select(1)
	fmt.Println("Connection ID:", newInstance.GetID())
}
