package model

type Contact struct {
	Phone string
	Email string
}

type User struct {
	ID         string
	FullName   string
	SocialName string
	Contact    Contact
}
