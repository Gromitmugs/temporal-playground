package service

import (
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type Workflow struct {
	Definition interface{}
	Activities []interface{}
}

func GetTemporalClient() (client.Client, error) {
	c, err := client.Dial(client.Options{
		HostPort: os.Getenv("TEMPORAL_ADDRESS"),
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
		return nil, err
	}
	return c, nil
}

func InitWorker(c *client.Client, name string, workflow interface{}, activities ...interface{}) {
	w := worker.New(*c, name, worker.Options{})
	w.RegisterWorkflow(workflow)
	for _, activity := range activities {
		w.RegisterActivity(activity)
	}
	err := w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
