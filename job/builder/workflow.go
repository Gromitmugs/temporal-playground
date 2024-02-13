package builder

import (
	"time"

	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/Gromitmugs/temporal-playground/thirdparty/client"
	"go.temporal.io/sdk/workflow"
)

var Workflow *service.Workflow = &service.Workflow{
	Definition: BuilderWorkflow,
	Activities: []interface{}{},
}

func BuilderWorkflow(ctx workflow.Context, broadcastMsg string) (string, error) {
	opt := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, opt)
	logger := workflow.GetLogger(ctx)

	var result string
	err := workflow.ExecuteActivity(ctx, CloneRepo, broadcastMsg).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	var recordMessageResult client.MessageCreateResult
	err = workflow.ExecuteActivity(ctx, BuildImage, broadcastMsg).Get(ctx, &recordMessageResult)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	return result, nil
}
