package types

// CreateUserRequest model
type CreateUserRequest struct {
	FullName   string `json:"full_name"`
	SocialName string `json:"social_name"`
	CPF        string `json:"cpf"`
	ContactID  string `json:"contact_id"`
}

// GetUserByIdRequest model
type GetUserByIdRequest struct {
	ID string `json:"id"`
}

// UpdateUserRequest model
type UpdateUserRequest struct {
	ID         string `json:"id"`
	FullName   string `json:"full_name"`
	SocialName string `json:"social_name"`
	CPF        string `json:"cpf"`
	ContactID  string `json:"contact_id"`
}

// UserResponse model
type UserResponse struct {
	ID         string `json:"id"`
	FullName   string `json:"full_name"`
	SocialName string `json:"social_name"`
	CPF        string `json:"cpf"`
	ContactID  string `json:"contact_id"`
}

// CreateContactRequest model
type CreateContactRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

// GetContactByIdRequest model
type GetContactByIdRequest struct {
	ID string `json:"id"`
}

// UpdateContactByIdRequest model
type UpdateContactByIdRequest struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

// ContactResponse model
type ContactResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
