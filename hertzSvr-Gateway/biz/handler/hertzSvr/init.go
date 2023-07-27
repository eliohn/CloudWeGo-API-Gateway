package hertzSvr

import (
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"hertzSvr-Gateway/biz/handler/hertzSvr/utils"
)

type ClientInfo struct {
	Provider *generic.ThriftContentProvider
	Cli      genericclient.Client
}

// Resolver 初始化etcdresolver
var Resolver = utils.NewResolver()

// Clients 初始化clientInfo
var Clients = make(map[string]ClientInfo)
