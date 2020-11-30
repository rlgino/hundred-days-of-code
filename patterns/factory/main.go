package main

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/factory/mobileinterface"
	"github.com/rlgino/hundred-days-of-code/patterns/factory/webinterface"
)

func main() {
	factory := mobileinterface.New()
	mobileBtn := factory.CreateButton()
	mobileBtn.Draw()
	fmt.Println("Click ==>")
	mobileBtn.Click()

	fmt.Println()

	factory = webinterface.New()
	webBtn := factory.CreateButton()
	webBtn.Draw()
	fmt.Println("Click ==>")
	webBtn.Click()
}
