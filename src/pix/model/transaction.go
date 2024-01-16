package model

import (
	"github.com/shopspring/decimal"
	pb "mentoria/protobuf/pix/v1"
	"time"
)

// Transaction model
type Transaction struct {
	ID         string          `gorm:"primarykey;type:uuid;default:gen_random_uuid();column:id"`
	Date       time.Time       `gorm:"type:date;column:date"`
	Amount     decimal.Decimal `gorm:"type:decimal(15,2); column:amount"`
	SenderID   string          `gorm:"type:uuid;column:sender_id"`
	ReceiverID string          `gorm:"type:uuid;column:receiver_id"`
}

// FromTransactionModelToProto converts Transaction from model to proto
func FromTransactionModelToProto(transaction *Transaction) *pb.Transaction {
	return &pb.Transaction{
		Id:         transaction.ID,
		Date:       transaction.Date.String(),
		Amount:     transaction.Amount.String(),
		SenderId:   transaction.SenderID,
		ReceiverId: transaction.ReceiverID,
	}
}
