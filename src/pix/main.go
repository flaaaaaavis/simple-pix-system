package main

import (
	"google.golang.org/grpc"
	"log"
	pb "mentoria/protobuf/pix/v1"
	"mentoria/server"
	"mentoria/src/config"
	"mentoria/src/pix/repository"
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

	newRepoBankAccount := repository.NewBankAccount(conn)
	newRepoTransaction := repository.NewTransaction(conn)
	newRepoPixCode := repository.NewPixCode(conn)
	newRepoPix := repository.NewPix(conn)
	newGrpc := grpc.NewServer()

	newServerPix := server.PixServer{
		BankAccountSvc: newRepoBankAccount,
		PixSvc:         newRepoPix,
		PixCodeSvc:     newRepoPixCode,
		TransactionSvc: newRepoTransaction,
	}

	pb.RegisterPixServiceServer(newGrpc, &newServerPix)

	err = newGrpc.Serve(listener)
}
