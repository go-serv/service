package net

import (
	"context"
	net_req "github.com/go-serv/service/internal/grpc/request/net"
	net_res "github.com/go-serv/service/internal/grpc/response/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (mw *netMiddleware) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		var (
			wrappedReq z.RequestInterface
			wrappedRes z.ResponseInterface
		)
		clnt := mw.Client()
		wrappedReq, err = net_req.NewClientRequest(req, clnt.Meta(), nil)
		if err != nil {
			return
		}
		md := metadata.MD{}
		opts = append(opts, grpc.Header(&md))
		wrappedRes = net_res.NewResponse(reply, md)
		netCtx := ctx.(z.NetClientContextInterface)
		netCtx.WithClientInvoker(invoker, cc, opts)
		netCtx.WithRequest(wrappedReq)
		netCtx.WithResponse(wrappedRes)
		// Request chain
		_, err = mw.newRequestChain().passThrough(netCtx)
		if err != nil {
			return
		}
		err = wrappedRes.Meta().Hydrate()
		if err != nil {
			return
		}
		// Response chain
		_, err = mw.newResponseChain().passThrough(wrappedRes)
		return
	}
}