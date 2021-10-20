package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	rint "BankAccount/repositories/interfaces"
	respErr "BankAccount/respErr"
)

type BankAccountsHandler struct {
	Repository rint.IAccountRepository
}

func NewBankAccountsHandler(rep rint.IAccountRepository) *BankAccountsHandler {
	return &BankAccountsHandler{Repository: rep}
}

func (handler *BankAccountsHandler) Serve(w http.ResponseWriter,
	r *http.Request) {
	err := handler.Handler(w, r)
	if err.Err != nil {
		http.Error(w, err.Error(), err.Status())
	}
	w.WriteHeader(err.Status())
}

func (handler *BankAccountsHandler) Handler(w http.ResponseWriter,
	r *http.Request) respErr.RespErr {

	var err respErr.RespErr
	switch r.Method {
	case "GET":
		err = handler.GetHandler(w, r)
	}

	return err
}

func (handler *BankAccountsHandler) GetHandler(w http.ResponseWriter,
	r *http.Request) respErr.RespErr {
	fName := "handlers.BankAccountsHandler.GetHandler"

	accounts, err := handler.Repository.ReadAll()
	if err != nil {
		return respErr.RespErr{errors.Wrap(err, fName), 500}
	}

	json, jErr := json.Marshal(accounts)
	if jErr != nil {
		return respErr.RespErr{errors.Wrap(jErr, fName), 500}
	}

	fmt.Fprint(w, string(json))

	return respErr.RespErr{nil, 200}
}
