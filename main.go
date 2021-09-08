package main

import (
	//"encoding/csv"	"log"
	//"os"
	"ex_craw/crawl"
	"log"

	_ "github.com/go-sql-driver/mysql"
)



func main() {
	db, err := crawl.ConnectDb()
	if err != nil {
		log.Println(" Error: ", err)
		return 
	}

	crawl.Crawler(db)
}
