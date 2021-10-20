package implementations

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	ent "BankAccount/entities"
	rint "BankAccount/repositories/interfaces"
)

type AccountRepository struct {
	DbConn *gorm.DB
}

var _ rint.IAccountRepository = &AccountRepository{}

func NewAccountRepository(dbConn *gorm.DB) rint.IAccountRepository {
	accountRep := AccountRepository{DbConn: dbConn}
	return &accountRep
}

func (rep *AccountRepository) Add(account ent.BankAccount) error {
	fName := "repositories.implementations.AccountRepository.Add"

	res := rep.DbConn.Create(&account)
	if res.Error != nil {
		return errors.Wrap(res.Error, fName)
	}

	return nil
}

func (rep *AccountRepository) Delete(accountId int32) error {
	fName := "repositories.implementations.AccountRepository.Delete"

	res := rep.DbConn.Delete(&ent.BankAccount{}, accountId)
	if res.Error != nil {
		return errors.Wrap(res.Error, fName)
	}

	return nil
}

func (rep *AccountRepository) Read(accountId int32) (ent.BankAccount, error) {
	fName := "repositories.implementations.AccountRepository.Read"

	var bankAccount ent.BankAccount
	res := rep.DbConn.First(&bankAccount, accountId)
	if res.Error != nil {
		return ent.BankAccount{}, errors.Wrap(res.Error, fName)
	}

	return ent.BankAccount{}, nil
}

func (rep *AccountRepository) ReadAll() ([]ent.BankAccount, error) {
	fName := "repositories.implementations.AccountRepository.Read"

	var bankAccounts []ent.BankAccount
	res := rep.DbConn.Find(&bankAccounts)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, fName)
	}

	return bankAccounts, nil
}

func (rep *AccountRepository) Update(account ent.BankAccount) error {
	fName := "repositories.implementations.AccountRepository.Update"

	res := rep.DbConn.Model(&ent.BankAccount{}).Updates(account)
	if res.Error != nil {
		return errors.Wrap(res.Error, fName)
	}

	return nil
}
