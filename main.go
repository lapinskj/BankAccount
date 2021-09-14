package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	conf "BankAccount/configuration"
)

func main () {
	fName := "main.main"

	configPath := "config.json"
	conf, confErr := conf.LoadConfig(configPath)
	if confErr != nil {
		log.Fatalln(errors.Wrap(confErr, fName))
	}

	dbConn, err := conf.DbConfig.Connect()
	if err != nil {
		log.Fatalln(errors.Wrap(err, fName))
	}
	fmt.Println("Connected to database successfully")

	Run(conf, dbConn)
}
