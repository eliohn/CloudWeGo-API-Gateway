package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	service "kitexSvr-AdvancedCalService/kitex_gen/kitex/service/advancedcalservice"
	"log"
	"net"
)

func InitEtcdRegistry(s *AdvancedCalServiceImpl, serviceName string, addr *net.TCPAddr) server.Server {
	//r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	// ...
	//consulClient, err := consulapi.NewClient(config)
	// ...
	r, err := consul.NewConsulRegister("10.3.5.103:8500")
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
	addr, _ := net.ResolveTCPAddr("tcp", ":9993")
	s := new(AdvancedCalServiceImpl)

	svr := InitEtcdRegistry(s, "AdvancedCalService", addr)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
