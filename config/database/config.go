package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type DBConfig interface {
	Dsn() string
	DbName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbName string
	dsn    string
	dbPort int
}

func NewConfig() *config {
	var cfg config
	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPass = os.Getenv("DB_PASS")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbName = os.Getenv("DB_NAME")
	var err error
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln("configuration error: ", err)
	}
	cfg.dsn = fmt.Sprintf("mongodb://%s:%d/%s", cfg.dbHost, cfg.dbPort, cfg.dbName)

	return &cfg
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DbName() string {
	return c.dbName
}
