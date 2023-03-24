package main

import (
	imooc "cap-imooc/proto/cap"
	"context"
	"fmt"

	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Name("cap.imooc.client"),
	)
	service.Init()
	capImooc := imooc.NewCapService("cap.imooc.server", service.Client())

	res, err := capImooc.SayHello(context.TODO(), &imooc.SayRequest{Message: "跟着cap老师学微服务"})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Answer)
}
