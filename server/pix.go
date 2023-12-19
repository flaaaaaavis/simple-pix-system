package server

import (
	gofrs "github.com/gofrs/uuid"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"log"
	pb "mentoria/protobuf/pix/v1"
	"mentoria/src/pix/model"
	"mentoria/src/pix/service"
	"time"
)

// PixServer communication interface between bff and proto
type PixServer struct {
	proto          pb.UnimplementedPixServiceServer
	bankAccountSvc service.BankAccountRepo
	pixSvc         service.PixRepo
	pixCodeSvc     service.PixCodeRepo
	transactionSvc service.TransactionRepo
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

	res, err := s.bankAccountSvc.CreateBankAccount(newBankAccount)
	if err != nil {
		log.Fatalf("Error creating new bankAccount, %v", err)
		return nil, err
	}

	bankAccount := model.FromBankAccountModelToProto(res)

	return bankAccount, nil
}

// GetBankAccountById sends bff's bank account obtaining request to service through proto
func (s *PixServer) GetBankAccountById(id string) (*pb.BankAccount, error) {
	res, err := s.bankAccountSvc.GetBankAccountById(id)
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
	userId := gofrs.FromStringOrNil(proto.UserId)
	bankAccountId := gofrs.FromStringOrNil(proto.BankAccountId)
	balance, err := decimal.NewFromString(proto.Balance)

	newPix := &model.Pix{
		UserID:        uuid.UUID(userId),
		BankAccountID: uuid.UUID(bankAccountId),
		Balance:       balance,
	}

	res, err := s.pixSvc.CreatePix(newPix)
	if err != nil {
		log.Fatalf("Error creating new pix, %v", err)
		return nil, err
	}

	pix := model.FromPixModelToProto(res)

	return pix, nil
}

// GetPixById sends bff's pix obtaining request to service through proto
func (s *PixServer) GetPixById(id string) (*pb.Pix, error) {
	res, err := s.pixSvc.GetPixById(id)
	if err != nil {
		log.Fatalf("Error getting pix by id, %v", err)
		return nil, err
	}

	pix := model.FromPixModelToProto(res)

	return pix, nil
}

// UpdatePixBalance sends bff's pix update balance request to service through proto
func (s *PixServer) UpdatePixBalance(proto *pb.Pix) (*pb.Pix, error) {
	id := gofrs.FromStringOrNil(proto.Id)
	balance, err := decimal.NewFromString(proto.Balance)

	newPix := &model.Pix{
		ID:      uuid.UUID(id),
		Balance: balance,
	}

	res, err := s.pixSvc.UpdatePixBalance(newPix)
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
	pixId := gofrs.FromStringOrNil(proto.PixId)

	newPixCode := &model.PixCode{
		PixID: uuid.UUID(pixId),
		Type:  model.PixType(proto.Type),
		Code:  proto.Code,
	}

	res, err := s.pixCodeSvc.CreatePixCode(newPixCode)
	if err != nil {
		log.Fatalf("Error creating new pix code, %v", err)
		return nil, err
	}

	pixCode := model.FromPixCodeModelToProto(res)

	return pixCode, nil
}

// GetPixCodeByPixId sends bff's pix code obtaining request to service through proto
func (s *PixServer) GetPixCodeByPixId(id string) (*pb.PixCode, error) {
	res, err := s.pixCodeSvc.GetPixCodeByPixId(id)
	if err != nil {
		log.Fatalf("Error getting pix code by pix id, %v", err)
		return nil, err
	}

	pixCode := model.FromPixCodeModelToProto(res)

	return pixCode, nil
}

// GetPixCodeByCode sends bff's pix code obtaining request to service through proto
func (s *PixServer) GetPixCodeByCode(code string) (*pb.PixCode, error) {
	res, err := s.pixCodeSvc.GetPixCodeByCode(code)
	if err != nil {
		log.Fatalf("Error getting pix code by code, %v", err)
		return nil, err
	}

	pixCode := model.FromPixCodeModelToProto(res)

	return pixCode, nil
}

// DeletePixCode sends bff's pix code deletion request to service through proto
func (s *PixServer) DeletePixCode(code string) error {
	err := s.pixCodeSvc.DeletePixCode(code)
	if err != nil {
		log.Fatalf("Error deleting pix code, %v", err)
		return err
	}

	return nil
}

// Transaction

// CreateTransaction sends bff's transaction creation request to service through proto
func (s *PixServer) CreateTransaction(proto *pb.Transaction) (*pb.Transaction, error) {
	senderId := gofrs.FromStringOrNil(proto.SenderId)
	receiverId := gofrs.FromStringOrNil(proto.ReceiverId)
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
		SenderID:   uuid.UUID(senderId),
		ReceiverID: uuid.UUID(receiverId),
	}

	res, err := s.transactionSvc.CreateTransaction(newTransaction)
	if err != nil {
		log.Fatalf("Error creating new transaction, %v", err)
		return nil, err
	}

	transaction := model.FromTransactionModelToProto(res)

	return transaction, nil
}

// ListUserTransactionsById sends bff's user's transactions listing request to service through proto
func (s *PixServer) ListUserTransactionsById(id string) ([]*pb.Transaction, error) {
	res, err := s.transactionSvc.ListUserTransactionsById(id)
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
