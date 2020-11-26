package rule2

import (
	"errors"

	"github.com/rlgino/hundred-days-of-code/patterns/observer/model"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/observer"
)

type statusChanging struct{}

// New amountCalculator
func New() observer.SaleObserver {
	return &statusChanging{}
}

// ExecutePreSale calculate total amount of the sale
func (c statusChanging) ExecutePreSale(sale *model.Sale) error {
	if sale.Status != model.PENDING {
		sale.Status = model.CANCELED
		return errors.New("Error in status")
	}

	sale.Status = model.FINISH

	return nil
}

// ExecutePostSale unimplemented
func (c statusChanging) ExecutePostSale(sale *model.Sale) error {
	return nil
}
