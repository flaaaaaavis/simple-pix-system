package model

import (
	"github.com/google/uuid"
)

// BankAccount model
type BankAccount struct {
	ID            uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	BankCode      string    `json:"bank_code"`
	BankName      string    `json:"bank_name"`
	BankBranch    string    `json:"bank_branch"`
	AccountNumber string    `json:"account_number"`
}
