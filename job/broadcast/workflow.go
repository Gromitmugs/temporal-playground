package broadcast

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

const TaskName string = "Broadcast"

func Workflow(ctx workflow.Context, broadcastMsg string) (string, error) {
	opt := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, opt)

	logger := workflow.GetLogger(ctx)
	logger.Info("Broadcast workflow started")

	var result string
	err := workflow.ExecuteActivity(ctx, BroadcastActivity, broadcastMsg).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("Broadcast workflow completed.", "result", result)
	return result, nil
}
