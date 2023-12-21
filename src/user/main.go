package main

import (
	"google.golang.org/grpc"
	"log"
	pb "mentoria/protobuf/user/v1"
	"mentoria/server"
	"mentoria/src/config"
	"mentoria/src/user/postgres/repository"
	"net"
)

func main() {
	cfg := *config.NewConfig()

	conn, err := config.Connection(cfg.Type)

	listener, err := net.Listen("tcp", "localhost:9003")
	if err != nil {
		log.Fatalf("error creating new tcp listener: %v", err)
	}
	log.Println("Listening on port 9003")

	if err != nil {
		log.Fatalf("erro ao setar nova conexão com o GORM: %s", err)
	}

	newRepoUser := repository.NewUser(conn)
	newReposContact := repository.NewContact(conn)
	newGrpc := grpc.NewServer()

	newServerUser := server.UserServer{
		UserSvc:    newRepoUser,
		ContactSvc: newReposContact,
	}

	pb.RegisterUserServiceServer(newGrpc, &newServerUser)

	//_, err = config.CreateTables(cfg.DynamoDBConfig)
	//if err != nil {
	//	fmt.Println("erro aqui")
	//	log.Fatalf("error %v", err)
	//}

	err = newGrpc.Serve(listener)
}
