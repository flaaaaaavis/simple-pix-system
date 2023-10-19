package model

import (
	"gorm.io/gorm"
)

type PixType string

const (
	PixTypeEmail PixType = "EMAIL"

	PixTypePhone PixType = "PHONE"

	PixTypeCPF PixType = "CPF"

	PixTypeRandom PixType = "RANDOM"
)

type PixCode struct {
	gorm.Model
	ID    string  `gorm:"primaryKey;type:varchar(255);column:id"`
	PixID string  `gorm:"type:varchar(255);column:pix_id"`
	Type  PixType `gorm:"type: varchar(255);column:type"`
	Code  string  `gorm:"type: varchar(255);column:code"`
}
