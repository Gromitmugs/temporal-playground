package scheduler

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

// DoSomething is an Activity
func DoSomething(ctx context.Context, scheduleByID string, startTime time.Time) error {
	activity.GetLogger(ctx).Info("Schedulde job running.", "scheduleByID", scheduleByID, "startTime", startTime)
	// Query database, call external API, or do any other non-deterministic action.
	return nil
}
