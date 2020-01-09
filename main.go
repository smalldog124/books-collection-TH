package main

import (
	"crawler/db"
	"fmt"
	"log"

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
	var detail []string
	c := colly.NewCollector()

	c.OnHTML("input[value]", func(e *colly.HTMLElement) {
		detail = append(detail, e.Attr("value"))

	})
	c.Visit("http://e-service.nlt.go.th/ISBNReq/Detail/205789")
	var book db.Book
	book.Name = detail[0]
	book.Writer = detail[1]
	book.Translator = detail[2]
	book.EditionNote = detail[4]
	book.PrintYear = detail[6]
	book.NumberOfPage = detail[7]
	book.ISBNumber = detail[8]
	book.Publisher = detail[9]

	if err = bookMongo.SaveBooks(book); err != nil {
		log.Println("data base err: ", err)
	}

	fmt.Println("Yes !!!")
}
