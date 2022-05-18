package server

import (
	"crypto/tls"
	rt "github.com/go-serv/service/internal/runtime"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type endpoint struct {
	srv                       serverInterface
	lis                       net.Listener
	grpcSrv                   *grpc.Server
	GrpcSrvUnaryInterceptors  []grpc.UnaryServerInterceptor
	GrpcSrvStreamInterceptors []grpc.StreamServerInterceptor
}

func (e *endpoint) withServer(s serverInterface) {
	e.srv = s
}

func (e *endpoint) GrpcServer() *grpc.Server {
	return e.grpcSrv
}

type tcpEndpoint struct {
	endpoint
	hostname      string
	port          int
	httpTransport bool
	tlsCfg        *tls.Config
}

func (e *tcpEndpoint) serveInit() {
	// Create a new gRPC server
	e.grpcSrv = grpc.NewServer(e.srv.GrpcServerOptions()...)
	// Register all network gRPC services
	for _, svc := range rt.Runtime().NetworkServices() {
		svc.Service_Register(e.grpcSrv)
	}
}

func (e *tcpEndpoint) Address() string {
	addr := e.hostname + ":" + strconv.Itoa(e.port)
	return addr
}

func (e *tcpEndpoint) listen(network string) error {
	if !e.httpTransport {
		addr := e.Address()
		lis, err := net.Listen(network, addr)
		if err != nil {
			return err
		}
		e.lis = lis
	} else {
		lis, err := tls.Listen(network, e.Address(), e.tlsCfg)
		if err != nil {
			return err
		}
		e.lis = lis
	}
	return nil
}

func (e *tcpEndpoint) tcpServe() error {
	e.serveInit()
	if !e.httpTransport {
		if err := e.grpcSrv.Serve(e.lis); err != nil {
			return err
		}
	} else {

	}
	return nil
}

//
// TCP 4 endpoint
//
type tcp4Endpoint struct {
	tcpEndpoint
}

func (e *tcp4Endpoint) Listen() error {
	return e.listen(Tcp4Network)
}

//
// TCP 6 endpoint
//
type tcp6Endpoint struct {
	tcpEndpoint
	fallback *tcp4Endpoint
}

func (e *tcp6Endpoint) Listen() error {
	return e.listen(Tcp6Network)
}

//func (e *endpoint) netServeInit() {
//	for _, svc := range rt.Runtime().NetworkServices() {
//		svc.Service_Register(e.grpcSrv)
//	}
//}
//
//func (e *localEndpoint) unixServe() error {
//	e.serveInit()
//	if err := e.grpcSrv.Serve(e.lis); err != nil {
//		return err
//	}
//	return nil
//}

//
// Local endpoint
//
type localEndpoint struct {
	endpoint
	pathname string
}

func (e *localEndpoint) Address() string {
	return e.pathname
}