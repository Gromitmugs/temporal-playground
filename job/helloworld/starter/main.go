package main

import (
	"context"
	"log"

	"github.com/Gromitmugs/temporal-playground/job/helloworld"
	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := service.GetTemporalClient()
	service.PanicIfErr(err)
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        uuid.New().String(),
		TaskQueue: helloworld.TaskName, // needs to be the same for both worker and starter
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, helloworld.Workflow, "Temporal")
	service.PanicIfErr(err)

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	service.PanicIfErr(err)
	log.Println("Workflow result:", result)
}
