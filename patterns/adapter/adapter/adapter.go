package adapter

import "github.com/rlgino/hundred-days-of-code/patterns/adapter/dollarcash"

const (
	dollarPrice = 89
)

// SavingBoxAdapter the ARS saving box
type SavingBoxAdapter interface {
	AddARS(amount int64)
	SubstractARS(amount int64)
	GetAmount() int64
}

// New adapter constructor
func New(initialAmount int64) SavingBoxAdapter {
	return &adapterARS{
		dollarcash.New(initialAmount),
	}
}

type adapterARS struct {
	dollarCash dollarcash.DollarSavingsBox
}

func (adapter *adapterARS) AddARS(amount int64) {
	total := amount / dollarPrice
	adapter.dollarCash.AddCash(total)
}

func (adapter *adapterARS) SubstractARS(amount int64) {
	total := amount / dollarPrice
	adapter.dollarCash.SubstractCash(total)

}

func (adapter *adapterARS) GetAmount() int64 {
	return adapter.dollarCash.GetCash() * dollarPrice
}
