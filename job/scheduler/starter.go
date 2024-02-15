package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/Gromitmugs/temporal-playground/service"
	"github.com/google/uuid"
	temporalclient "go.temporal.io/sdk/client"
)

func Starter() error {
	client, err := service.GetTemporalClient()
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := context.Background()

	// This schedule ID can be user business logic identifier as well.
	// Create the schedule, start with no spec so the schedule will not run.
	scheduleHandle, err := client.ScheduleClient().Create(ctx, temporalclient.ScheduleOptions{
		ID:   uuid.NewString(),
		Spec: temporalclient.ScheduleSpec{},
		Action: &temporalclient.ScheduleWorkflowAction{
			ID:        uuid.NewString(),
			Workflow:  SampleScheduleWorkflow,
			TaskQueue: TaskQueueName,
		},
	})
	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}
	// Delete the schedule once the sample is done
	defer func() {
		log.Println("Deleting schedule", "ScheduleID", scheduleHandle.GetID())
		err = scheduleHandle.Delete(ctx)
		if err != nil {
			log.Fatalln("Unable to delete schedule", err)
		}
	}()

	// Manually trigger the schedule once
	// err = scheduleHandle.Trigger(ctx, temporalclient.ScheduleTriggerOptions{
	// 	Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
	// })
	// if err != nil {
	// 	log.Fatalln("Unable to trigger schedule", err)
	// }

	// Update the schedule with a spec so it will run periodically,
	log.Println("Updating schedule", "ScheduleID", scheduleHandle.GetID())
	err = scheduleHandle.Update(ctx, temporalclient.ScheduleUpdateOptions{
		DoUpdate: func(schedule temporalclient.ScheduleUpdateInput) (*temporalclient.ScheduleUpdate, error) {
			schedule.Description.Schedule.Spec = &temporalclient.ScheduleSpec{
				// Run the schedule at 5pm on Friday
				Calendars: []temporalclient.ScheduleCalendarSpec{
					{
						Hour: []temporalclient.ScheduleRange{
							{
								Start: 17,
							},
						},
						DayOfWeek: []temporalclient.ScheduleRange{
							{
								Start: 5,
							},
						},
					},
				},
				// Run the schedule every 5s
				Intervals: []temporalclient.ScheduleIntervalSpec{
					{
						Every: 5 * time.Second,
					},
				},
			}
			// Start the schedule paused to demonstrate how to unpause a schedule
			schedule.Description.Schedule.State.Paused = true
			schedule.Description.Schedule.State.LimitedActions = true
			schedule.Description.Schedule.State.RemainingActions = 10
			return &temporalclient.ScheduleUpdate{
				Schedule: &schedule.Description.Schedule,
			}, nil
		},
	})
	if err != nil {
		log.Fatalln("Unable to update schedule", err)
	}

	// Unpause schedule
	log.Println("Unpausing schedule", "ScheduleID", scheduleHandle.GetID())
	err = scheduleHandle.Unpause(ctx, temporalclient.ScheduleUnpauseOptions{})
	if err != nil {
		log.Fatalln("Unable to unpause schedule", err)
	}
	// Wait for the schedule to run 10 actions
	log.Println("Waiting for schedule to complete 10 actions", "ScheduleID", scheduleHandle.GetID())

	for {
		description, err := scheduleHandle.Describe(ctx)
		if err != nil {
			log.Fatalln("Unable to describe schedule", err)
		}
		if description.Schedule.State.RemainingActions != 0 {
			log.Println("Schedule has remaining actions", "ScheduleID", scheduleHandle.GetID(), "RemainingActions", description.Schedule.State.RemainingActions)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	return nil
}
