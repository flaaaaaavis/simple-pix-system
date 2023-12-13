package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID         uuid.UUID       `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	Date       time.Time       `gorm:"type:date;column:date"`
	Amount     decimal.Decimal `gorm:"type:decimal(15,2); column:amount"`
	SenderID   uuid.UUID       `gorm:"type:uuid;column:sender_id"`
	ReceiverID uuid.UUID       `gorm:"type:uuid;column:receiver_id"`
}
