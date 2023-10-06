package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"projeto.com/src/user/model"
)

type PixType string

const (
	PixTypeEmail PixType = "EMAIL"

	PixTypePhone PixType = "PHONE"

	PixTypeCPF PixType = "CPF"

	PixTypeRandom PixType = "RANDOM"
)

type PixCode struct {
	ID   uint    `json:"id" gorm:"foreignKey:PixCodeID"`
	Type PixType `json:"type"`
	Code string  `json:"code"`
}

type Account struct {
	gorm.Model
	ID            uint            `json:"id" gorm:"primaryKey"`
	User          model.User      `json:"user"`
	BankAccountID uint            `json:"bank_account"`
	PixCodeID     []uint          `json:"pix_codes"`
	Balance       decimal.Decimal `json:"balance"`
	Statement     []Transaction   `json:"statement"`
}
