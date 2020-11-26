package main

import (
	"encoding/json"
	"fmt"

	"github.com/rlgino/hundred-days-of-code/patterns/observer/model"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/observer"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/rule1"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/rule2"
	"github.com/rlgino/hundred-days-of-code/patterns/observer/rule3"
)

func main() {
	sale := model.Sale{
		ID:      1,
		Amounts: []int{1, 2, 3},
		Status:  model.PENDING,
	}

	container := subscribeObservers()
	errPre := container.ExecutePreSaleRules(&sale)
	if errPre != nil {
		fmt.Println(errPre.Error())
	}
	errPost := container.ExecutePostSaleRules(&sale)
	if errPost != nil {
		fmt.Println(errPost.Error())
	}

	saleData, _ := json.Marshal(sale)
	fmt.Println("Sale result:", string(saleData))
}

func subscribeObservers() (container observer.ObserversContainer) {
	container = observer.ObserversContainer{}
	container.Subscribe(rule1.New())
	container.Subscribe(rule2.New())
	container.Subscribe(rule3.New())
	return
}
