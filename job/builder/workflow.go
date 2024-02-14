package builder

import (
	"time"

	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/Gromitmugs/temporal-playground/thirdparty/client"
	"go.temporal.io/sdk/workflow"
)

var Workflow *service.Workflow = &service.Workflow{
	Definition: BuilderWorkflow,
	Activities: []interface{}{
		CloneRepo,
		BuildImage,
		RemoveClonedRepo,
	},
}

func BuilderWorkflow(ctx workflow.Context, broadcastMsg string) (string, error) {
	opt := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, opt)
	logger := workflow.GetLogger(ctx)

	var clonePath string
	err := workflow.ExecuteActivity(ctx, CloneRepo, broadcastMsg).Get(ctx, &clonePath)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	defer workflow.ExecuteActivity(ctx, RemoveClonedRepo, clonePath).Get(ctx, nil)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	var buildImageLog client.MessageCreateResult
	err = workflow.ExecuteActivity(ctx, BuildImage, broadcastMsg).Get(ctx, &buildImageLog)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	return "Successfully build and clone the repository", nil
}
