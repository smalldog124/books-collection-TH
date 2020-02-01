package main

import (
	"crawler/cmd/handlers"
	"crawler/internal/books"
	"crawler/internal/database"
	"flag"
	"fmt"
	"log"
	"net/http"

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
	booksAPI := handlers.BooksAPI{
		Books: &bookPostgres,
	}
	engine := gin.Default()
	engine.GET("/books", func(c *gin.Context) {
		books, err := bookPostgres.GetAllBooks()
		if err != nil {
			log.Println(err)
			return
		}
		bookDetail := []handlers.BookResponse{
			{
				ISBN:       books[0].ISBN,
				Name:       books[0].Name,
				Writer:     books[0].Writer,
				Translator: books[0].Translator,
				Publisher:  books[0].Publisher,
				PrintYear:  books[0].PrintYear,
				Updatated:  books[0].Updated.Format("2006-01-02"),
			},
		}
		c.JSON(http.StatusOK, bookDetail)
	})
	engine.GET("/api/v1/scan/:isbn", booksAPI.BookScanHandler)
	log.Fatal(engine.Run(fmt.Sprintf(":%s", port)))
}
