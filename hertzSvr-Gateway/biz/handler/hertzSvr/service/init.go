package service

import (
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"hertzSvr-Gateway/biz/handler/hertzSvr/utils"
)

type ClientInfo struct {
	Provider *generic.ThriftContentProvider
	Cli      genericclient.Client
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

service FirstLevelCalService{
    Response Add(1: Request request)(api.post="/gateway/FirstLevelCalService/add")
    Response Sub(1: Request request)(api.post="/gateway/FirstLevelCalService/sub")
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

service SecondLevelCalService{
    Response Mul(1: Request request)(api.post="/gateway/SecondLevelCalService/mul")
    Response Div(1: Request request)(api.post="/gateway/SecondLevelCalService/div")
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

service AdvancedCalService{
    Response Fact(1: Request request)(api.post="/gateway/AdvancedCalService/fact")
    Response Fib(1: Request request)(api.post="/gateway/AdvancedCalService/fib")
}`

// 初始化etcdresolver
var resolver = utils.NewResolver()

// 初始化三个服务对应的provider
var firstLevelCalProvider = utils.NewProvider(idlContentA)
var secondLevelCalProvider = utils.NewProvider(idlContentB)
var advancedCalProvider = utils.NewProvider(idlContentC)

// 初始化clientInfo
var Clients = map[string]ClientInfo{
	"FirstLevelCalService": {
		Provider: firstLevelCalProvider,
		Cli:      utils.NewClient("FirstLevelCalService", firstLevelCalProvider, resolver),
	},

	"SecondLevelCalService": {
		Provider: secondLevelCalProvider,
		Cli:      utils.NewClient("SecondLevelCalService", secondLevelCalProvider, resolver),
	},

	"AdvancedCalService": {
		Provider: advancedCalProvider,
		Cli:      utils.NewClient("AdvancedCalService", advancedCalProvider, resolver),
	},
}
