package main

import (
	"context"
	service "kitexSvr-serviceB/kitex_gen/kitex/service"
	"log"
)

// BServiceImpl implements the last service interface defined in the IDL.
type BServiceImpl struct{}

// RequestB implements the BServiceImpl interface.
func (s *BServiceImpl) RequestB(ctx context.Context, req *service.BReq) (resp *service.BResp, err error) {
	// TODO: Your code here...
	log.Println("Server B receive request: " + req.Data)
	resp = &service.BResp{
		Success: true,
		Message: "B success",
	}
	return
}
