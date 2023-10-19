package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"projeto.com/src/user/model"
)

type Pix struct {
	gorm.Model
	ID          string          `gorm:"primaryKey;type:varchar(255);column:id"`
	User        model.User      `gorm:"type:varchar(255);column:user"`
	BankAccount BankAccount     `gorm:"type:varchar(255);column:bank_account"`
	Balance     decimal.Decimal `gorm:"type:varchar(255);column:balance"`
}
