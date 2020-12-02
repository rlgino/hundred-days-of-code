package main_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/adapter/adapter"
)

func Test_Success(t *testing.T) {
	savingBox := adapter.New(0)
	savingBox.AddARS(89)
	savingBox.AddARS(89)

	result := savingBox.GetAmount()
	if result != 178 {
		t.Error("The result should be 178")
	}
}
