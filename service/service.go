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

func StartWorkflow(c *client.Client, workflow interface{}, activities ...interface{}) {
	w := worker.New(*c, "hello-world", worker.Options{})
	w.RegisterWorkflow(workflow)
	for _, activity := range activities {
		w.RegisterActivity(activity)
	}
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
