package billing

import (
	"time"

	"github.com/Gromitmugs/temporal-playground/service"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/workflow"
)

var Workflow *service.Workflow = &service.Workflow{
	Definition: BillingWorkflow,
	Activities: []interface{}{
		Payment,
	},
}

func BillingWorkflow(ctx workflow.Context) error {
	opt := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, opt)
	info := workflow.GetInfo(ctx)

	// Workflow Executions started by a Schedule have the following additional properties appended to their search attributes
	scheduledByIDPayload := info.SearchAttributes.IndexedFields["TemporalScheduledById"]
	var scheduledByID string
	err := converter.GetDefaultDataConverter().FromPayload(scheduledByIDPayload, &scheduledByID)
	if err != nil {
		return err
	}

	startTimePayload := info.SearchAttributes.IndexedFields["TemporalScheduledStartTime"]
	var startTime time.Time
	err = converter.GetDefaultDataConverter().FromPayload(startTimePayload, &startTime)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, Payment, scheduledByID, startTime).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("schedule workflow failed.", "Error", err)
		return err
	}
	return nil
}
