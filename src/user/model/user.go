package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         string  `gorm:"primaryKey;type:varchar(255);column:id"`
	FullName   string  `gorm:"type: varchar(255); column:full_name"`
	SocialName string  `gorm:"type: varchar(255); column:social_name"`
	CPF        string  `gorm:"type: varchar(255); column:cpf"`
	Contact    Contact `gorm:"foreignKey:ID"`
}
