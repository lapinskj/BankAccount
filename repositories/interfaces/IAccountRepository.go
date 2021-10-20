package interfaces

import (
	ent "BankAccount/entities"
)

type IAccountRepository interface {
	Add(account ent.BankAccount) error
	Delete(accountId int32) error
	Read(accountId int32) (ent.BankAccount, error)
	ReadAll() ([]ent.BankAccount, error)
	Update(accountId ent.BankAccount) error
}