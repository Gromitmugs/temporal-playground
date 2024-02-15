package billing

import (
	"context"
	"log"
	"time"

	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/google/uuid"
	"go.temporal.io/api/enums/v1"
	temporalclient "go.temporal.io/sdk/client"
)

func Starter() error {
	client, err := service.GetTemporalClient()
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := context.Background()

	scheduleHandle, err := client.ScheduleClient().Create(ctx, temporalclient.ScheduleOptions{
		ID: uuid.NewString(),
		Spec: temporalclient.ScheduleSpec{
			Calendars: []temporalclient.ScheduleCalendarSpec{},
			Intervals: []temporalclient.ScheduleIntervalSpec{
				{
					Every: 2 * time.Second, // every seconds that ends with the input number
				},
			},
		},
		Action: &temporalclient.ScheduleWorkflowAction{
			ID:        uuid.NewString(),
			Workflow:  BillingWorkflow,
			TaskQueue: TaskQueueName,
		},
		Overlap:            enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
		RemainingActions:   5,
		TriggerImmediately: true,
		Note:               "Billing Service For Person ID: 1",
		Paused:             false,
	})
	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}
	defer func() {
		err = scheduleHandle.Delete(ctx)
		if err != nil {
			log.Fatalln("Unable to delete schedule", err)
		}
	}()

	if err := handlerWaitUntilFinish(ctx, scheduleHandle); err != nil {
		return err
	}

	return nil
}

func triggerScheduleOnce(ctx context.Context, handler temporalclient.ScheduleHandle) error {
	return handler.Trigger(ctx, temporalclient.ScheduleTriggerOptions{
		Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
	})
}

func handlerWaitUntilFinish(ctx context.Context, handler temporalclient.ScheduleHandle) error {
	// for {
	// 	description, err := handler.Describe(ctx)
	// 	if err != nil {
	// 		log.Fatalln("Unable to describe schedule", err)
	// 	}
	// 	if description.Schedule.State.RemainingActions != 0 {
	// 		log.Println("Schedule has remaining actions", "ScheduleID", handler.GetID(), "RemainingActions", description.Schedule.State.RemainingActions)
	// 		time.Sleep(5 * time.Second)
	// 	} else {
	// 		break
	// 	}
	// }
	time.Sleep(time.Second * 60)
	return nil
}
