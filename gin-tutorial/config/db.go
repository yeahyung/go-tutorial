package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(DsnFormat,
		DbUsername,
		DbPassword,
		DbHost,
		DbPort,
		DbName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
