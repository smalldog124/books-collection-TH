package main

import (
	"crawler/db"
	"crawler/find"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo"
)

func main() {
	resp, err := http.Get("http://www.e-service.nlt.go.th/ISBNReq/ListSearchPub?keywordTypeKey=1")
	if err != nil {
		log.Println("call fail", err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	table := find.GetTableData(body)
	books := find.BooksDetail(table)
	// err = ioutil.WriteFile("books.html", body, 0644)
	// if err != nil {
	// 	fmt.Println("writeFile fail")
	// }

	fmt.Printf("books !!!\n %v", books)
	DBConnection, err := mgo.Dial(db.MongoEndPoint)
	if err != nil {
		log.Fatalf("Cannot connect database %s\n", err.Error())
	}
	bookMongo := db.BookMongo{
		DBConnection: DBConnection,
		Database:     db.Database,
	}
	var booksDetail []db.Book
	for _, book := range books {
		b := db.Book{
			Name:      book[0],
			Writer:    book[1],
			Publisher: book[2],
			Updatated: time.Now(),
		}
		booksDetail = append(booksDetail, b)
	}
	log.Println(bookMongo.SaveBooks(booksDetail))
	fmt.Println("Yes !!!")
}
