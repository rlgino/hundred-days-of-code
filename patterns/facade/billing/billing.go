package billing

import (
	"errors"
	"fmt"
)

// Service of billing in another API
type Service struct{}

// Billing method send the sale ID to billing in AFIP API
func (s Service) Billing(saleID int) error {
	if saleID < 1 {
		return errors.New("The sale ID is invalid")
	}
	fmt.Println("Send sale", saleID, "to AFIP")
	return nil
}
