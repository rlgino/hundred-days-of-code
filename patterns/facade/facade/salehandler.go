package facade

import (
	"github.com/rlgino/hundred-days-of-code/patterns/facade/billing"
	"github.com/rlgino/hundred-days-of-code/patterns/facade/entities"
	"github.com/rlgino/hundred-days-of-code/patterns/facade/shipping"
	"github.com/rlgino/hundred-days-of-code/patterns/facade/stock"
)

// SaleHandler change the differents status of the finish
type SaleHandler interface {
	Finish(entities.Sale) error
}

// New constructor
func New() SaleHandler {
	return &saleHandler{
		stock.Controller{},
		shipping.Service{},
		billing.Service{},
	}
}

type saleHandler struct {
	stock    stock.Controller
	shipping shipping.Service
	billing  billing.Service
}

func (s saleHandler) Finish(sale entities.Sale) error {
	for _, prod := range sale.ProductsList {
		err := s.stock.StockHandler(prod)
		if err != nil {
			return err
		}
	}

	err := s.shipping.SendProducts(sale.ProductsList)
	if err != nil {
		return err
	}

	err = s.billing.Billing(sale.ID)
	if err != nil {
		return err
	}

	return nil
}
