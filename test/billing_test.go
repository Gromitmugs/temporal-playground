package test

import (
	"fmt"
	"testing"

	"github.com/Gromitmugs/temporal-playground/job/billing"
	"github.com/google/uuid"
)

const id = "6acdcabe-7323-41c2-876e-529ad3b1aa71"

func init() {
	uuid.EnableRandPool()
}

func TestBilling(t *testing.T) {
	if err := billing.Starter(uuid.NewString()); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
}

func TestCancelBilling(t *testing.T) {
	if err := billing.CancelScheduleById(id); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

}
