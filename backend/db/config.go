package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
)

var DB *sql.DB

const (
	dbUsername = "postgres"
	dbPassword = "200103287sdu"
	dbName     = "booking"
	dbHost     = "127.0.0.1"
	dbPort     = 5432
)

func init() {
	initDB()
	configDB()
}

func initDB() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName,
	)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("We have problems with connection database", err)
		os.Exit(1)
	}
}

func configDB() {
	tables()
}

func tables() {
	st, ioErr := ioutil.ReadFile("db/tables.sql")
	if ioErr != nil {
		fmt.Println("Cannot read db/tables.sql")
		os.Exit(1)
	}
	if _, err := DB.Exec(string(st)); err != nil {
		fmt.Println(err)
	}
}
