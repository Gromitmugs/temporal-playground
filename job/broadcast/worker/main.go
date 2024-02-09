package main

import (
	"github.com/Gromitmugs/temporal-playground/job/broadcast"
	"github.com/Gromitmugs/temporal-playground/service"
)

func main() {
	c, err := service.GetTemporalClient()
	service.HandleErr(err)
	defer c.Close()

	service.RegisterWorkflowAndActivity(&c, broadcast.TaskName, broadcast.Workflow, broadcast.BroadcastActivity)
}
