package broadcast

import (
	"context"
	"fmt"

	"github.com/Gromitmugs/temporal-playground/thirdparty/client"
	"go.temporal.io/sdk/activity"
)

func BroadcastMessage(ctx context.Context, broadcastMsg string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "Broadcast Activity Started")
	return fmt.Sprint("Broadcast Message: ", broadcastMsg), nil
}

// The following functions sends a request to 3rd party services
// In this case, we use a GQL client as an interface.
func RecordMessage(ctx context.Context, message string) (*client.MessageCreateResult, error) {
	c := client.New(client.EndpointUrl, map[string]string{})
	result, err := c.MessageCreate(ctx, message)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ErrorProduce(ctx context.Context, errMsg string) error {
	c := client.New(client.EndpointUrl, map[string]string{})
	if err := c.ErrorCreate(ctx, errMsg); err != nil {
		return err
	}
	return nil
}
