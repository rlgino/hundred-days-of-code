package interfaces

// DialogFactory factory of dialog object
type DialogFactory interface {
	CreateButton() Button
}

// Button product to render
type Button interface {
	Click()
	Draw()
}
