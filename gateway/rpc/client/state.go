package client

import (
	"context"
	"time"

	"github.com/haozhe-qi/ChatIm/common/config"
	"github.com/haozhe-qi/ChatIm/common/prpc"
	"github.com/haozhe-qi/ChatIm/state/rpc/service"
)

// gateway把state.go作为client端，调用state的delconn和sendmsg服务
var stateClient service.StateClient

func initStateClient() {
	pCli, err := prpc.NewPClient(config.GetStateServiceName())
	if err != nil {
		panic(err)
	}
	stateClient = service.NewStateClient(pCli.Conn())

}

func CancleConn(ctx *context.Context, endpoint string, fd int32, payLoad []byte) error {
	rpcCtx, _ := context.WithTimeout(*ctx, 100*time.Millisecond)
	stateClient.CancelConn(rpcCtx, &service.StateRequest{
		Endpoint: endpoint,
		Fd:       fd,
		Data:     payLoad,
	})
	return nil
}

// gateway读取固定长度去除len，按len读出payload，直接发给state server
func SendMsg(ctx *context.Context, endpoint string, fd int32, payLoad []byte) error {
	rpcCtx, _ := context.WithTimeout(*ctx, 100*time.Millisecond)
	//request放进去直接调用 state server收到后会把这个放channel里，开个协程去channel里处理
	_, err := stateClient.SendMsg(rpcCtx, &service.StateRequest{
		Endpoint: endpoint,
		Fd:       fd,
		Data:     payLoad,
	})
	if err != nil {
		panic(err)
	}
	return nil
}
