package main

import (
	"crawler/db"
	"fmt"
	"log"
	"sync"
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

	var detail []db.Book
	var wg sync.WaitGroup
	c := colly.NewCollector()
	for round := 0; round <= 200; round += 10 {
		stratTime := time.Now()
		wg.Add(1)
		go visitWebsite(c, round, &wg, detail)
		wg.Wait()

		if err = bookMongo.SaveBooks(detail); err != nil {
			log.Println("data base err: ", err)
		}
		fmt.Printf("s: %v book: %d\n", time.Since(stratTime), round)
	}
	fmt.Println("Yes !!!")
}

func visitWebsite(c *colly.Collector, round int, wg *sync.WaitGroup, detail []db.Book) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		index := (round * 10) + i
		b := crawler(c, index)
		detail = append(detail, b)
	}
}

func crawler(c *colly.Collector, index int) db.Book {
	var detail []string
	url := fmt.Sprintf("http://e-service.nlt.go.th/ISBNReq/Detail/%d", index)
	c.OnHTML("input[value]", func(e *colly.HTMLElement) {
		detail = append(detail, e.Attr("value"))

	})
	if err := c.Visit(url); err != nil {
		log.Println("function crawler: ", err)
	}
	return formatSaveBook(detail)
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
