package main

import (
	"fmt"
	"os"

	"github.com/Gromitmugs/temporal-playground/job/billing"
	"github.com/Gromitmugs/temporal-playground/job/broadcast"
	"github.com/Gromitmugs/temporal-playground/job/builder"
	"github.com/Gromitmugs/temporal-playground/job/scheduler"
	"github.com/Gromitmugs/temporal-playground/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
		please specify your operation through input arguments
		go run main.go <operation_name>
		`)
		return
	}
	client, err := service.GetTemporalClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	operation := os.Args[1]
	switch operation {
	case "broadcast":
		service.InitWorker(&client, broadcast.TaskQueueName, broadcast.Workflow.Definition, broadcast.Workflow.Activities...)
	case "builder":
		service.InitWorker(&client, builder.TaskQueueName, builder.Workflow.Definition, builder.Workflow.Activities...)
	case "scheduler":
		service.InitWorker(&client, scheduler.TaskQueueName, scheduler.Workflow.Definition, scheduler.Workflow.Activities...)
	case "billing":
		service.InitWorker(&client, billing.TaskQueueName, billing.Workflow.Definition, billing.Workflow.Activities...)
	default:
		fmt.Println("no operation found")
		return
	}

}
