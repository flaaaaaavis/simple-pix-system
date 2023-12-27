package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mentoria/src/pix/model"
	model2 "mentoria/src/user/postgres/model"
)

func Connection(config DatabaseConfig) (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Errorf("erro ao abrir a configuração do GORM: %s", err)
		return nil, err
	}

	db.AutoMigrate(&model2.Contact{}, &model2.User{}, &model.BankAccount{}, &model.PixCode{}, &model.Pix{}, &model.Transaction{})

	return db, nil
}
