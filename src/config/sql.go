package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (c *GormConnection) Connection(config DatabaseConfig) error {
	dns := fmt.Sprintf("Host=%s Port=%s User=%s Pass=%s DbName=%s SSLMode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Errorf("erro ao abrir a configuração do GORM: %s", err)
		return err
	}

	/*  db.AutoMigrate(&UserValue{}, &PixValue{})  */

	c.Gorm = db

	return nil
}

func NewConnection(dbData DatabaseConfig) (*GormConnection, error) {
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
