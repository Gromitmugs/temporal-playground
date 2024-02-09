package broadcast

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"
)

func BroadcastActivity(ctx context.Context, broadcastMsg string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "Broadcast Activity Started")
	return fmt.Sprint("Broadcast Message: ", broadcastMsg), nil
}
