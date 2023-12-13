package main

import (
	"fmt"
	"log"
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

	_, err := config.CreateTables(cfg.DynamoDBConfig)
	if err != nil {
		fmt.Println("erro aqui")
		log.Fatalf("error %v", err)
	}

	/*log.Printf("retorno %s", res.DescribeTable)*/

	server.NewGRPC()
	log.Println("Listening on port 9003")

	log.Fatalln("terminou")
}
