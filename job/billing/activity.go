package billing

import (
	"context"
	"fmt"
	"time"

	"github.com/Gromitmugs/temporal-playground/job/broadcast"
)

func Payment(ctx context.Context, scheduleByID string, startTime time.Time) error {
	message := fmt.Sprintf(`
		Payment has been processed.
		ScheduleById %s.
		StartTime %s.
	`, scheduleByID, startTime.String())
	_, err := broadcast.RecordMessage(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
