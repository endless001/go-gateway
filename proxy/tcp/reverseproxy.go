package tcp

import (
	"context"
	"go-gateway/loadbalancer"
	"io"
	"log"
	"net"
	"time"
)

type TcpReverseProxy struct {
	ctx                  context.Context
	Addr                 string
	KeepAlivePeriod      time.Duration
	DialTimeout          time.Duration
	DialContext          func(ctx context.Context, network, address string) (net.Conn, error)
	OnDialError          func(src net.Conn, dstDialErr error)
	ProxyProtocolVersion int
}

func NewLoadBalanceReverseProxy(balance loadbalancer.LoadBalance) *TcpReverseProxy {
	return func() *TcpReverseProxy {
		nextAddr, err := balance.Get("")
		if err != nil {
			log.Fatal("get next addr fail")
		}
		return &TcpReverseProxy{
			Addr:            nextAddr,
			KeepAlivePeriod: time.Second,
			DialTimeout:     time.Second,
		}
	}()
}

func (t *TcpReverseProxy) dialTimeOut() time.Duration {

	if t.DialTimeout > 0 {
		return t.DialTimeout
	}
	return 10 * time.Second

}

var defaultDialer = new(net.Dialer)

func (t *TcpReverseProxy) dialContext() func(ctx context.Context, network, address string) (net.Conn, error) {
	if t.DialContext != nil {
		return t.DialContext
	}
	return (&net.Dialer{
		Timeout:   t.DialTimeout,
		KeepAlive: t.KeepAlivePeriod,
	}).DialContext
}

func (t *TcpReverseProxy) keepAlivePeriod() time.Duration {
	if t.KeepAlivePeriod != 0 {
		return t.KeepAlivePeriod
	}
	return time.Minute
}

func (t *TcpReverseProxy) ServeTCP(ctx context.Context, src net.Conn) {

}

func (t *TcpReverseProxy) onDialError() func(src net.Conn, dstDialErr error) {
	if t.OnDialError != nil {
		return t.OnDialError
	}
	return func(src net.Conn, dstDialErr error) {
		log.Printf("tcpproxy for incomming conn %v,error dialing %q: %v",
			src.RemoteAddr().String(), t.Addr, dstDialErr)
		src.Close()
	}
}

func (t *TcpReverseProxy) proxyCopy(errc chan<- error, dst, src net.Conn) {
	_, err := io.Copy(dst, src)
	errc <- err
}
