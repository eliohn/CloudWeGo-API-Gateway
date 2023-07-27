package main

import (
	"context"
	"kitexSvr-FirstLevelCalService/kitex_gen/kitex/service"
	"log"
)

// FirstLevelCalServiceImpl implements the last service interface defined in the IDL.
type FirstLevelCalServiceImpl struct{}

// Add implements the FirstLevelCalServiceImpl interface.
func (s *FirstLevelCalServiceImpl) Add(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received add request--By FirstLevelCal Service")
	result := request.Operand_1 + request.Operand_2
	resp = &service.Response{
		Success: true,
		Message: "FirstLevelCal Service Add method success",
		Data:    result,
	}
	return
}

// Sub implements the FirstLevelCalServiceImpl interface.
func (s *FirstLevelCalServiceImpl) Sub(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received sub request--By FirstLevelCal Service")
	result := request.Operand_1 - request.Operand_2
	resp = &service.Response{
		Success: true,
		Message: "FirstLevelCal Service Sub method success",
		Data:    result,
	}
	return
}
