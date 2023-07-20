package main

import (
	"context"
	service "kitexSvr-serviceC/kitex_gen/kitex/service"
	"log"
)

// CServiceImpl implements the last service interface defined in the IDL.
type CServiceImpl struct{}

// RequestA implements the CServiceImpl interface.
func (s *CServiceImpl) RequestA(ctx context.Context, req *service.CReq) (resp *service.CResp, err error) {
	// TODO: Your code here...
	log.Println("Server C receive request: " + req.Data)
	resp = &service.CResp{
		Success: true,
		Message: "C success",
	}
	return
}
