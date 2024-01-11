package main

import (
	"google.golang.org/grpc"
	"log"
)

func main() {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	grpcClient, err := grpc.Dial("localhost:9003", creds)
	if err != nil {
		log.Fatal("error on connection client grpc ")
	}
}
