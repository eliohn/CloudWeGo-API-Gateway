package main

import (
	"context"
	service "kitexSvr-serviceC/kitex_gen/kitex/service"
	"log"
)

// HertzSvrImpl implements the last service interface defined in the IDL.
type HertzSvrImpl struct{}

// Request implements the HertzSvrImpl interface.
func (s *HertzSvrImpl) Request(ctx context.Context, request *service.SvrRequest) (resp *service.SvrResponse, err error) {
	// TODO: Your code here...
	log.Println("Server C receive request: " + request.BizParams)
	resp = &service.SvrResponse{
		Success: true,
		Message: "C success",
	}
	return
}
