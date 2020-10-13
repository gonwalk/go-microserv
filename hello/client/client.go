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

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2379"}
	})
	// 创建一个新的服务
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter.Client"),
		micro.Registry(reg),
	)
	// 初始化
	service.Init()

	// 创建 Greeter 客户端
	greeter := proto.NewGreeterService("go.micro.srv.greeter", service.Client())

	// 远程调用 Greeter 服务的 Hello 方法
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "sym"})
	if err != nil {
		logrus.Errorln("远程调用greeter.Hello失败：" + err.Error())
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
