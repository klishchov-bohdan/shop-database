package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	MysqlUserName        string
	MysqlPassword        string
	MysqlDBName          string
	MysqlMaxIdleCons     int
	MysqlMaxOpenCons     int
	MysqlConsMaxLifeTime int
}

func NewConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load("config/mysql.env")
	if err != nil {
		log.Fatal("Can`t load mysql.env")
	}
	cfg.MysqlUserName = os.Getenv("user")
	cfg.MysqlPassword = os.Getenv("pwd")
	cfg.MysqlDBName = os.Getenv("dbName")
	cfg.MysqlMaxIdleCons, _ = strconv.Atoi(os.Getenv("maxIdleConns"))
	cfg.MysqlMaxOpenCons, _ = strconv.Atoi(os.Getenv("maxOpenConns"))
	cfg.MysqlConsMaxLifeTime, _ = strconv.Atoi(os.Getenv("connMaxLifetime"))

	return cfg
}
