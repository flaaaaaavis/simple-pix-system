package dynamo_model

type Pix struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	BankAccountID string `json:"bank_account_id"`
	Balance       string `json:"balance"`
}
