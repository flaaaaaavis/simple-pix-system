package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mentoria/src/pix/model"
	userModel "mentoria/src/user/model/postgres_model"
)

func Connection(config DatabaseConfig) (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Errorf("erro ao abrir a configuração do GORM: %s", err)
		return nil, err
	}

	db.AutoMigrate(&userModel.Contact{}, &userModel.User{}, &model.BankAccount{}, &model.PixCode{}, &model.Pix{}, &model.Transaction{})

	//contactId := uuid.New()
	//db.Create(&userModel.Contact{
	//	ID:          contactId,
	//	PhoneNumber: "(12)12345-6789",
	//	Email:       "email@email.com",
	//})
	//
	//contactId2 := uuid.New()
	//db.Create(&userModel.Contact{
	//	ID:          contactId2,
	//	PhoneNumber: "(12)12345-6787",
	//	Email:       "email2@email.com",
	//})
	//
	//userId := uuid.New()
	//db.Create(&userModel.User{
	//	ID:         userId,
	//	FullName:   "FullName",
	//	SocialName: "SocialName",
	//	CPF:        "CPF",
	//	ContactID:  contactId,
	//})
	//
	//userId2 := uuid.New()
	//db.Create(&userModel.User{
	//	ID:         userId2,
	//	FullName:   "FullName2",
	//	SocialName: "SocialName2",
	//	CPF:        "CPF2",
	//	ContactID:  contactId2,
	//})
	//
	//bankAccountId := uuid.New()
	//db.Create(&model.BankAccount{
	//	ID:            bankAccountId,
	//	BankCode:      "001",
	//	BankName:      "Banco do Brasil",
	//	BankBranch:    "4498",
	//	AccountNumber: "1138233-3",
	//})
	//
	//bankAccountId2 := uuid.New()
	//db.Create(&model.BankAccount{
	//	ID:            bankAccountId2,
	//	BankCode:      "001",
	//	BankName:      "Banco do Brasil",
	//	BankBranch:    "4498",
	//	AccountNumber: "1138273-3",
	//})
	//
	//pixId := uuid.New()
	//db.Create(&model.Pix{
	//	ID:            pixId,
	//	UserID:        userId,
	//	BankAccountID: bankAccountId,
	//	Balance:       decimal.NewFromFloat(7543.28),
	//})
	//
	//pixId2 := uuid.New()
	//db.Create(&model.Pix{
	//	ID:            pixId2,
	//	UserID:        userId2,
	//	BankAccountID: bankAccountId2,
	//	Balance:       decimal.NewFromFloat(7543.28),
	//})
	//
	//pixCodeId := uuid.New()
	//db.Create(&model.PixCode{
	//	ID:    pixCodeId,
	//	PixID: pixId,
	//	Type:  model.PixTypePhone,
	//	Code:  "(87)99619-7228",
	//})
	//
	//pixCodeId2 := uuid.New()
	//db.Create(&model.PixCode{
	//	ID:    pixCodeId2,
	//	PixID: pixId2,
	//	Type:  model.PixTypePhone,
	//	Code:  "(87)99619-7229",
	//})
	//
	//transactionId := uuid.New()
	//db.Create(&model.Transaction{
	//	ID:         transactionId,
	//	Date:       time.Now(),
	//	Amount:     decimal.NewFromFloat(500.00),
	//	SenderID:   pixId,
	//	ReceiverID: pixId2,
	//})
	return db, nil
}
