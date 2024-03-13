package test

import (
	"fmt"
	"testing"

	"github.com/Gromitmugs/temporal-playground/job/broadcast"
)

func TestBroadcast(t *testing.T) {
	err := broadcast.Starter("TestMessage")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
