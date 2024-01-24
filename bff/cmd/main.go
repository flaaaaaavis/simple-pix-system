package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"mentoria/bff/server/router"
	"mentoria/bff/server/user"
	pb "mentoria/protobuf/user/v1"
	"net/http"
)

func main() {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	grpcClient, err := grpc.Dial("localhost:9003", creds)
	if err != nil {
		log.Fatal("error on connection client grpc ")
	}

	newUser := pb.NewUserServiceClient(grpcClient)
	newGrpcBackend := user.NewGrpcBackend(newUser)

	userRouter := user.NewRouter(newGrpcBackend)

	var routes []router.Route

	routes = append(routes, userRouter.Routes()...)

	ServerListener("9002", routes...)
}

func ServerListener(port string, routes ...router.Route) error {
	chiRouter := chi.NewRouter()

	for _, item := range routes {
		method := item.Method()
		path := item.Path()
		handler := item.Handler()

		chiRouter.Method(method, path, handler)
	}

	http.ListenAndServe(port, chiRouter)
	log.Printf("Serve is running  on port: %v", port)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), chiRouter)
}
