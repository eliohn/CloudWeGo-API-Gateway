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
namespace go hertzSvr.service

//--------------------update request & response--------------
struct UpdateReq {
    1: string idl(api.body = 'idl')
}

struct UpdateResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//--------------------A request & response--------------
struct AReq {
    1: string data(api.body='data')
}

struct AResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//--------------------B request & response--------------
struct BReq {
    1: string data(api.body='data')
}

struct BResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//--------------------C request & response--------------
struct CReq {
    1: string data(api.body='data')
}

struct CResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//--------------------D request & response--------------
struct DReq {
    1: string data(api.body='data')
}

struct DResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------update service-------------------
service UpdateService {
    UpdateResp Update(1: UpdateReq req)(api.get = '/update')
}

//----------------------A service-------------------
service AService {
    AResp RequestA(1: AReq req)(api.post = '/A-req')
}

//----------------------B service-------------------
service BService {
    BResp RequestB(1: BReq req)(api.post = '/B-req')
}

//----------------------C service-------------------
service CService {
    CResp RequestC(1: CReq req)(api.post = '/C-req')
}

//----------------------D service-------------------
service DService {
    DResp RequestD(1: DReq req)(api.post = '/D-req')
}
`
var idlAContent = `
namespace go kitex.service

//--------------------A request & response--------------
struct AReq {
    1: string data(api.body='data')
}

struct AResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------A service-------------------
service AService {
    AResp RequestA(1: AReq req)(api.post = '/A-req')
}

`
var idlBContent = `
namespace go kitex.service

//--------------------B request & response--------------
struct BReq {
    1: string data(api.body='data')
}

struct BResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------A service-------------------
service BService {
    BResp RequestB(1: BReq req)(api.post = '/B-req')
}

`
var idlCContent = `
namespace go kitex.service

//--------------------C request & response--------------
struct CReq {
    1: string data(api.body='data')
}

struct CResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------A service-------------------
service CService {
    CResp RequestC(1: CReq req)(api.post = '/C-req')
}

`
var idlDContent = `
namespace go kitex.service

//--------------------D request & response--------------
struct DReq {
    1: string data(api.body='data')
}

struct DResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------D service-------------------
service DService {
    DResp RequestD(1: DReq req)(api.post = '/D-req')
}

`

var provider = utils.NewProvider(idlAContent)
var resolver = utils.NewResolver()
var clientInfo = ClientInfo{
	name: "AService",
	cli:  utils.NewClient("AService", provider, resolver),
}

var cliA = utils.NewClient("AService", utils.NewProvider(idlAContent), resolver)
var cliB = utils.NewClient("BService", utils.NewProvider(idlBContent), resolver)
var cliC = utils.NewClient("CService", utils.NewProvider(idlCContent), resolver)
var cliD = utils.NewClient("DService", utils.NewProvider(idlDContent), resolver)
