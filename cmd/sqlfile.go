package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/tanimutomo/sqlfile"
)

func main() {
	db := newDB()
	_, err := sqlfile.Exec(db, "./example.sql")
	if err != nil {
		fmt.Println("Exec / err: ", err.Error())
	}
}

func newDB() *sql.DB {
	DBMS := os.Getenv("DB_TYPE")
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")
	DBHOST := os.Getenv("DB_HOST")
	DBPORT := os.Getenv("DB_PORT")
	CONNECT := USER + ":" + PASS + "@(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=true"

	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}

	return db
}
