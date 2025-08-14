package state

import (
	"context"
	"fmt"

	"github.com/haozhe-qi/ChatIm/common/config"
	"github.com/haozhe-qi/ChatIm/common/prpc"
	"github.com/haozhe-qi/ChatIm/state/rpc/client"
	"github.com/haozhe-qi/ChatIm/state/rpc/service"
	"google.golang.org/grpc"
)

var cmdChannel chan *service.CmdContext

// RunMain 启动网关服务
func RunMain(path string) {
	config.Init(path)
	cmdChannel = make(chan *service.CmdContext, config.GetStateCmdChannelNum())

	s := prpc.NewPServer(
		prpc.WithServiceName(config.GetStateServiceName()),
		prpc.WithIP(config.GetStateServiceAddr()),
		prpc.WithPort(config.GetSateServerPort()), prpc.WithWeight(config.GetStateRPCWeight()))

	s.RegisterService(func(server *grpc.Server) {
		service.RegisterStateServer(server, &service.Service{CmdChannel: cmdChannel})
	})

	//初始化rpc客户端
	client.Init()
	//处理写协程
	go cmdHandler()
	//启动rpc server
	s.Start(context.TODO())
}

func cmdHandler() {
	for cmd := range cmdChannel {
		switch cmd.Cmd {
		case service.CancelConnCmd:
			fmt.Printf("cancelconn endpoint:%s, fd:%d, data:%+v", cmd.Endpoint, cmd.FD, cmd.Payload)
		case service.SendMsgCmd:
			fmt.Println("cmdHandler", string(cmd.Payload))
			//收到gateway发来的payload由于没实现业务层，这里直接再发回去
			client.Push(cmd.Ctx, int32(cmd.FD), cmd.Payload)
		}
	}
}
