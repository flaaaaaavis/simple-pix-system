package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "mentoria/protos/protobuf/user/v1"
	"mentoria/src/config"
	"mentoria/src/server"
)

func main() {
	cfg := *config.NewConfig()

	// _, err := config.Connection(cfg.Type)

	// if err != nil {
	// 	log.Fatalf("erro ao setar nova conexão com o GORM: %s", err)
	// }

	// cl := *config.ConnectionDynamo()

	// _ = dynamo.NewDynamoClient(cl)

	lis, err := net.Listen("tcp", "localhost:9003")
	if err != nil {
		log.Fatalf("error on start rpc Server, %v", err)
	}

	usr := server.UserService{}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &usr)

	log.Println("Listening on port 9003")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed server %v", err)
	}

	_, err = config.CreateTables(cfg.DynamoDBConfig)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	log.Fatalln("terminou")

}
