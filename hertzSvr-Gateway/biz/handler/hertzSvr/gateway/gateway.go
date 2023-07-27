package gateway

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertzSvr-Gateway/biz/handler/hertzSvr/service"
	"hertzSvr-Gateway/biz/handler/hertzSvr/utils"
)

// Gateway API gateway
func Gateway(ctx context.Context, c *app.RequestContext) {
	// 从路由中获取服务名
	svcName := c.Param("svc")
	// 获取服务对应的clientInfo
	clientInfo := service.Clients[svcName]

	// 进行HTTP泛化调用
	resp := utils.GetHTTPGenericResponse(ctx, c, "", clientInfo.Cli)

	c.JSON(consts.StatusOK, resp)
}
