package main

import (
	"context"
	"kitexSvr-serviceB/kitex_gen/kitex/service"
	"log"
)

// HertzSvrImpl implements the last service interface defined in the IDL.
type HertzSvrImpl struct{}

// Mul implements the HertzSvrImpl interface.
func (s *HertzSvrImpl) Mul(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received mul request--By B Service")
	result := request.Operand_1 * request.Operand_2
	resp = &service.Response{
		Success: true,
		Message: "B Service Mul method success",
		Data:    result,
	}
	return
}

// Div implements the HertzSvrImpl interface.
func (s *HertzSvrImpl) Div(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received div request--By B Service")
	result := request.Operand_1 / request.Operand_2
	resp = &service.Response{
		Success: true,
		Message: "B Service Div method success",
		Data:    result,
	}
	return
}
