package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = ""
	DB_NAME     = "test"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_TYPE     = "mysql"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	if DB_TYPE == "mysql" {
		Db = connectDBSql()
	} else {
		Db = connectDBPostgree()
	}
	return Db
}

func connectDBSql() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error", err)
		return nil
	}

	return db
}

func connectDBPostgree() *gorm.DB {
	var err error
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DB_HOST, DB_PORT, DB_USERNAME, DB_NAME, DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error", err)
		return nil
	}

	return db
}
