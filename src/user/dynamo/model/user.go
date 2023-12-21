package model

type User struct {
	ID         string `json:"id"`
	FullName   string `json:"full_name"`
	SocialName string `json:"social_name"`
	CPF        string `json:"cpf"`
	ContactID  string `json:"contact_id"`
}
