package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	service "kitexSvr-serviceA-2/kitex_gen/kitex/service/hertzsvr"
	"log"
	"net"
)

func InitEtcdRegistry(s *HertzSvrImpl, serviceName string, addr *net.TCPAddr) server.Server {
	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		log.Fatal("Error: fail to new etcd registry---" + err.Error())
	}

	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: serviceName,
		Tags:        map[string]string{"Cluster": "AServiceCluster"},
	}

	svr := service.NewServer(s, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(ebi))

	return svr
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9991")
	s := new(HertzSvrImpl)

	svr := InitEtcdRegistry(s, "AService", addr)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
