package main

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/factory/mobileinterface"
	"github.com/rlgino/hundred-days-of-code/patterns/factory/webinterface"
)

func main() {
	factory := mobileinterface.New()
	btnMobile := factory.CreateButton()
	btnMobile.Draw()
	fmt.Println("Click ==>")
	btnMobile.Click()

	fmt.Println()

	factory = webinterface.New()
	webBtn := factory.CreateButton()
	webBtn.Draw()
	fmt.Println("Click ==>")
	webBtn.Click()
}
