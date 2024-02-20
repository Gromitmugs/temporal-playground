package builder

import (
	"time"

	"github.com/Gromitmugs/temporal-playground/service"
	"go.temporal.io/sdk/workflow"
)

var Workflow *service.Workflow = &service.Workflow{
	Definition: BuilderWorkflow,
	Activities: []interface{}{
		KanikoCloneAndBuildImage,
	},
}

func BuilderWorkflow(ctx workflow.Context, broadcastMsg string) (string, error) {
	opt := workflow.ActivityOptions{
		ScheduleToCloseTimeout: time.Minute * 10,
		ScheduleToStartTimeout: time.Minute * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, opt)
	logger := workflow.GetLogger(ctx)

	var result string
	err := workflow.ExecuteActivity(ctx, KanikoCloneAndBuildImage, broadcastMsg).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	return "Successfully build the repository", nil
}
