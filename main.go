package main

import (
	"github.com/mariogao/microservice/controller"

	pb "github.com/mariogao/microservice/proto"

	"github.com/micro/go-micro/v2"
)

func main() {
	server := micro.NewService(
		micro.Name("myMicro"),
	)
	server.Init()
	_ = pb.RegisterHelloHandler(server.Server(), controller.NewHelloService())

	server.Run()
}
