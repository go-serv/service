package net

import (
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type clientCtx struct {
	netContext
	client  z.NetworkClientInterface
	cc      *grpc.ClientConn
	invoker grpc.UnaryInvoker
	opts    []grpc.CallOption
}

func (ctx *clientCtx) WithClientInvoker(invoker grpc.UnaryInvoker, cc *grpc.ClientConn, opts []grpc.CallOption) {
	ctx.invoker = invoker
	ctx.cc = cc
	ctx.opts = opts
}

func (ctx *clientCtx) Client() z.NetworkClientInterface {
	return ctx.client
}

func (ctx *clientCtx) WithClient(client z.NetworkClientInterface) {
	ctx.client = client
}

// Invoke prepares metadata and calls gRPC method.
func (ctx *clientCtx) Invoke() (err error) {
	var (
		md metadata.MD
	)
	if md, err = ctx.req.Meta().Dehydrate(); err != nil {
		return
	}
	ctx.Context = metadata.NewOutgoingContext(ctx.Context, md)
	methodName := ctx.req.MethodReflection().SlashFullName()
	if err = ctx.invoker(
		ctx,
		methodName,
		ctx.req.DataFrame(),
		ctx.res.DataFrame(),
		ctx.cc,
		ctx.opts...); err != nil {
		return
	}
	// Response meta data.
	err = ctx.res.Meta().Hydrate()
	return
}
