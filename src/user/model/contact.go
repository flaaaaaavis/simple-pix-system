package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID          string `gorm:"primaryKey;type:varchar(255);column:id"`
	PhoneNumber string `gorm:"type:varchar(255);column:phone_number"`
	Email       string `gorm:"type:varchar(255);column:email"`
}
