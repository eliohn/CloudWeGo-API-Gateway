package main

import (
	"context"
	service "kitexSvr-serviceD/kitex_gen/kitex/service"
	"log"
)

// DServiceImpl implements the last service interface defined in the IDL.
type DServiceImpl struct{}

// RequestA implements the DServiceImpl interface.
func (s *DServiceImpl) RequestA(ctx context.Context, req *service.DReq) (resp *service.DResp, err error) {
	// TODO: Your code here...
	log.Println("Server D receive request: " + req.Data)
	resp = &service.DResp{
		Success: true,
		Message: "D success",
	}
	return
}
