package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Gromitmugs/temporal-playground/job/broadcast"
	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/google/uuid"
	temporalclient "go.temporal.io/sdk/client"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`
		please specify your message through a first argument
		go run main.go <your_message>
		`)
		return
	}
	message := os.Args[1]
	fmt.Println("your broadcast message is: ", message)

	client, err := service.GetTemporalClient()
	service.HandleErr(err)
	defer client.Close()

	startOpt := temporalclient.StartWorkflowOptions{
		ID:        uuid.New().String(),
		TaskQueue: broadcast.TaskName,
	}

	workflowRun, err := client.ExecuteWorkflow(context.Background(), startOpt, broadcast.Workflow, message)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
		return
	}

	var resp string
	err = workflowRun.Get(context.Background(), &resp)
	service.HandleErr(err)
	log.Println("Workflow result:", resp)
}
