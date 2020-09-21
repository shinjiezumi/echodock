package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Conn *gorm.DB

func Initialize() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("DB_PROTOCOL")
	dbName := os.Getenv("DB_NAME")
	dsn := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=true"

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Conn = conn
}