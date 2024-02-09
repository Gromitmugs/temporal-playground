package main

import (
	"fmt"

	"github.com/Gromitmugs/temporal-playground/job/broadcast"
	"github.com/Gromitmugs/temporal-playground/service"
)

func main() {
	client, err := service.GetTemporalClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	service.RegisterWorkflowAndActivity(&client, broadcast.Workflow, broadcast.Activity)
}
