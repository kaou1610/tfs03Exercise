package crawl

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/gocolly/colly"
)

type Movie struct {
	//key string
	name string
	year string
	rate string

}

func Crawler(db *sql.DB) {
	// fName := "movie.csv"
	// file, err := os.Create(fName)
	// if err != nil {
	// 	log.Fatal("Could not create file")
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()
	c := colly.NewCollector(
	//colly.AllowedDomains("https://www.imdb.com/"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %s", r.URL)
		fmt.Println()
	})
	c.OnError(func(r *colly.Response, e error) {
		log.Println("Error: ", e)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visited: %s", r.Request.URL)
	})
	fmt.Println()

	c.OnHTML("tr", func(h *colly.HTMLElement) {
		movie := Movie{}
		movie.name = h.ChildText(".titleColumn > a")
		movie.year = h.ChildText(".titleColumn .secondaryInfo")
		movie.rate = h.ChildText(".ratingColumn > strong")
		
		stmt, err := db.Prepare("INSERT INTO movie (name,year,rate) values (?,?,?)")
		handleError(err)
		res, err1 := stmt.Exec(movie.name, movie.year, movie.rate)
		handleError(err1)
		_, err2 := res.LastInsertId()
		if err != nil {
			log.Fatal(err2)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Crawler done", r.Request.URL)
	})
	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}