package main_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/Gromitmugs/temporal-playground/job/broadcast"
	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/google/uuid"
	temporalclient "go.temporal.io/sdk/client"
)

func TestBroadCastStarter(t *testing.T) {
	client, err := service.GetTemporalClient()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer client.Close()

	startOpt := temporalclient.StartWorkflowOptions{
		ID:        uuid.New().String(),
		TaskQueue: "TaskQueueBroadcast",
	}

	workflowRun, err := client.ExecuteWorkflow(context.Background(), startOpt, broadcast.Workflow, "TestMsg2")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
		return
	}

	// Synchronously wait for the workflow completion.
	var result string
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
		return
	}
	log.Println("Workflow result:", result)
}
