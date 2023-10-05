package model

// BankAccount model
type BankAccount struct {
	BankCode   string `json:"bank_code"`
	BankName   string `json:"bank_name"`
	BankBranch string `json:"bank_branch"`
	Account    string `json:"account"`
}
