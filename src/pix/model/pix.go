package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type BankAccount struct {
	BankCode   string
	BankName   string
	BankBranch string
	Account    string
}

type PixType string

const (
	PixTypeEmail PixType = "EMAIL"

	PixTypePhone PixType = "PHONE"

	PixTypeCPF PixType = "CPF"

	PixTypeRandom PixType = "RANDOM"
)

type PixCode struct {
	Type PixType
	Code string
}

type Pix struct {
	ID          string
	BankAccount BankAccount
	PixCodes    []PixCode
}

type TransactionType string

const (
	TransactionTypeCharge TransactionType = "CHARGE"

	TransactionTypePayment TransactionType = "PAYMENT"

	/* TransactionTypeRefund TransactionType = "REFUND" */
)

type Transaction struct {
	Type      TransactionType
	Date      time.Time
	Status    string
	Amount    decimal.Decimal
	Sender    BankAccount
	Recipient BankAccount
}

type Account struct {
	BankAccount BankAccount
	Balance     decimal.Decimal
	Statement   []Transaction
}
