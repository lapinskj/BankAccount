package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	ent "BankAccount/entities"
	rint "BankAccount/repositories/interfaces"
	respErr "BankAccount/respErr"
)

type BankAccountHandler struct {
	Repository rint.IAccountRepository
}

func NewBankAccountHandler(rep rint.IAccountRepository) *BankAccountHandler {
	return &BankAccountHandler{Repository: rep}
}

func (handler *BankAccountHandler) Serve(w http.ResponseWriter,
	r *http.Request) {
	err := handler.Handler(w, r)
	if err.Err != nil {
		http.Error(w, err.Error(), err.Status())
	}
	w.WriteHeader(err.Status())
}

func (handler *BankAccountHandler) Handler(w http.ResponseWriter,
	r *http.Request) respErr.RespErr {

	var err respErr.RespErr
	switch r.Method {
	case "GET":
		err = handler.GetHandler(w, r)
	case "POST":
		err = handler.PostHandler(w, r)
	}

	return err
}

func (handler *BankAccountHandler) GetHandler(w http.ResponseWriter,
	r *http.Request) respErr.RespErr {
	fName := "handlers.BankAccountHandler.GetHandler"

	params := mux.Vars(r)
	idStr, ok := params["ID"]
	if !ok {
		return respErr.RespErr{errors.Wrap(errors.New("No ID parameter in the request"),
			fName), 500}
	}
	accountId, cErr := strconv.Atoi(idStr)
	if cErr != nil {
		return respErr.RespErr{errors.Wrap(cErr, fName), 500}
	}
	accountId32 := int32(accountId)

	account, err := handler.Repository.Read(accountId32)
	if err != nil {
		return respErr.RespErr{errors.Wrap(err, fName), 500}
	}

	json, jErr := json.Marshal(account)
	if jErr != nil {
		return respErr.RespErr{errors.Wrap(jErr, fName), 500}
	}

	fmt.Fprint(w, string(json))

	return respErr.RespErr{nil, 200}
}

func (handler *BankAccountHandler) PostHandler(w http.ResponseWriter,
	r *http.Request) respErr.RespErr {
	fName := "handlers.BankAccountHandler.PostHandler"

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return respErr.RespErr{errors.Wrap(err, fName), 500}
	}

	var bankAccount ent.BankAccount
	jErr := json.Unmarshal(body, &bankAccount)
	if jErr != nil {
		w.WriteHeader(500)
		return respErr.RespErr{errors.Wrap(err, fName), 500}
	}

	sErr := handler.Repository.Add(bankAccount)
	if sErr != nil {
		w.WriteHeader(500)
		return respErr.RespErr{errors.Wrap(err, fName), 500}
	}

	return respErr.RespErr{nil, 201}
}
