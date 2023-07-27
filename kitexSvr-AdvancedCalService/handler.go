package main

import (
	"context"
	service "kitexSvr-AdvancedCalService/kitex_gen/kitex/service"
	"log"
)

// AdvancedCalServiceImpl implements the last idlManager interface defined in the IDL.
type AdvancedCalServiceImpl struct{}

// Fact implements the AdvancedCalServiceImpl interface.
func (s *AdvancedCalServiceImpl) Fact(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received factorial request--By AdvancedCal Service")
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
			Message: "AdvancedCal Service Fact method success",
			Data:    int32(calculateFact(int(n))),
		}
	}
	return
}

// Fib implements the AdvancedCalServiceImpl interface.
func (s *AdvancedCalServiceImpl) Fib(ctx context.Context, request *service.Request) (resp *service.Response, err error) {
	log.Println("Received fibonacci request--By AdvancedCal Service")
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
			Message: "AdvancedCal Service Fib method success",
			Data:    int32(calculateFib(int(n))),
		}
	}
	return
}

// calculateFact Helper function to calculate factorial
func calculateFact(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFact(n-1)
}

// calculateFib Helper function to calculate fibonacci
func calculateFib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return calculateFib(n-1) + calculateFib(n-2)
}
