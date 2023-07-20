package main

import (
	"context"
	service "kitexSvr-serviceA/kitex_gen/kitex/service"
	"log"
)

// AServiceImpl implements the last service interface defined in the IDL.
type AServiceImpl struct{}

// RequestA implements the AServiceImpl interface.
func (s *AServiceImpl) RequestA(ctx context.Context, req *service.AReq) (resp *service.AResp, err error) {
	// TODO: Your code here...
	log.Println("Server A receive request: " + req.Data)
	resp = &service.AResp{
		Success: true,
		Message: "A success",
	}
	return
}
