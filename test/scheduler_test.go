package test

import (
	"fmt"
	"testing"

	"github.com/Gromitmugs/temporal-playground/job/scheduler"
)

func TestScheduler(t *testing.T) {
	if err := scheduler.Starter(); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
}
