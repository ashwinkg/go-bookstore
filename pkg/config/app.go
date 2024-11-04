package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Connect() {
	// Connection string format: <username>:<password>@tcp(<hostname>:<port>)/<database>?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:rootpassword@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
