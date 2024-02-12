package main

import (
	"fmt"
	"os"

	"github.com/Gromitmugs/temporal-playground/job/broadcast"
	"github.com/Gromitmugs/temporal-playground/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
		please specify your operation through input arguments
		go run main.go <operation_name>
		`)
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
		service.RegisterWorkflowAndActivity(&client, "Broadcast", broadcast.Workflow.Definition, broadcast.Workflow.Activities...)
	default:
		fmt.Println("no operation found")
		return
	}

}
