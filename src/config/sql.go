package config

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	postgres2 "mentoria/src/pix/model/postgres"
	userModel "mentoria/src/user/model"
	"time"
)

func Connection(config DatabaseConfig) (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Errorf("erro ao abrir a configuração do GORM: %s", err)
		return nil, err
	}

	db.AutoMigrate(&userModel.Contact{}, &userModel.User{}, &postgres2.BankAccount{}, &postgres2.PixCode{}, &postgres2.Pix{}, &postgres2.Transaction{})

	contactId := uuid.New()
	db.Create(&userModel.Contact{
		ID:          contactId,
		PhoneNumber: "(12)12345-6789",
		Email:       "email@email.com",
	})

	contactId2 := uuid.New()
	db.Create(&userModel.Contact{
		ID:          contactId2,
		PhoneNumber: "(12)12345-6787",
		Email:       "email2@email.com",
	})

	userId := uuid.New()
	db.Create(&userModel.User{
		ID:         userId,
		FullName:   "FullName",
		SocialName: "SocialName",
		CPF:        "CPF",
		ContactID:  contactId,
	})

	userId2 := uuid.New()
	db.Create(&userModel.User{
		ID:         userId2,
		FullName:   "FullName2",
		SocialName: "SocialName2",
		CPF:        "CPF2",
		ContactID:  contactId2,
	})

	bankAccountId := uuid.New()
	db.Create(&postgres2.BankAccount{
		ID:            bankAccountId,
		BankCode:      "001",
		BankName:      "Banco do Brasil",
		BankBranch:    "4498",
		AccountNumber: "1138233-3",
	})

	bankAccountId2 := uuid.New()
	db.Create(&postgres2.BankAccount{
		ID:            bankAccountId2,
		BankCode:      "001",
		BankName:      "Banco do Brasil",
		BankBranch:    "4498",
		AccountNumber: "1138273-3",
	})

	pixId := uuid.New()
	db.Create(&postgres2.Pix{
		ID:            pixId,
		UserID:        userId,
		BankAccountID: bankAccountId,
		Balance:       decimal.NewFromFloat(7543.28),
	})

	pixId2 := uuid.New()
	db.Create(&postgres2.Pix{
		ID:            pixId2,
		UserID:        userId2,
		BankAccountID: bankAccountId2,
		Balance:       decimal.NewFromFloat(7543.28),
	})

	pixCodeId := uuid.New()
	db.Create(&postgres2.PixCode{
		ID:    pixCodeId,
		PixID: pixId,
		Type:  postgres2.PixTypePhone,
		Code:  "(87)99619-7228",
	})

	pixCodeId2 := uuid.New()
	db.Create(&postgres2.PixCode{
		ID:    pixCodeId2,
		PixID: pixId2,
		Type:  postgres2.PixTypePhone,
		Code:  "(87)99619-7229",
	})

	transactionId := uuid.New()
	db.Create(&postgres2.Transaction{
		ID:         transactionId,
		Type:       postgres2.TransactionTypePayment,
		Date:       time.Now(),
		Amount:     decimal.NewFromFloat(500.00),
		SenderID:   pixId,
		ReceiverID: pixId2,
		Status:     postgres2.TransactionStatusPending,
	})
	return db, nil
}
