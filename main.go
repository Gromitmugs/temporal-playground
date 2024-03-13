package main

import (
	"context"
	"fmt"
	"io/ioutil"
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

	operation := os.Args[1]
	debugTools(operation)

	client, err := service.GetTemporalClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

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

func debugTools(operation string) {
	switch operation {
	case "kaniko":
		builder.TestCloneAndBuild()
		os.Exit(0)
	case "clone":
		clonedPath, err := builder.CloneRepo(context.TODO(), "https://github.com/Gromitmugs/hello-world-docker")
		fmt.Println(err.Error())
		fmt.Println(clonedPath)
		os.Exit(0)
	case "ls":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a path for ls command")
			os.Exit(0)
		}
		fileInfos, err := ioutil.ReadDir(os.Args[2])
		if err != nil {
			fmt.Println("Error in accessing directory:", err)
		}
		for _, file := range fileInfos {
			fmt.Println(file.Name())
		}
		os.Exit(0)
	}
}
