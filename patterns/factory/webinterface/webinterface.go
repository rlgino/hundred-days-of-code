package webinterface

import (
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/factory/interfaces"
)

type webInterface struct{}

// New constructor to web interface
func New() interfaces.DialogFactory {
	return &webInterface{}
}

// CreateButton creates the button fo web interface
func (webInterface) CreateButton() interfaces.Button {
	return &webButton{}
}

type webButton struct{}

func (btn *webButton) Draw() {
	fmt.Println("<input type=\"button\">Saludar!</>")
}

func (btn *webButton) Click() {
	fmt.Println("Hola mundo!")
}
