package builder

import (
	"context"

	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/google/uuid"
	temporalclient "go.temporal.io/sdk/client"
)

func Starter(repoUrl string) error {
	client, err := service.GetTemporalClient()
	if err != nil {
		return err
	}
	defer client.Close()

	startOpt := temporalclient.StartWorkflowOptions{
		ID:        uuid.New().String(),
		TaskQueue: TaskQueueName,
	}

	workflowRun, err := client.ExecuteWorkflow(context.Background(), startOpt, Workflow.Definition, repoUrl)
	if err != nil {
		return err
	}

	var resp string
	if err = workflowRun.Get(context.Background(), &resp); err != nil {
		return err
	}
	return nil
}
