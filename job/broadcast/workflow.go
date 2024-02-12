package broadcast

import (
	"time"

	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/Gromitmugs/temporal-playground/thirdparty/client"
	"go.temporal.io/sdk/workflow"
)

var Workflow *service.Workflow = &service.Workflow{
	Definition: BroadcastWorkflow,
	Activities: []interface{}{
		BroadcastMessage,
	},
}

func BroadcastWorkflow(ctx workflow.Context, broadcastMsg string) (string, error) {
	opt := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, opt)

	logger := workflow.GetLogger(ctx)
	logger.Info("Broadcast workflow started")

	var result string
	err := workflow.ExecuteActivity(ctx, BroadcastMessage, broadcastMsg).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	var recordMessageResult client.MessageCreateResult
	err = workflow.ExecuteActivity(ctx, RecordMessage, broadcastMsg).Get(ctx, &recordMessageResult)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("Broadcast workflow completed.", "result", result)
	return result, nil
}
