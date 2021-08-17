package crawl

import (
	"database/sql"
	"fmt"
	"time"
)

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Successful connection to database")
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}