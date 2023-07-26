package service

import (
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"hertzSvr/biz/handler/hertzSvr/utils"
)

type ClientInfo struct {
	provider *generic.ThriftContentProvider
	cli      genericclient.Client
}

var idlContent = `
namespace go kitex.service

struct SvrRequest{
    1: string svrName (api.body="svrName")
    2: string bizParams (api.body="bizParams") //应该是JSON，用string代替
}

struct RegisterIDL{
    1: string name (api.body="name")
    2: string version(api.body="version")
    3: string idl (api.body="idl")
}

struct SvrResponse{
    1: bool success (api.body="success")
    2: string message (api.body="message") //应该是JSON，用string代替
}

service HertzSvr{
    SvrResponse Request(1: SvrRequest request)(api.post="/gateway/:svr/request")
    SvrResponse RegisterIDL(1: RegisterIDL idl)(api.post="/registerIDL")
}
`

// 初始化etcdresolver
var resolver = utils.NewResolver()

// 初始化三个服务对应的provider
var providerA = utils.NewProvider(idlContent)
var providerB = utils.NewProvider(idlContent)
var providerC = utils.NewProvider(idlContent)

// 初始化clientInfo
var clients = map[string]ClientInfo{
	"AService": {
		provider: providerA,
		cli:      utils.NewClient("AService", providerA, resolver),
	},

	"BService": {
		provider: providerB,
		cli:      utils.NewClient("BService", providerB, resolver),
	},

	"CService": {
		provider: providerC,
		cli:      utils.NewClient("CService", providerC, resolver),
	},
}
