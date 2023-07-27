package main

import (
	"context"
	service "kitexSvr-SecondLevelCalService/kitex_gen/kitex/service"
	"log"
)

// SecondLevelCalServiceImpl implements the last service interface defined in the IDL.
type SecondLevelCalServiceImpl struct{}

// Mul implements the SecondLevelCalServiceImpl interface.
func (s *SecondLevelCalServiceImpl) Mul(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received mul request--By SecondLevelCal Service")
	result := request.Operand_1 * request.Operand_2
	resp = &service.Response{
		Success: true,
		Message: "SecondLevelCal Service Mul method success",
		Data:    result,
	}
	return
}

// Div implements the SecondLevelCalServiceImpl interface.
func (s *SecondLevelCalServiceImpl) Div(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received div request--By SecondLevelCal Service")
	var result int32
	if request.Operand_2 == 0 {
		resp = &service.Response{
			Success: false,
			Message: "Zero div error occurred",
			Data:    -1,
		}
	} else {
		result = request.Operand_1 / request.Operand_2
		resp = &service.Response{
			Success: true,
			Message: "B Service Div method success",
			Data:    result,
		}
	}
	return
}
