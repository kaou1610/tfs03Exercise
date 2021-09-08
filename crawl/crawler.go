package crawl

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

type Movie struct {
	//key string
	name string
	year string
	rate string

}

func Crawler(db *gorm.DB) {
	// fName := "movie.csv"
	// file, err := os.Create(fName)
	// if err != nil {
	// 	log.Fatal("Could not create file")
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %s", r.URL)
	})
	c.OnError(func(r *colly.Response, e error) {
		log.Println("Error: ", e)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visited: %s", r.Request.URL)
	})

	c.OnHTML("tr", func(h *colly.HTMLElement) {
		movie := Movie{}
		movie.name = h.ChildText(".titleColumn > a")
		movie.year = h.ChildText(".titleColumn .secondaryInfo")
		movie.rate = h.ChildText(".ratingColumn > strong")
		db.Exec("INSERT movie SET name=?, year=?, rate=?", movie.name, movie.year, movie.rate) //error in here
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

// var db *gorm.DB
// var err error

// func dbConnection() (*gorm.DB, error)  {
// 	db, err = gorm.Open("mysql", "mphuong:16101999@tcp(127.0.0.1:3306)/movies?charset=utf8&parseTime=True")
//     if err != nil {
//         log.Printf("Error %s when opening DBn", err)
//         return nil, err
//     }

//     return db, nil
// }