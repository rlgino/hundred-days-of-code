package mobileinterface

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/factory/interfaces"
)

type mobileInterface struct{}

// New constructor to mobile interface
func New() interfaces.DialogFactory {
	return &mobileInterface{}
}

// CreateButton creates the mobile button
func (mobile *mobileInterface) CreateButton() interfaces.Button {
	return &mobileButton{}
}

type mobileButton struct {
	wasDrawed bool
}

func (btn *mobileButton) Draw() {
	fmt.Println("----------------")
	fmt.Println("|---Greeting---|")
	fmt.Println("----------------")
	btn.wasDrawed = true
}

func (btn *mobileButton) Click() {
	if !btn.wasDrawed {
		panic("The button wasn't drawed")
	}

	fmt.Println("Hello World!")
}
