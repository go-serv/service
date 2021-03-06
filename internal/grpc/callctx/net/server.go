package net

import (
	"github.com/go-serv/service/internal/grpc/codec"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type srvContext struct {
	netContext
	server  z.NetworkServerInterface
	sess    z.SessionInterface
	handler grpc.UnaryHandler
}

func (ctx *srvContext) Invoke() (err error) {
	var (
		ifreplay any
	)
	if ifreplay, err = ctx.handler(ctx, ctx.req.DataFrame().Interface()); err != nil {
		return
	}
	msg := ifreplay.(proto.Message)
	ctx.res.WithDataFrame(codec.NewDataFrame(msg))
	err = ctx.res.Populate(msg)
	return
}

func (s *srvContext) Session() z.SessionInterface {
	return s.sess
}

func (s *srvContext) WithSession(sess z.SessionInterface) {
	s.sess = sess
}

func (s *srvContext) Server() z.NetworkServerInterface {
	return s.server
}

func (s *srvContext) WithServer(srv z.NetworkServerInterface) {
	s.server = srv
}
