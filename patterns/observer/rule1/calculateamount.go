package rule1

import (
	"errors"

	"github.com/rlgino/hundred-days-of-code/patterns/observer/model"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/observer"
)

type amountCalculator struct{}

// New amountCalculator
func New() observer.SaleObserver {
	return &amountCalculator{}
}

// ExecutePreSale calculate total amount of the sale
func (c amountCalculator) ExecutePreSale(sale *model.Sale) error {
	if len(sale.Amounts) == 0 {
		sale.Status = model.CANCELED
		return errors.New("Empty amounts")
	}

	tot := 0
	for _, am := range sale.Amounts {
		tot += am
	}
	sale.TotalAmount = tot

	return nil
}

// ExecutePostSale unimplemented
func (c amountCalculator) ExecutePostSale(sale *model.Sale) error {
	return nil
}
