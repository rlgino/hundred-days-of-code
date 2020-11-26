package main_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/bridge/agregation"
	"github.com/rlgino/hundred-days-of-code/patterns/bridge/bridge"
	"github.com/rlgino/hundred-days-of-code/patterns/bridge/substraction"
)

func Test_mainAgregationTest(t *testing.T) {
	agregationImpl := bridge.New(agregation.New())
	result := agregationImpl.Calculate(1, 1)

	if result != 2 {
		t.Error("The result should be 2")
	}
}

func Test_mainSubstractionTest(t *testing.T) {
	substractionImpl := bridge.New(substraction.New())
	result := substractionImpl.Calculate(1, 1)

	if result != 0 {
		t.Error("The result should be 0")
	}
}
