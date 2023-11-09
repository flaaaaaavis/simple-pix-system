package main

import (
	"fmt"
	"projeto.com/src/config"
)

func main() {
	cfg := config.NewConfig()

	_, err := config.Connection(cfg.Type)

	if err != nil {
		fmt.Errorf("erro ao setar nova conexão com o GORM: %s", err)
	}
}
