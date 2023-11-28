package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	ID         uuid.UUID         `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	Type       TransactionType   `gorm:"type:varchar(255);column:type"`
	Date       time.Time         `gorm:"type:date;column:date"`
	Amount     decimal.Decimal   `gorm:"type:decimal(15,2); column:amount"`
	SenderID   uuid.UUID         `gorm:"type:uuid;column:sender_id"`
	ReceiverID uuid.UUID         `gorm:"type:uuid;column:receiver_id"`
	Status     TransactionStatus `gorm:"type:varchar(20);default:'pending';column:status"`
}
