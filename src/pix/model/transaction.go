package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type TransactionType string

const (
	TransactionTypeCharge TransactionType = "CHARGE"

	TransactionTypePayment TransactionType = "PAYMENT"

	/* TransactionTypeRefund TransactionType = "REFUND" */
)

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "PENDING"

	TransactionStatusDone TransactionStatus = "DONE"
)

type Transaction struct {
	Type     TransactionType    `json:"type"`
	Date     time.Time          `json:"date"`
	Amount   decimal.Decimal    `json:"amount"`
	Accounts map[string]Account `json:"sender"`
	Status   TransactionStatus  `json:"status"`
}
