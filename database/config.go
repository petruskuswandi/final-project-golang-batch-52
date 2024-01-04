package database

import (
	"fmt"
	"log"

	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	if config.DB_CONNECTION == "mysql" {
		dsnMySQL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_DATABASE)
		DB, err = gorm.Open(mysql.Open(dsnMySQL), &gorm.Config{})
	}

	if config.DB_CONNECTION == "pgsql" {
		dsnPostgres := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.DB_HOST, config.DB_USERNAME, config.DB_PASSWORD, config.DB_DATABASE, config.DB_PORT)
		DB, err = gorm.Open(postgres.Open(dsnPostgres), &gorm.Config{})
	}

	if err != nil {
		panic("Can't connect to database.")
	}

	log.Println("Connected to database.")
}
