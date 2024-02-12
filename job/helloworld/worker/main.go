package main

import (
	"github.com/Gromitmugs/temporal-playground/job/helloworld"
	"github.com/Gromitmugs/temporal-playground/service"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := service.GetTemporalClient()
	service.PanicIfErr(err)
	defer c.Close()

	service.RegisterWorkflowAndActivity(&c, helloworld.TaskName, helloworld.Workflow, helloworld.Activity)
}
