package main

import (
	"log"
	"net"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user/userservice"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}), // server name
		server.WithServiceAddr(addr), server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
