package main_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/observer/model"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/observer"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/rule1"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/rule2"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/rule3"
)

func Test_mainTest(t *testing.T) {
	// Setup
	sale := model.Sale{
		ID:      1,
		Amounts: []int{1, 2, 3},
		Status:  model.PENDING,
	}

	container := observer.ObserversContainer{}
	container.Subscribe(rule1.New())
	container.Subscribe(rule2.New())
	container.Subscribe(rule3.New())

	// Execute
	err1 := container.ExecutePreSaleRules(&sale)
	err2 := container.ExecutePostSaleRules(&sale)

	// Verify
	if err1 != nil {
		t.Error(err1.Error())
	}
	if err2 != nil {
		t.Error(err2.Error())
	}
	if sale.ID != 123 {
		t.Error("Invalid ID")
	}
	if sale.TotalAmount != 6 {
		t.Error("Invalid Amount")
	}
}
