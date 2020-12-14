package controller

import (
	"context"
	"fmt"

	pb "github.com/mariogao/microservice/proto"
)

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (h *HelloService) MyHelloService(ctx context.Context, req *pb.HelloReq, res *pb.HelloRes) error {
	fmt.Println("传入的结构体：", req.Name)
	//具体逻辑，假装查了一下 数据库
	res.ResName = "你好" + req.Name
	fmt.Println("返回的结构体：", res.ResName)
	return nil
}
