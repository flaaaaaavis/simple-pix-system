package model

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type User struct {
	ID         string  `json:"id"`
	FullName   string  `json:"full_name"`
	SocialName string  `json:"social_name"`
	CPF        string  `json:"cpf"`
	Contact    Contact `json:"contact"`
}
