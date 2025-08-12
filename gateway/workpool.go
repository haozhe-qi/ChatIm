package gateway

import (
	"fmt"

	"github.com/haozhe-qi/ChatIm/common/config"
	"github.com/panjf2000/ants/v2"
)

var wPool *ants.Pool

func initWorkPoll() {
	var err error
	if wPool, err = ants.NewPool(config.GetGatewayWorkerPoolNum()); err != nil {
		fmt.Printf("InitWorkPoll.err :%s num:%d\n", err.Error(), config.GetGatewayWorkerPoolNum())
	}
}
