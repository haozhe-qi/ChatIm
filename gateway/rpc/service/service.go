package service

import "context"

const (
	DelConnCmd = 1
	PushCmd    = 2
)

type CmdContext struct {
	Ctx     *context.Context
	Cmd     int32
	FD      int
	Payload []byte
}

//	type Service struct {
//		CmdChannel chan *CmdContext
//	}
type Service struct {
	UnimplementedGatewayServer // 必须嵌入，以满足 gRPC 接口要求
	CmdChannel                 chan *CmdContext
}

func (s *Service) DelConn(ctx context.Context, gr *GatewayRequest) (*GatewayResponse, error) {
	c := context.TODO()
	s.CmdChannel <- &CmdContext{
		Ctx: &c,
		Cmd: DelConnCmd,
		FD:  int(gr.GetFd()),
	}
	return &GatewayResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}

// 收到state server发来的消息也是根据request打包一个CmdContext放channel里等着被处理
func (s *Service) Push(ctx context.Context, gr *GatewayRequest) (*GatewayResponse, error) {
	c := context.TODO()
	s.CmdChannel <- &CmdContext{
		Ctx:     &c,
		Cmd:     PushCmd,
		FD:      int(gr.GetFd()),
		Payload: gr.GetData(),
	}
	return &GatewayResponse{
		Code: 0,
		Msg:  "success",
	}, nil
}
