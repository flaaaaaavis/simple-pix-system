package model

import (
	"github.com/google/uuid"
)

type PixType string

const (
	PixTypeEmail PixType = "EMAIL"

	PixTypePhone PixType = "PHONE"

	PixTypeCPF PixType = "CPF"

	PixTypeRandom PixType = "RANDOM"
)

type PixCode struct {
	ID    uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	PixID uuid.UUID `gorm:"type:uuid;column:pix_id"`
	Type  PixType   `gorm:"type: varchar(255);column:type"`
	Code  string    `gorm:"type: varchar(255);column:code"`
}
