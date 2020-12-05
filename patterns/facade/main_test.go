package main_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/facade/entities"
	"github.com/rlgino/hundred-days-of-code/patterns/facade/facade"
)

func Test_SaleHander(t *testing.T) {
	test := []struct {
		name     string
		sale     entities.Sale
		errorMsg string
	}{
		{
			name: "Sale finishing should be OK",
			sale: entities.Sale{
				ID:           12,
				ProductsList: []string{"TV", "Computer", "Table"},
			},
		},
		{
			name: "Sale should be fail because it hasn't products",
			sale: entities.Sale{
				ID:           12,
				ProductsList: []string{},
			},
			errorMsg: "Product list is empty",
		},
		{
			name: "Sale finishing should be fail because the product is invalid",
			sale: entities.Sale{
				ID:           12,
				ProductsList: []string{""},
			},
			errorMsg: "The product is invalid",
		},
		{
			name: "Sale finishing should be fail because the ID is invalid",
			sale: entities.Sale{
				ID:           -12,
				ProductsList: []string{"TV"},
			},
			errorMsg: "The sale ID is invalid",
		},
	}

	handler := facade.New()
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.Finish(tt.sale)
			if err != nil {
				if err.Error() != tt.errorMsg {
					t.Error("The expected error is", err.Error())
				}
			} else {
				if tt.errorMsg != "" {
					t.Error("The expected error is", err.Error())
				}
			}
		})
	}
}
