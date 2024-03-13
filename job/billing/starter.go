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

func Starter(scheduleId string) error {
	client, err := service.GetTemporalClient()
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := context.Background()
	_, err = client.ScheduleClient().Create(ctx, temporalclient.ScheduleOptions{
		ID: scheduleId,
		Spec: temporalclient.ScheduleSpec{
			StartAt: time.Now(),
			Intervals: []temporalclient.ScheduleIntervalSpec{
				{
					Every: 60 * time.Second,
				},
			},
		},
		Action: &temporalclient.ScheduleWorkflowAction{
			ID:        uuid.NewString(),
			Workflow:  BillingWorkflow,
			TaskQueue: TaskQueueName,
		},
		Overlap:            enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
		RemainingActions:   3,
		TriggerImmediately: false,
		Note:               "Billing Service For Person ID: 1",
		Paused:             false,
	})
	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}

	return nil
}

func triggerScheduleOnce(ctx context.Context, handler temporalclient.ScheduleHandle) error {
	return handler.Trigger(ctx, temporalclient.ScheduleTriggerOptions{
		Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
	})
}

func handlerWaitUntilFinish(ctx context.Context, handler temporalclient.ScheduleHandle) error {
	for {
		description, err := handler.Describe(ctx)
		if err != nil {
			log.Fatalln("Unable to describe schedule", err)
		}
		if description.Schedule.State.RemainingActions != 0 {
			log.Println("Schedule has remaining actions", "ScheduleID", handler.GetID(), "RemainingActions", description.Schedule.State.RemainingActions)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	return nil
}

func deleteScheduleById(ctx context.Context, id string) error {
	client, err := service.GetTemporalClient()
	if err != nil {
		return err
	}
	defer client.Close()

	return client.ScheduleClient().GetHandle(ctx, id).Delete(ctx)
}

func CancelScheduleById(id string) error {
	return deleteScheduleById(context.TODO(), id)
}
