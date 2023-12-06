package main

import (
	"fmt"
	"log"

	"mentoria/src/config"
)

func main() {
	cfg := *config.NewConfig()

	// _, err := config.Connection(cfg.Type)

	// if err != nil {
	// 	log.Fatalf("erro ao setar nova conexão com o GORM: %s", err)
	// }

	// cl := *config.ConnectionDynamo()

	// _ = dynamo.NewDynamoClient(cl)

	res, err := config.CreateTables(cfg.DynamoDBConfig)
	if err != nil {
		fmt.Println("erro aqui")
		log.Fatalf("error %v", err)
	}

	log.Println("retorno %s", res.DescribeTable)

	log.Fatalln("terminou")

}
