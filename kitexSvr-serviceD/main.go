package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	service "kitexSvr-serviceD/kitex_gen/kitex/service/dservice"
	"log"
	"net"
)

func InitEtcdRegistry(s *DServiceImpl, serviceName string, addr *net.TCPAddr) server.Server {
	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		log.Fatal("Error: fail to new etcd registry---" + err.Error())
	}

	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: serviceName,
		Tags:        map[string]string{"Cluster": "DServiceCluster"},
	}

	svr := service.NewServer(s, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(ebi))

	return svr
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9993")
	s := new(DServiceImpl)

	svr := InitEtcdRegistry(s, "DService", addr)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
