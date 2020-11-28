package grpc

import (
	"context"
	"github.com/mwitkow/grpc-proxy/proxy"
	"go-gateway/internal/loadbalancer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func NewGrpcLoadBalanceHandler(loadBalance loadbalancer.LoadBalance) grpc.StreamHandler {
	return func() grpc.StreamHandler {
		nextAddr, err := loadBalance.Get("")
		if err != nil {
			log.Fatal("get next addr fail")
		}
		director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
			c, err := grpc.DialContext(ctx, nextAddr, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
			md, _ := metadata.FromIncomingContext(ctx)
			outCtx, _ := context.WithCancel(ctx)
			outCtx = metadata.NewIncomingContext(outCtx, md.Copy())
			return outCtx, c, err
		}
		return proxy.TransparentHandler(director)
	}()
}
