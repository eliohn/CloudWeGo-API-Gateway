package service

import (
	"github.com/cloudwego/kitex/client/genericclient"
	"hertzSvr/biz/handler/hertzSvr/utils"
)

type ClientInfo struct {
	name string
	cli  genericclient.Client
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
    SvrResponse Request(1: SvrRequest request)(api.post="/request")
    SvrResponse RegisterIDL(1: RegisterIDL idl)(api.post="/registerIDL")
}
`

// 初始化etcdresolver
var resolver = utils.NewResolver()

// 初始化四个服务对应的client（后续会用缓存层优化）
var cliA = utils.NewClient("AService", utils.NewProvider(idlContent), resolver)
var cliB = utils.NewClient("BService", utils.NewProvider(idlContent), resolver)
var cliC = utils.NewClient("CService", utils.NewProvider(idlContent), resolver)
var cliD = utils.NewClient("DService", utils.NewProvider(idlContent), resolver)

var clients = map[string]genericclient.Client{
	"AService": cliA,
	"BService": cliB,
	"CService": cliC,
	"DService": cliD,
}
