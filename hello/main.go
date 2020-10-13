package main

import (
	"context"
	"fmt"

	proto "github.com/love666666shen/go-microserv/hello/proto"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/sirupsen/logrus"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, resp *proto.HelloResponse) (err error) {
	resp.Greeting = "你好，" + req.Name

	return
}

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2379"}
	})

	// 创建一个新的greeterService服务
	greeterService := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Registry(reg),
	)

	// 服务初始化，会解析命令行参数
	greeterService.Init()

	// 注册处理器，调用Greeter服务接口处理请求
	if err := proto.RegisterGreeterHandler(greeterService.Server(), new(Greeter)); err != nil {
		logrus.Errorf("服务注册失败:%v", err)
	}

	// 启动服务
	if err := greeterService.Run(); err != nil {
		logrus.Panic(fmt.Sprintln("启动微服务失败:", err))
	}
}
