package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Gromitmugs/temporal-playground/job/builder"
)

func TestBuilder(t *testing.T) {
	ctx := context.Background()
	clonePath, err := builder.CloneRepo(ctx)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()

	}
	if err := builder.BuildImage(ctx, clonePath); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
}
