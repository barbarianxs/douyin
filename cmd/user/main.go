package main

import (
	user "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user/userservice"
	server "github.com/cloudwego/kitex/server"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	
	
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.UserServiceAddr)
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
