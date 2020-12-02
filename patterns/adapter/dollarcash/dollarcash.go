package dollarcash

// DollarSavingsBox is the dollar cash
type DollarSavingsBox interface {
	AddCash(amount int64)
	SubstractCash(amount int64)
	GetCash() int64
}

// New constructor
func New(initialAmount int64) DollarSavingsBox {
	return &dollarSavingsBox{
		cash: initialAmount,
	}
}

type dollarSavingsBox struct {
	cash int64
}

func (box *dollarSavingsBox) AddCash(amount int64) {
	box.cash += amount
}

func (box *dollarSavingsBox) SubstractCash(amount int64) {
	box.cash -= amount
}

func (box *dollarSavingsBox) GetCash() int64 {
	return box.cash
}
