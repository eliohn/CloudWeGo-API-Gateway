package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertzSvr/biz/handler/hertzSvr/utils"
	"hertzSvr/biz/model/hertzSvr/service"
)

func Gateway(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.SvrRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 从路由中获取服务名
	svcName := c.Param("svc")
	clientInfo := clients[svcName]

	resp := utils.GetHTTPGenericResponse(ctx, c, "", clientInfo.cli)

	c.JSON(consts.StatusOK, resp)
}
