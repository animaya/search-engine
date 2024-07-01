package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() {
	dburl := os.Getenv("DATABASE_URL")

	var err error

	DBConn, err = gorm.Open(postgres.Open(dburl))

	if err != nil {
		fmt.Println("failed to connect to datase")
		panic("failed to connect to datase")
	}

	err = DBConn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error

	if err != nil {
		fmt.Println("cant install uuid extension")
		panic(err)
	}

	err = DBConn.AutoMigrate()
	if err != nil {

		panic(err)
	}
}

func GetDB() *gorm.DB {
	return DBConn
}
