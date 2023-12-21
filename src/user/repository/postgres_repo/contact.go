package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"mentoria/src/user/model/postgres_model"
	"mentoria/src/user/service"
)

type contactRepository struct {
	gormConnection *gorm.DB
}

func (c contactRepository) CreateContact(ctx context.Context, newContact *model.Contact) (*model.Contact, error) {
	log.Println("entrou")
	err := c.gormConnection.Create(newContact)
	if err.Error != nil {
		log.Fatalf("Error when creating new contact: %s", err.Error)

		return nil, err.Error
	}

	return newContact, nil
}

func (c contactRepository) GetContactById(ctx context.Context, req *model.GetContactByIdRequest) (*model.Contact, error) {
	contact := &model.Contact{}

	condition := fmt.Sprintf("id=%s", req.ID)

	result := c.gormConnection.First(model.Contact{}, condition)
	if result.Error != nil {
		log.Fatalf("Error when getting contact: %s", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		log.Fatalf("Error when getting contact: %s", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, contact)
	if err != nil {
		log.Fatalf("Error when getting contact: %s", err.Error())

		return nil, errors.New(err.Error())
	}

	return contact, nil
}

func (c contactRepository) UpdateContactById(ctx context.Context, newContact *model.Contact) (*model.Contact, error) {
	err := c.gormConnection.Model(newContact).Where("id IN (?)", newContact.ID).Updates(newContact)
	if err.Error != nil {
		log.Fatalf("Error when updating contact: %s", err.Error)

		return nil, err.Error
	}

	return newContact, nil
}

func NewContact(db *gorm.DB) service.ContactService {
	return &contactRepository{
		gormConnection: db,
	}
}
