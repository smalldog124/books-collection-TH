package main

import (
	"crawler/internal/books"
	"crawler/internal/database"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var port, dbhost, dbschema, dbusername, dbpassword, disableTLS string
	var dbport int
	flag.StringVar(&port, "port", "3030", "port for open service")
	flag.StringVar(&dbhost, "dbhost", "localhost", "database host name")
	flag.IntVar(&dbport, "dbport", 5432, "database port")
	flag.StringVar(&dbschema, "dbschema", "books_collection_th", "database schema name")
	flag.StringVar(&dbusername, "dbusername", "smalldoc", "database user name")
	flag.StringVar(&dbpassword, "dbpassword", "example", "database password")
	flag.StringVar(&disableTLS, "disableTLS", "Y", "database disableTLS[Y/n]")
	flag.Parse()
	var databaseTSL bool
	if disableTLS == "n" {
		databaseTSL = false
	} else {
		databaseTSL = true
	}
	dbConfig := database.Config{
		User:       dbusername,
		Password:   dbpassword,
		Host:       dbhost,
		Port:       dbport,
		Name:       dbschema,
		DisableTLS: databaseTSL,
	}

	db, err := database.Open(dbConfig)
	if err != nil {
		log.Fatal("connecting database fail", err)
	}
	bookPostgres := books.BookDB{
		Connection: db,
	}
	engine := gin.Default()
	type BookDetail struct {
		ISBNumber    string    `json:"ISBNumber"`
		Name         string    `json:"name"`
		Writer       string    `json:"writer"`
		Translator   string    `json:"translator"`
		Publisher    string    `json:"publisher"`
		EditionNote  string    `json:"edition_note"`
		PrintYear    string    `json:"print_year"`
		NumberOfPage string    `json:"no_of_page"`
		Updatated    time.Time `json:"updated"`
	}
	engine.GET("/books", func(c *gin.Context) {
		books, err := bookPostgres.GetAllBooks()
		if err != nil {
			log.Println(err)
			return
		}
		bookDetail := BookDetail{
			ISBNumber:  books[0].ISBN,
			Name:       books[0].Name,
			Writer:     books[0].Author,
			Translator: books[0].Translator,
			Publisher:  books[0].Publisher,
			PrintYear:  books[0].PrintYear,
			Updatated:  books[0].Updated,
		}
		c.JSON(http.StatusOK, bookDetail)
	})
	log.Fatal(engine.Run(fmt.Sprintf(":%s", port)))
}
