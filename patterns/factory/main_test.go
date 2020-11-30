package main_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/factory/mobileinterface"
	"github.com/rlgino/hundred-days-of-code/patterns/factory/webinterface"
)

func Test_success_mobileFactory(t *testing.T) {
	factory := mobileinterface.New()
	if factory == nil {
		t.Error("Factory shouldn't be null")
	}

	mobileBtn := factory.CreateButton()
	if mobileBtn == nil {
		t.Error("mobile button shouldn't be null")
	}

	mobileBtn.Draw()
	mobileBtn.Click()
}

func Test_fail_mobileFactory(t *testing.T) {
	factory := mobileinterface.New()
	if factory == nil {
		t.Error("Factory shouldn't be null")
	}

	mobileBtn := factory.CreateButton()
	if mobileBtn == nil {
		t.Error("mobile button shouldn't be null")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("It should fail")
		}
	}()

	mobileBtn.Click()

}

func Test_success_webFactory(t *testing.T) {
	factory := webinterface.New()
	if factory == nil {
		t.Error("Factory shouldn't be null")
	}

	webBtn := factory.CreateButton()
	if webBtn == nil {
		t.Error("web button shouldn't be null")
	}

	webBtn.Draw()
	webBtn.Click()
}
