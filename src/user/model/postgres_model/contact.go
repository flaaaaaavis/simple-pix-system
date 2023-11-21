package model

import (
	"github.com/google/uuid"
)

type Contact struct {
	ID          uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	Email       string    `gorm:"type:varchar(255);column:email"`
	PhoneNumber string    `gorm:"type:varchar(255);column:phone_number"`
}
