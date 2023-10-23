package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Pix struct {
	gorm.Model
	ID            string          `gorm:"primaryKey;type:varchar(255);column:id"`
	UserID        string          `gorm:"type:varchar(255);column:user_id"`
	BankAccountID string          `gorm:"type:varchar(255);column:bank_account_id"`
	Balance       decimal.Decimal `gorm:"type:varchar(255);column:balance"`
}
