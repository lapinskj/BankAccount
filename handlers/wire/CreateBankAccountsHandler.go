//+build wireinject

package handlers

import (
	h "BankAccount/handlers"
	rimp "BankAccount/repositories/implementations"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func CreateBankAccountsHandler(dbConn *gorm.DB) (*h.BankAccountsHandler, error) {
	panic(wire.Build(
		rimp.NewAccountRepository,
		h.NewBankAccountsHandler,
	))
}
