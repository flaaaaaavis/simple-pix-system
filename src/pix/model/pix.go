package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PixType string

const (
	PixTypeEmail PixType = "EMAIL"

	PixTypePhone PixType = "PHONE"

	PixTypeCPF PixType = "CPF"

	PixTypeRandom PixType = "RANDOM"
)

type BankAccount struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	BankCode     string    `json:"bank_code,omitempty"`
	BankName     string    `json:"bank_name,omitempty"`
	BankBranch   string    `json:"bank_branch,omitempty"`
	Account      string    `json:"account,omitempty"`
	Code         string    `json:"code,omitempty"`
	Type         string    `json:"type,omitempty"`
	Date         time.Time `gorm:"type:date"`
	Status       string    `json:"status,omitempty"`
	Amount       int       `json:"amount,omitempty"`
	Balance      int       `json:"balance,omitempty"`
	Transactions string    `json:"transactions,omitempty"`
}

type TransactionType string

const (
	TransactionTypeCharge TransactionType = "CHARGE"

	TransactionTypePayment TransactionType = "PAYMENT"

	/* TransactionTypeRefund TransactionType = "REFUND" */
)

type Account struct {
	gorm.Model
	ID          int         `gorm:"primarykey;type:int:;default:gen_random_uuid()"`
	BankAccount BankAccount `gorm:"foreignKey:ID;references:ID" `
}
