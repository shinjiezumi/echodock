package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conn DBコネクション
var Conn *gorm.DB

// Initialize はDBコネクションを初期化します
func Initialize() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("DB_PROTOCOL")
	dbName := os.Getenv("DB_NAME")
	// TODO JawsDBではタイムゾーンの変更ができない
	//dsn := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	dsn := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=true"

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Conn = conn
}
