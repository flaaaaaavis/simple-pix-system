package model

import (
	"github.com/shopspring/decimal"
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
	Type PixType `json:"type"`
	Code string  `json:"code"`
}

type Account struct {
	User        model.User      `json:"user"`
	BankAccount BankAccount     `json:"bank_account"`
	PixCodes    []PixCode       `json:"pix_codes"`
	Balance     decimal.Decimal `json:"balance"`
	Statement   []Transaction   `json:"statement"`
}
