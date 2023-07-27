package main

import (
	"context"
	service "kitexSvr-serviceA-1/kitex_gen/kitex/service"
	"log"
)

// HertzSvrImpl implements the last service interface defined in the IDL.
type HertzSvrImpl struct{}

// Add implements the HertzSvrImpl interface.
func (s *HertzSvrImpl) Add(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received add request--By A Service")
	result := request.Operand_1 + request.Operand_2
	resp = &service.Response{
		Success: true,
		Message: "A Service Add method success",
		Data:    result,
	}
	return
}

// Sub implements the HertzSvrImpl interface.
func (s *HertzSvrImpl) Sub(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received sub request--By A Service")
	result := request.Operand_1 - request.Operand_2
	resp = &service.Response{
		Success: true,
		Message: "A Service Sub method success",
		Data:    result,
	}
	return
}
