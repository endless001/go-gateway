package grpc

import "google.golang.org/grpc"

func NewGrpcLoadBalanceHandler() grpc.StreamHandler {
	return func() grpc.StreamHandler {
		return nil
	}()
}
