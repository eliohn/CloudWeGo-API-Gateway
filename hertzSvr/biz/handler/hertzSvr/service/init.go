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

var idlContentA = `
namespace go kitex.service

struct Request{
    1: i32 operand_1 (api.body="operand_1")
    2: i32 operand_2 (api.body="operand_2")
}

struct Response{
    1: bool success (api.body="success")
    2: string message (api.body="message") //应该是JSON，用string代替
    3: i32 data (api.body="data")
}

service HertzSvr{
    Response Add(1: Request request)(api.post="/gateway/AService/add")
    Response Sub(1: Request request)(api.post="/gateway/AService/sub")
}`
var idlContentB = `namespace go kitex.service

struct Request{
    1: i32 operand_1 (api.body="operand_1")
    2: i32 operand_2 (api.body="operand_2")
}

struct Response{
    1: bool success (api.body="success")
    2: string message (api.body="message") //应该是JSON，用string代替
    3: i32 data (api.body="data")
}

service HertzSvr{
    Response Mul(1: Request request)(api.post="/gateway/BService/mul")
    Response Div(1: Request request)(api.post="/gateway/BService/div")
}`
var idlContentC = `namespace go kitex.service

struct Request{
    1: i32 operand (api.body="operand")
}

struct Response{
    1: bool success (api.body="success")
    2: string message (api.body="message") //应该是JSON，用string代替
    3: i32 data (api.body="data")
}

service HertzSvr{
    Response Fact(1: Request request)(api.post="/gateway/CService/fact")
    Response Fib(1: Request request)(api.post="/gateway/CService/fib")
}`

// 初始化etcdresolver
var resolver = utils.NewResolver()

// 初始化三个服务对应的provider
var providerA = utils.NewProvider(idlContentA)
var providerB = utils.NewProvider(idlContentB)
var providerC = utils.NewProvider(idlContentC)

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
