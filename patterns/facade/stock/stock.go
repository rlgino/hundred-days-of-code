package stock

import (
	"errors"
	"fmt"
)

// Controller struct
type Controller struct{}

// StockHandler add or substract stock of differents products
func (c Controller) StockHandler(prod string) error {
	if prod == "" {
		return errors.New("The product is invalid")
	}
	fmt.Println("Handler stock of", prod)
	return nil
}
