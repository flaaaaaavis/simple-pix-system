package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"mentoria/src/user/postgres/model"
	"mentoria/src/user/postgres/service"
)

type contactRepository struct {
	gormConnection *gorm.DB
}

func (c contactRepository) CreateContact(_ context.Context, newContact *model.Contact) (*model.Contact, error) {
	err := c.gormConnection.Create(newContact)
	if err.Error != nil {
		log.Fatalf("Error when creating new contact: %s", err.Error)

		return nil, err.Error
	}

	return newContact, nil
}

func (c contactRepository) GetContactById(_ context.Context, req *model.GetContactByIdRequest) (*model.Contact, error) {
	contact := &model.Contact{}

	id := fmt.Sprintf("%v", req.ID)
	condition := &model.Contact{
		ID: id,
	}

	c.gormConnection.First(contact, condition)

	return contact, nil
}

func (c contactRepository) UpdateContactById(_ context.Context, newContact *model.Contact) (*model.Contact, error) {
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
