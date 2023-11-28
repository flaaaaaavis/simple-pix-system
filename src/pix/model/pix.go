package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Pix struct {
	ID            uuid.UUID       `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	UserID        uuid.UUID       `gorm:"type:uuid;column:user_id"`
	BankAccountID uuid.UUID       `gorm:"type:uuid;column:bank_account_id"`
	Balance       decimal.Decimal `gorm:"type:uuid;column:balance"`
}
