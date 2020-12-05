package shipping

import (
	"errors"
	"fmt"
)

// Service of shipping in other API
type Service struct{}

// SendProducts to the shipping API
func (s Service) SendProducts(products []string) error {
	quantity := len(products)
	if quantity == 0 {
		return errors.New("Product list is empty")
	}
	fmt.Println(fmt.Sprintf("Send %d products to address", quantity))
	return nil
}
