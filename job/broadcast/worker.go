package broadcast

import (
	"github.com/Gromitmugs/temporal-playground/job/broadcast/constant"
	"github.com/Gromitmugs/temporal-playground/service"
)

func Worker() {
	c, err := service.GetTemporalClient()
	service.PanicIfErr(err)
	defer c.Close()

	service.RegisterWorkflowAndActivity(&c, constant.TaskName, Workflow, BroadcastActivity)
}
