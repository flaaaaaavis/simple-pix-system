package main

import (
	"fmt"
	"mentoria/src/config"
)

func main() {
	cfg := config.NewConfig()

	_, err := config.NewConnection(cfg.Type)

	if err != nil {
		fmt.Errorf("erro ao setar nova conexão com o GORM: %s", err)
	}
}
