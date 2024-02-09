package broadcast

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context, broadcastMsg string) (string, error) {
	opt := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, opt)

	logger := workflow.GetLogger(ctx)
	logger.Info("Broadcast workflow started")

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, broadcastMsg).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("Broadcast workflow completed.", "result", result)
	return result, nil
}

func Activity(ctx context.Context, broadcastMsg string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "Broadcast Activity Started")
	return fmt.Sprint("Broadcast Message:", broadcastMsg), nil
}
