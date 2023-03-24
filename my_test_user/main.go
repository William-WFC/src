package main

import (
	"my_test_user/handler"
	pb "my_test_user/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("my_test_user"),
	)

	// Register handler
	pb.RegisterMy_test_userHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
