package configuration

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type DbConfig struct {
	DbName	string
}


func (dbConf DbConfig) Connect() (*gorm.DB, error){
	fName := "configuration.DbConfig.Connect"

	dsn := fmt.Sprintf("host=localhost dbname=%s port=5432 sslmode=disable",
		dbConf.DbName)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, fName)
	}

	return gormDB, nil
}
