package database_test

import (
	"testing"

	"github.com/rlgino/hundred-days-of-code/patterns/singleton/database"
)

func Test_GetInstanceTest(t *testing.T) {
	instanceA := database.GetInstance()
	instanceB := database.GetInstance()
	if instanceA.GetID() != instanceB.GetID() {
		t.Errorf("The IDs must be equals")
	}
}
