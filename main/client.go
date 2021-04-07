package main

import (
	"context"
	"demo1/v1/proto"
	"fmt"

	"github.com/micro/go-micro/v2"
)


func main() {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.Version("v1"))
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("greeter.service", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(),
		&proto.Request{Name: "Micro中国"},
		) // 声明version
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应结果
	fmt.Println(rsp.Msg)
}
