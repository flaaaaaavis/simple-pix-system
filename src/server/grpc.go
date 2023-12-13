package server

import (
	"google.golang.org/grpc"
	"log"
	pb "mentoria/protobuf/user/v1"
	"mentoria/server"
	"net"
)

func NewGRPC() (*grpc.Server, error) {
	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", "localhost:9003")
	if err != nil {
		log.Fatalf("error creating new tcp listener: %v", err)
		return nil, err
	}

	userSvc := server.UserServer{}

	pb.RegisterUserServiceServer(grpcServer, &userSvc)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("error serving new tcp listener, %v", err)
		return nil, err
	}

	return grpcServer, nil
}
