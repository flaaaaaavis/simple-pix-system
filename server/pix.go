package server

import (
	"github.com/shopspring/decimal"
	"log"
	pb "mentoria/protobuf/pix/v1"
	"mentoria/src/pix/model"
	"mentoria/src/pix/service"
	"time"
)

// PixServer communication interface between bff and proto
type PixServer struct {
	pb.UnimplementedPixServiceServer
	BankAccountSvc service.BankAccountService
	PixSvc         service.PixService
	PixCodeSvc     service.PixCodeService
	TransactionSvc service.TransactionService
}

// BankAccount

// CreateBankAccount sends bff's bank account creation request to service through proto
func (s *PixServer) CreateBankAccount(proto *pb.BankAccount) (*pb.BankAccount, error) {
	newBankAccount := &model.BankAccount{
		BankCode:      proto.BankCode,
		BankName:      proto.BankName,
		BankBranch:    proto.BankBranch,
		AccountNumber: proto.AccountNumber,
	}

	res, err := s.BankAccountSvc.CreateBankAccount(newBankAccount)
	if err != nil {
		log.Fatalf("Error creating new bankAccount, %v", err)
		return nil, err
	}

	bankAccount := model.FromBankAccountModelToProto(res)

	return bankAccount, nil
}

// GetBankAccountById sends bff's bank account obtaining request to service through proto
func (s *PixServer) GetBankAccountById(id string) (*pb.BankAccount, error) {
	res, err := s.BankAccountSvc.GetBankAccountById(id)
	if err != nil {
		log.Fatalf("Error getting bankAccount by id, %v", err)
		return nil, err
	}

	bankAccount := model.FromBankAccountModelToProto(res)

	return bankAccount, nil
}

// Pix

// CreatePix sends bff's pix creation request to service through proto
func (s *PixServer) CreatePix(proto *pb.Pix) (*pb.Pix, error) {
	balance, err := decimal.NewFromString(proto.Balance)

	newPix := &model.Pix{
		UserID:        proto.UserId,
		BankAccountID: proto.BankAccountId,
		Balance:       balance,
	}

	res, err := s.PixSvc.CreatePix(newPix)
	if err != nil {
		log.Fatalf("Error creating new pix, %v", err)
		return nil, err
	}

	pix := model.FromPixModelToProto(res)

	return pix, nil
}

// GetPixById sends bff's pix obtaining request to service through proto
func (s *PixServer) GetPixById(id string) (*pb.Pix, error) {
	res, err := s.PixSvc.GetPixById(id)
	if err != nil {
		log.Fatalf("Error getting pix by id, %v", err)
		return nil, err
	}

	pix := model.FromPixModelToProto(res)

	return pix, nil
}

// UpdatePixBalance sends bff's pix update balance request to service through proto
func (s *PixServer) UpdatePixBalance(proto *pb.Pix) (*pb.Pix, error) {
	balance, err := decimal.NewFromString(proto.Balance)

	newPix := &model.Pix{
		ID:      proto.Id,
		Balance: balance,
	}

	res, err := s.PixSvc.UpdatePixBalance(newPix)
	if err != nil {
		log.Fatalf("Error updating pix balance, %v", err)
		return nil, err
	}

	pix := model.FromPixModelToProto(res)

	return pix, nil
}

// PixCode

// CreatePixCode sends bff's pix code creation request to service through proto
func (s *PixServer) CreatePixCode(proto *pb.PixCode) (*pb.PixCode, error) {
	newPixCode := &model.PixCode{
		PixID: proto.PixId,
		Type:  model.PixType(proto.Type),
		Code:  proto.Code,
	}

	res, err := s.PixCodeSvc.CreatePixCode(newPixCode)
	if err != nil {
		log.Fatalf("Error creating new pix code, %v", err)
		return nil, err
	}

	pixCode := model.FromPixCodeModelToProto(res)

	return pixCode, nil
}

// GetPixCodeByPixId sends bff's pix code obtaining request to service through proto
func (s *PixServer) GetPixCodeByPixId(id string) (*pb.PixCode, error) {
	res, err := s.PixCodeSvc.GetPixCodeByPixId(id)
	if err != nil {
		log.Fatalf("Error getting pix code by pix id, %v", err)
		return nil, err
	}

	pixCode := model.FromPixCodeModelToProto(res)

	return pixCode, nil
}

// GetPixCodeByCode sends bff's pix code obtaining request to service through proto
func (s *PixServer) GetPixCodeByCode(code string) (*pb.PixCode, error) {
	res, err := s.PixCodeSvc.GetPixCodeByCode(code)
	if err != nil {
		log.Fatalf("Error getting pix code by code, %v", err)
		return nil, err
	}

	pixCode := model.FromPixCodeModelToProto(res)

	return pixCode, nil
}

// DeletePixCode sends bff's pix code deletion request to service through proto
func (s *PixServer) DeletePixCode(code string) error {
	err := s.PixCodeSvc.DeletePixCode(code)
	if err != nil {
		log.Fatalf("Error deleting pix code, %v", err)
		return err
	}

	return nil
}

// Transaction

// CreateTransaction sends bff's transaction creation request to service through proto
func (s *PixServer) CreateTransaction(proto *pb.Transaction) (*pb.Transaction, error) {
	amount, err := decimal.NewFromString(proto.Amount)

	layout := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(layout, proto.Date)
	if err != nil {
		log.Fatalf("Error parsing new transaction date, %v", err)
		return nil, err
	}

	newTransaction := &model.Transaction{
		Date:       date,
		Amount:     amount,
		SenderID:   proto.SenderId,
		ReceiverID: proto.ReceiverId,
	}

	res, err := s.TransactionSvc.CreateTransaction(newTransaction)
	if err != nil {
		log.Fatalf("Error creating new transaction, %v", err)
		return nil, err
	}

	transaction := model.FromTransactionModelToProto(res)

	return transaction, nil
}

// ListUserTransactionsById sends bff's user's transactions listing request to service through proto
func (s *PixServer) ListUserTransactionsById(id string) ([]*pb.Transaction, error) {
	res, err := s.TransactionSvc.ListUserTransactionsById(id)
	if err != nil {
		log.Fatalf("Error getting user transactions, %v", err)
		return nil, err
	}

	var transactions []*pb.Transaction

	for i := 0; i < len(res); i++ {
		transactions = append(transactions, model.FromTransactionModelToProto(&res[i]))
	}

	return transactions, nil
}
