package gateway

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertzSvr-Gateway/biz/handler/hertzSvr"
	"hertzSvr-Gateway/biz/handler/hertzSvr/idlManager"
	"hertzSvr-Gateway/biz/handler/hertzSvr/utils"
)

// Gateway API gateway
func Gateway(ctx context.Context, c *app.RequestContext) {
	// 从路由中获取服务名
	svcName := c.Param("svc")

	// 获取服务对应的clientInfo
	var clientInfo hertzSvr.ClientInfo

	// 查找缓存中是否有对应service的client，若无则进行创建并放入缓存
	if _, isOk := hertzSvr.Clients[svcName]; isOk {
		clientInfo = hertzSvr.Clients[svcName]
	} else {
		// 调用idl管理平台API，查找对应service的idl
		idlContent := idlManager.GetIDLContent(svcName)
		if idlContent == "" {
			panic("Error: cannot find idl of service " + svcName)
		}
		provider := utils.NewProvider(idlContent)
		clientInfo = hertzSvr.ClientInfo{
			Provider: provider,
			Cli:      utils.NewClient(svcName, provider, hertzSvr.Resolver),
		}
		hertzSvr.Clients[svcName] = clientInfo
	}

	// 进行HTTP泛化调用
	resp := utils.GetHTTPGenericResponse(ctx, c, "", clientInfo.Cli)

	c.JSON(consts.StatusOK, resp.Body)
}
