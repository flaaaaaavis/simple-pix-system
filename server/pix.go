package server

import (
	gofrs "github.com/gofrs/uuid"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"log"
	pb "mentoria/protobuf/pix/v1"
	"mentoria/src/pix/model"
	"mentoria/src/pix/service"
	"mentoria/utils"
)

type PixServer struct {
	proto          pb.UnimplementedPixServiceServer
	bankAccountSvc service.BankAccountRepo
	pixSvc         service.PixRepo
	pixCodeSvc     service.PixCodeRepo
	transactionSvc service.TransactionRepo
}

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

	bankAccount := utils.FromBankAccountModelToProto(res)

	return bankAccount, nil
}

func (s *PixServer) GetBankAccountById(id string) (*pb.BankAccount, error) {
	res, err := s.bankAccountSvc.GetBankAccountById(id)
	if err != nil {
		log.Fatalf("Error getting bankAccount by id, %v", err)
		return nil, err
	}

	bankAccount := utils.FromBankAccountModelToProto(res)

	return bankAccount, nil
}

// Pix

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

	pix := utils.FromPixModelToProto(res)

	return pix, nil
}

func (s *PixServer) GetPixById(id string) (*pb.Pix, error) {
	res, err := s.pixSvc.GetPixById(id)
	if err != nil {
		log.Fatalf("Error getting pix by id, %v", err)
		return nil, err
	}

	pix := utils.FromPixModelToProto(res)

	return pix, nil
}

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

	pix := utils.FromPixModelToProto(res)

	return pix, nil
}
