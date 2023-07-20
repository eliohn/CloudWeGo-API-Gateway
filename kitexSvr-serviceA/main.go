package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	service "kitexSvr-serviceA/kitex_gen/kitex/service/aservice"
	"log"
	"net"
)

func InitEtcdRegistry(s *AServiceImpl, serviceName string, addr *net.TCPAddr) server.Server {
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
	addr, _ := net.ResolveTCPAddr("tcp", ":9990")
	s := new(AServiceImpl)

	svr := InitEtcdRegistry(s, "AService", addr)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
