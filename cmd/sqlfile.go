package main

import (
	"database/sql"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/tanimutomo/sqlfile"
)

func main() {
	db := newDB()
	sqlfile.Exec(db, "./example.sql")
}

func newDB() *sql.DB {
	godotenv.Load()

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
