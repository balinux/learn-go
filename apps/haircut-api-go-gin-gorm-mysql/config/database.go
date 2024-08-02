package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "balinuxsv1:123B@l!nuxV1@tcp(127.0.0.1:3306)/haircut?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:1234Usd!@tcp(127.0.0.1:3306)/haircut?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = database
}
