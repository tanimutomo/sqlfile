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

	s := sqlfile.New()

	if err := s.File("./examples/create_tables.sql"); err != nil {
		return
	}
	res, err := s.Exec(db)
	fmt.Println("Tables are created.")

	s = sqlfile.New()
	if err = s.File("./examples/insert.sql"); err != nil {
		return
	}
	fmt.Println("Load / s: ", s)
	fmt.Println("Load / err: ", err)

	res, err = s.Exec(db)
	fmt.Println("Exec / err: ", err)
	for i, r := range res {
		fmt.Printf("Query: %d\n", i)

		id, err := r.LastInsertId()
		if err != nil {
			fmt.Printf("when calling LastInseredId(): %s", err)
			return
		}
		fmt.Println("Exec / res.LastIntertedId: ", id)

		num, err := r.RowsAffected()
		if err != nil {
			fmt.Printf("when calling RowsAffected(): %s", err)
			return
		}
		fmt.Println("Exec / res.RowsAffected: ", num)
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
