package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	HOST     string
	USER     string
	PASSWORD string
	DB       string
	PORT     string
}

func newConfig() *config {
	return &config{
		HOST:     os.Getenv("MYSQL_HOST"),
		USER:     os.Getenv("MYSQL_USER"),
		PASSWORD: os.Getenv("MYSQL_PASSWORD"),
		DB:       os.Getenv("MYSQL_DB"),
		PORT:     os.Getenv("MYSQL_PORT"),
	}
}

var DB = connection()

func connection() *sql.DB {
	config := newConfig()

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.USER, config.PASSWORD, config.HOST, config.PORT, config.DB)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatalf("Fail in Open database: %v \n DNS: %s", err, dns)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Fail in PING Test to Database: %v \n DNS: %s", err, dns)
	}

	fmt.Println("Good connection with the DB \n DNS: ", dns)
	return db
}
