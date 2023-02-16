package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sakasikai/GoFive/cmd/api/rpc"
	"github.com/sakasikai/GoFive/cmd/user/dal/db"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive/userservice"
	"github.com/sakasikai/GoFive/pkg/constants"
	"log"
	"net"
)

func Init() {
	rpc.InitRPC()
	db.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	Init()

	svr := gofive.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithRegistry(r),                                             // registry
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
