package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	FullName   string    `gorm:"type: varchar(255); column:full_name"`
	SocialName string    `gorm:"type: varchar(255); column:social_name"`
	CPF        string    `gorm:"type: varchar(255); column:cpf"`
	ContactID  uuid.UUID `gorm:"type:uuid;column:contact_id"`
}
