package scheduler

import (
	"github.com/Gromitmugs/temporal-playground/service"
)

func Starter() error {
	client, err := service.GetTemporalClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// TODO JEDI: https://github.com/temporalio/samples-go/blob/main/schedule/starter/main.go
	return nil
}
