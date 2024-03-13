package test

import (
	"fmt"
	"testing"

	"github.com/Gromitmugs/temporal-playground/job/builder"
)

func TestBuilder(t *testing.T) {
	const testRepoUrl = "https://github.com/Gromitmugs/hello-world-docker"
	err := builder.Starter(testRepoUrl)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
