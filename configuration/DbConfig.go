package configuration

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	ent "BankAccount/entities"
)

type DbConfig struct {
	DbName	string
	Port int
}


func (dbConf DbConfig) Connect() (*gorm.DB, error){
	fName := "configuration.DbConfig.Connect"

	dsn := fmt.Sprintf("host=localhost dbname=%s port=%d sslmode=disable",
		dbConf.DbName, dbConf.Port)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, fName)
	}

	err = dbConn.AutoMigrate(
		&ent.BankAccount{},
		)
	if err != nil {
		fmt.Println("AutoMigrate error")
		return nil, errors.Wrap(err, fName)
	}

	return dbConn, nil
}
