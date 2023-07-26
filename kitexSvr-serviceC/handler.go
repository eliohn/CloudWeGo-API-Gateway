package main

import (
	"context"
	service "kitexSvr-serviceC/kitex_gen/kitex/service"
)

// HertzSvrImpl implements the last service interface defined in the IDL.
type HertzSvrImpl struct{}

// Fact implements the HertzSvrImpl interface.
func (s *HertzSvrImpl) Fact(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	n := request.Operand
	if n < 0 {
		resp = &service.Response{
			Success: false,
			Message: "Factorial Error: n cannot be negative",
			Data:    -1,
		}
	} else {
		resp = &service.Response{
			Success: true,
			Message: "C Service Fact method success",
			Data:    int32(calculateFact(int(n))),
		}
	}
	return
}

// Fib implements the HertzSvrImpl interface.
func (s *HertzSvrImpl) Fib(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	n := request.Operand
	if n < 0 {
		resp = &service.Response{
			Success: false,
			Message: "Fibonacci Error: n cannot be negative",
			Data:    -1,
		}
	} else {
		resp = &service.Response{
			Success: true,
			Message: "C Service Fib method success",
			Data:    int32(calculateFib(int(n))),
		}
	}
	return
}

func calculateFact(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFact(n-1)
}

func calculateFib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return calculateFib(n-1) + calculateFib(n-2)
}
