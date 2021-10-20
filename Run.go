package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	config "BankAccount/configuration"
	wire "BankAccount/handlers/wire"
)

func Run(conf config.AppConfig, dbConn *gorm.DB) {

	allowHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
	allowOrigins := handlers.AllowedOrigins([]string{"*"})
	allowMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST",
		"PUT", "OPTIONS", "DELETE", "PATCH"})

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	accountHandler, err := wire.CreateBankAccountHandler(dbConn)
	if err != nil {
		return
	}
	r.HandleFunc("/accounts/{ID}", accountHandler.Serve)

	accountsHandler, err := wire.CreateBankAccountsHandler(dbConn)
	if err != nil {
		return
	}
	r.HandleFunc("/accounts", accountsHandler.Serve)

	fmt.Println("Starting server at", conf.Port)
	log.Fatal(http.ListenAndServe(conf.Port, handlers.CORS(allowOrigins,
		allowHeaders, allowMethods)(r)))

}
