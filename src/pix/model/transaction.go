package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type TransactionType string

const (
	TransactionTypePayment TransactionType = "PAYMENT"

	TransactionTypeRefund TransactionType = "REFUND"
)

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "PENDING"

	TransactionStatusDone TransactionStatus = "DONE"
)

type Transaction struct {
	gorm.Model
	ID         string            `gorm:"primaryKey;type:varchar(255);column:id"`
	Type       TransactionType   `gorm:"column:type"`
	Date       time.Time         `gorm:"type:date;column:date"`
	Amount     decimal.Decimal   `gorm:"type:decimal(15,2); column:amount"`
	SenderID   string            `gorm:"type:varchar(255);column:sender_id"`
	ReceiverID string            `gorm:"type:varchar(255);column:receiver_id"`
	Status     TransactionStatus `gorm:"type:varchar(20);default:'pending';column:status"`
}
