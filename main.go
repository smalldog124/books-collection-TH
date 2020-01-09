package main

import (
	"crawler/db"
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/gocolly/colly"
)

func main() {
	DBConnection, err := mgo.Dial(db.MongoEndPoint)
	if err != nil {
		log.Fatalf("Cannot connect database %s\n", err.Error())
	}
	bookMongo := db.BookMongo{
		DBConnection: DBConnection,
		Database:     db.Database,
	}

	c := colly.NewCollector()
	for index := 1; index <= 200; index++ {
		stratTime := time.Now()
		url := fmt.Sprintf("http://e-service.nlt.go.th/ISBNReq/Detail/%d", index)
		detail, err := crawler(c, url)
		if err != nil {
			log.Println("function crawler: ", err)
		}
		book := formatSaveBook(detail)
		if err = bookMongo.SaveBooks(book); err != nil {
			log.Println("data base err: ", err)
		}
		fmt.Printf("s: %v book: %d\n", time.Since(stratTime), index)
	}
	fmt.Println("Yes !!!")
}

func crawler(c *colly.Collector, url string) ([]string, error) {
	var detail []string
	c.OnHTML("input[value]", func(e *colly.HTMLElement) {
		detail = append(detail, e.Attr("value"))

	})
	return detail, c.Visit(url)
}

func formatSaveBook(detail []string) db.Book {
	var book db.Book
	book.Name = detail[0]
	book.Writer = detail[1]
	book.Translator = detail[2]
	book.EditionNote = detail[4]
	book.PrintYear = detail[6]
	book.NumberOfPage = detail[7]
	book.ISBNumber = detail[8]
	book.Publisher = detail[9]
	book.Updatated = time.Now()
	return book
}
