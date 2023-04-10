package config

import (
	"os"
	"strconv"
)

var (
	APP_HOST string = "127.0.0.1"
	APP_PORT uint   = 3000
	DB_PORT  uint   = 3306
	DB_HOST  string = "db"
	DB_USER  string = "root"
	DB_PASS  string = "password"
	DB_NAME  string = "judge"
)

func init() {
	if v := os.Getenv("APP_HOST"); v != "" {
		APP_HOST = v
	}
	if v, err := strconv.ParseUint(os.Getenv("APP_PORT"), 10, 32); err != nil && v != 0 {
		APP_PORT = uint(v)
	}
	if v, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32); err != nil && v != 0 {
		DB_PORT = uint(v)
	}
	if v := os.Getenv("DB_HOST"); v != "" {
		DB_HOST = v
	}
	if v := os.Getenv("DB_USER"); v != "" {
		DB_USER = v
	}
	if v := os.Getenv("DB_PASS"); v != "" {
		DB_PASS = v
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		DB_NAME = v
	}
}
