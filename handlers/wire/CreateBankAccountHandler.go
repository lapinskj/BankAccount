//+build wireinject

package handlers

import (
	h "BankAccount/handlers"
	rimp "BankAccount/repositories/implementations"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func CreateBankAccountHandler(dbConn *gorm.DB) (*h.BankAccountHandler, error) {
	panic(wire.Build(
		rimp.NewAccountRepository,
		h.NewBankAccountHandler,
	))
}
