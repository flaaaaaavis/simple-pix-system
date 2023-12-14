package types

type User struct {
	FullName   string `json:"full_name"`
	SocialName string `json:"social_name"`
	Cpf        string `json:"cpf"`
	ContactID  string `json:"contact_id"`
}
