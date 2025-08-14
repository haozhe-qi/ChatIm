package client

import (
	"context"
	"fmt"
	"time"

	"github.com/haozhe-qi/ChatIm/common/config"
	"github.com/haozhe-qi/ChatIm/common/prpc"
	"github.com/haozhe-qi/ChatIm/gateway/rpc/service"
)

// state把该gateway.go作为客户端去请求gateway的delconn和push服务
var gatewayClient service.GatewayClient

func initGatewayClient() {
	pCli, err := prpc.NewPClient(config.GetGatewayServiceName())
	if err != nil {
		panic(err)
	}
	gatewayClient = service.NewGatewayClient(pCli.Conn())
}

func DelConn(ctx *context.Context, fd int32, playLoad []byte) error {
	rpcCtx, _ := context.WithTimeout(*ctx, 100*time.Millisecond)
	gatewayClient.DelConn(rpcCtx, &service.GatewayRequest{Fd: fd, Data: playLoad})
	return nil
}

func Push(ctx *context.Context, fd int32, playLoad []byte) error {
	rpcCtx, _ := context.WithTimeout(*ctx, 100*time.Second)
	//打包个request发回去
	resp, err := gatewayClient.Push(rpcCtx, &service.GatewayRequest{Fd: fd, Data: playLoad})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	return nil
}
