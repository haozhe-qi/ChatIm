package ipconf

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/haozhe-qi/ChatIm/common/config"
	"github.com/haozhe-qi/ChatIm/ipconf/domain"
	"github.com/haozhe-qi/ChatIm/ipconf/source"
)

// RunMain启动web容器
func RunMain(path string) {
	config.Init(path)
	source.Init() //数据源要先启动即etcd那套
	domain.Init() //初始化调度层
	s := server.Default(server.WithHostPorts(":6789"))
	s.GET("/ip/list", GetIpInfoList)
	s.Spin()

}
