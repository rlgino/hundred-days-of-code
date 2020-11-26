package observer

import (
	"github.com/rlgino/hundred-days-of-code/patterns/observer/model"
)

// SaleObserver observes the sale
type SaleObserver interface {
	ExecutePreSale(s *model.Sale) error
	ExecutePostSale(s *model.Sale) error
}

// ObserversContainer it's the container of the sale observers
type ObserversContainer struct {
	observers []SaleObserver
}

// Subscribe new observer into a subject
func (container *ObserversContainer) Subscribe(observer SaleObserver) {
	container.observers = append(container.observers, observer)
}

// Remove an observer of the list
func (container *ObserversContainer) Remove(observer SaleObserver) {
	res := -1
	for i, item := range container.observers {
		if item == observer {
			res = i
		}
	}
	observersLength := len(container.observers)
	container.observers[observersLength] = container.observers[res]
	container.observers[res] = nil
	container.observers = container.observers[:observersLength-1]
}

// ExecutePreSaleRules executes the pre sale rules
func (container ObserversContainer) ExecutePreSaleRules(sale *model.Sale) error {
	for _, obs := range container.observers {
		err := obs.ExecutePreSale(sale)
		if err != nil {
			return err
		}
	}
	return nil
}

// ExecutePostSaleRules executes the post sale rules
func (container ObserversContainer) ExecutePostSaleRules(sale *model.Sale) error {
	for _, obs := range container.observers {
		err := obs.ExecutePostSale(sale)
		if err != nil {
			return err
		}
	}
	return nil
}
