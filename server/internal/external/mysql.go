package external

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"yudegaki.github.com/rewrite-judge/internal/config"
)

var DB *gorm.DB

func InitDB() {
	port := config.DB_PORT
	host := config.DB_HOST
	user := config.DB_USER
	password := config.DB_PASS
	dbName := config.DB_NAME

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	Initialize()
}
