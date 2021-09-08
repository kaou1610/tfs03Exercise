package main

import (
	//"encoding/csv"	"log"
	//"os"
	"ex_craw/crawl"
	"log"

)



func main() {
	db, err := crawl.ConnectDb()
	if err != nil {
		log.Println(" Error: ", err)
		return 
	}
	defer db.Close()
	
	crawl.Crawler(db)
}
