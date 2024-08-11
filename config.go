package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Config struct {
	DB *gorm.DB
}

var ConfigInstance Config

func InitDB(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	ConfigInstance.DB = db
	return nil
}
