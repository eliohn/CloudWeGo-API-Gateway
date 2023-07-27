package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	service "kitexSvr-SecondLevelCalService/kitex_gen/kitex/service/secondlevelcalservice"
	"log"
	"net"
)

func InitEtcdRegistry(s *SecondLevelCalServiceImpl, serviceName string, addr *net.TCPAddr) server.Server {
	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		log.Fatal("Error: fail to new etcd registry---" + err.Error())
	}

	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: serviceName,
		Tags:        map[string]string{"Cluster": serviceName + "Cluster"},
	}

	svr := service.NewServer(s, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(ebi))

	return svr
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9992")
	s := new(SecondLevelCalServiceImpl)

	svr := InitEtcdRegistry(s, "SecondLevelCalService", addr)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
