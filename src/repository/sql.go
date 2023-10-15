package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"projeto.com/src/config"
	pixModel "projeto.com/src/pix/model"
)

type Gorm interface {
	Db() *gorm.DB
}

type GormConnection struct {
	Gorm *gorm.DB
}

func (c *GormConnection) Db() *gorm.DB {
	return c.Gorm
}

var GlobalConfig *GormConnection

func (c *GormConnection) Connection(config config.DatabaseConfig) error {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Errorf("erro ao abrir a configuração do GORM: %s", err)
		return err
	}

	err = db.AutoMigrate(&pixModel.BankAccount{})
	if err != nil {
		return err
	}

	c.Gorm = db

	return nil
}

func NewConnection(dbData config.DatabaseConfig) (*GormConnection, error) {
	if GlobalConfig == nil {
		GlobalConfig = &GormConnection{}
		err := GlobalConfig.Connection(dbData)

		if err != nil {
			fmt.Errorf("erro ao criar nova conexão com o GORM: %s", err)
			return nil, err
		}
	}

	return GlobalConfig, nil
}
