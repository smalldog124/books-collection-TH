package handlers

import (
	"crawler/internal/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BooksAPI struct {
	Books books.BooksInserface
}

type BookResponse struct {
	ID           int    `json:"id"`
	ISBN         string `json:"isbn"`
	Name         string `json:"name"`
	Writer       string `json:"writer"`
	Translator   string `json:"translator"`
	Publisher    string `json:"publisher"`
	EditionNote  string `json:"edition_note"`
	PrintYear    string `json:"print_year"`
	NumberOfPage string `json:"no_of_page"`
	Updatated    string `json:"updated"`
}

func (b BooksAPI) BookScanHandler(constext *gin.Context) {
	isbn := constext.Param("isbn")
	book, err := b.Books.GetBookBy(isbn)
	if err != nil {
		constext.String(http.StatusInternalServerError, err.Error())
		return
	}
	bookResponse := BookResponse{
		ID:         book.ID,
		ISBN:       book.ISBN,
		Name:       book.Name,
		Writer:     book.Writer,
		Translator: book.Translator,
		Publisher:  book.Publisher,
		PrintYear:  book.PrintYear,
		Updatated:  book.Updated.Format("2006-01-02"),
	}
	constext.JSON(http.StatusOK, bookResponse)
}
