package rule3

import (
	"errors"
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/observer/model"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/observer"
)

type billing struct{}

// New amountCalculator
func New() observer.SaleObserver {
	return &billing{}
}

// ExecutePreSale unimplemented
func (c billing) ExecutePreSale(sale *model.Sale) error {
	return nil
}

// ExecutePostSale billing with an external service ie.
func (c billing) ExecutePostSale(sale *model.Sale) error {
	if sale.TotalAmount == 0 || sale.TotalAmount < 0 {
		return errors.New("Invalid amounts")
	}
	if sale.Status != model.FINISH {
		return errors.New("Invalid status")
	}

	fmt.Println("Billing....")
	sale.ID = 123
	fmt.Println("Billed!")
	return nil
}
