//
// Implementation of NetParcel network service along with its runtime registration.

package server

import (
	"context"
	"crypto/rand"
	i "github.com/go-serv/service/internal"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var Name = protoreflect.FullName(proto.NetParcel_ServiceDesc.ServiceName)

type netParcel struct {
	i.NetworkServiceInterface
	proto.NetParcelServer
}

func (s *netParcel) Service_Register(srv *grpc.Server) {
	proto.RegisterNetParcelServer(srv, s)
}

func (s *netParcel) Service_OnNewSession(req i.RequestInterface) error {
	return nil
}

func (s *netParcel) GetCryptoNonce(ctx context.Context, req *proto.CryptoNonce_Request) (*proto.CryptoNonce_Response, error) {
	res := &proto.CryptoNonce_Response{}
	nonce := make([]byte, req.GetLen())
	rand.Read(nonce)
	res.Nonce = nonce
	return res, nil
}
