package main

import (
	"fmt"
	"projeto.com/src/config"
	"projeto.com/src/repository"
)

// consumir o NewConnection e subir o DB no docker compose

func main() {
	cfg := config.NewConfig()

	_, err := repository.NewConnection(cfg.Type)

	if err != nil {
		fmt.Errorf("erro ao setar nova conexão com o GORM: %s", err)
	}
}
