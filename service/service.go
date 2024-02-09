package service

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func GetTemporalClient() (client.Client, error) {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
		return nil, err
	}
	return c, nil
}

func RegisterWorkflowAndActivity(c *client.Client, workflow interface{}, activities ...interface{}) {
	w := worker.New(*c, "hello-world", worker.Options{})
	w.RegisterWorkflow(workflow)
	for _, activity := range activities {
		w.RegisterActivity(activity)
	}
	err := w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
