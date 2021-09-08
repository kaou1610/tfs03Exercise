package crawl

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "mphuong:16101999@tcp(127.0.0.1:3306)/movies")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Successful connection to database")

	return db, nil
}

func InsertDB(db *sql.DB, name string, year string, rating string)  {
	
}