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
	log.Printf("client grpc is running  on port: 9003")

	newUser := pb.NewUserServiceClient(grpcClient)
	backendUser := user.NewGRPC(newUser)
	userRouter := user.NewRouter(backendUser)

	var routes []router.Route
	routes = append(routes, userRouter.Routes()...)

	if err = Listen("9002", routes...); err != nil {
		log.Fatalf("Failed to listen server on port 9002: %v", err)
	}
}

func Listen(port string, routes ...router.Route) error {
	r := chi.NewRouter()

	for _, route := range routes {
		m := route.Method()
		p := route.Path()
		r.Method(m, p, route.Handler())
	}
	log.Printf("Serve is running  on port: %v", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
