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
		provider, err := utils.NewProvider(idlContent)
		if err != nil {
			c.JSON(consts.StatusBadRequest, &hertzSvr.Response{
				Success: false,
				Message: "Error: fail to load idl for service " + svcName + "." + err.Error(),
			})
			return
		}
		clientInfo.Provider = provider
		clientInfo.Cli, err = utils.NewClient(svcName, provider, hertzSvr.Resolver)
		if err != nil {
			c.JSON(consts.StatusBadRequest, &hertzSvr.Response{
				Success: false,
				Message: "Error: fail to make new client for service " + svcName + "." + err.Error(),
			})
		}

		hertzSvr.Clients[svcName] = clientInfo
	}

	// 进行HTTP泛化调用
	resp, err := utils.GetHTTPGenericResponse(ctx, c, "", clientInfo.Cli)
	if err != nil {
		c.JSON(consts.StatusBadGateway, hertzSvr.Response{
			Success: false,
			Message: "Error: fail to generic call " + svcName + "." + err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, resp.Body)
}
