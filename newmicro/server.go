package main

import (
	imooc "cap-imooc/proto/cap"
	"context"
	"fmt"

	"go-micro.dev/v4"
)

type CapServer struct{}

func (c *CapServer) SayHello(ctx context.Context, req *imooc.SayRequest, res *imooc.SayResponse) error {
	res.Answer = "我们的口号是：" + req.Message
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(
		micro.Name("cap.imooc.server"),
	)
	// 初始化方法
	service.Init()
	// 注册服务
	imooc.RegisterCapHandler(service.Server(), new(CapServer))
	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
