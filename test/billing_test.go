package test

import (
	"fmt"
	"testing"

	"github.com/Gromitmugs/temporal-playground/job/billing"
)

func TestBilling(t *testing.T) {
	if err := billing.Starter(); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
}
