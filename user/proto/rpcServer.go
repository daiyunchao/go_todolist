package proto

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	UnimplementedRemoteUserServer
}

func (s *Server) GetUserInfo(ctx context.Context, request *RequestGetUserInfo) (*ResponseGetUserInfo, error) {
	return nil, nil
}

func (s *Server) GetUserInfoByNickname(ctx context.Context, request *RequestGetUserInfoByNickname) (*ResponseGetUserInfoByNickname, error) {
	return &ResponseGetUserInfoByNickname{
		Nickname: request.Nickname,
		Id:       "10001",
		Password: "11223344",
	}, nil
}

func (s *Server) RegisterServer(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcServer := grpc.NewServer()
	RegisterRemoteUserServer(rpcServer, &Server{})
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
