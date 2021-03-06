package net

import (
	net_mw "github.com/go-serv/service/internal/grpc/middleware/net"
	enc_mw "github.com/go-serv/service/internal/middleware/net/enc_msg"
	session_mw "github.com/go-serv/service/internal/middleware/net/session"
	"github.com/go-serv/service/internal/server"
	_ "github.com/go-serv/service/internal/service/net_parcel/server" // this will enable the NetParcel service
	"github.com/go-serv/service/pkg/z"
)

func NewNetServer() *netServer {
	srv := new(netServer)
	srv.ServerInterface = server.NewServer()
	srv.ServerInterface.WithMiddlewareGroup(srv.defaultMiddlewareGroup())
	return srv
}

func (srv *netServer) defaultMiddlewareGroup() z.MiddlewareInterface {
	g := net_mw.NewMiddleware()
	session_mw.ServerInit(g)
	enc_mw.ServerInit(g)
	return g
}
