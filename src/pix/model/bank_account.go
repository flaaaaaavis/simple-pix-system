package model

import "gorm.io/gorm"

// BankAccount model
type BankAccount struct {
	gorm.Model
	ID         string `gorm:"primaryKey;type:varchar(255);column:id"`
	BankCode   string `json:"bank_code"`
	BankName   string `json:"bank_name"`
	BankBranch string `json:"bank_branch"`
	Account    string `json:"account"`
}
