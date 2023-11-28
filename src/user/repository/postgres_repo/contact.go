package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mentoria/src/user/model/postgres_model"
	"mentoria/src/user/service"
)

type contactRepository struct {
	gormConnection *gorm.DB
}

func (c contactRepository) CreateContact(newContact *model.Contact) (*model.Contact, error) {
	err := c.gormConnection.Create(newContact)
	if err.Error != nil {
		fmt.Sprintf("Error when creating new contact: %s", err.Error)

		return nil, err.Error
	}

	return newContact, nil
}

func (c contactRepository) GetContactById(id string) (*model.Contact, error) {
	contact := &model.Contact{}

	condition := fmt.Sprintf("id=%s", id)

	result := c.gormConnection.First(model.Contact{}, condition)
	if result.Error != nil {
		fmt.Sprintf("Error when getting contact: %s", result.Error)

		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		fmt.Sprintf("Error when getting contact: %s", err.Error())

		return nil, errors.New(err.Error())
	}

	err = result.ScanRows(rows, contact)
	if err != nil {
		fmt.Sprintf("Error when getting contact: %s", err.Error())

		return nil, errors.New(err.Error())
	}

	return contact, nil
}

func (c contactRepository) UpdateContactById(newContact *model.Contact) (*model.Contact, error) {
	err := c.gormConnection.Model(newContact).Where("id IN (?)", newContact.ID).Updates(newContact)
	if err.Error != nil {
		fmt.Sprintf("Error when updating contact: %s", err.Error)

		return nil, err.Error
	}

	return newContact, nil
}

func NewContact(db *gorm.DB) service.ContactRepo {
	return &contactRepository{
		gormConnection: db,
	}
}
