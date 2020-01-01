package main

import (
	"crawler/db"
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func main() {
	r := gin.Default()
	var dbEnpoint string
	flag.StringVar(&dbEnpoint, "dbEnpoint", "books-db", "database host name")
	flag.Parse()
	DBConnection, err := mgo.Dial(dbEnpoint)
	if err != nil {
		log.Fatalf("Cannot connect database %s\n", err.Error())
	}
	bookMongo := db.BookMongo{
		DBConnection: DBConnection,
		Database:     "books",
	}
	r.GET("/books", func(c *gin.Context) {
		books, err := bookMongo.GetBooks()
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, books)
	})
	log.Fatal(r.Run(":3030"))
}
