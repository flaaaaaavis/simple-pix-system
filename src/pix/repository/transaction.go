package service

import (
	"fmt"
	"projeto.com/src/pix/model"
	userModel "projeto.com/src/user/model"
)

func createTransaction() model.Transaction {
	return model.Transaction{}
}

func listUserTransactions(user userModel.User) []model.Transaction {
	fmt.Sprint(user.ID)

	return []model.Transaction{}
}

func getUserTransactions(user userModel.User) model.Transaction {
	fmt.Sprint(user.ID)

	return model.Transaction{}
}

func updateTransaction() model.Transaction {
	return model.Transaction{}
}
