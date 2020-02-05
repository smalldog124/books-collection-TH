package handlers

import (
	"crawler/internal/books"
	"log"
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

type AddBookShelfRequest struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
	Score  int `json:"score"`
}

type AddBookWishListRequest struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}

func (b BooksAPI) BookScanHandler(context *gin.Context) {
	isbn := context.Param("isbn")
	book, err := b.Books.GetBookBy(isbn)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
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
	context.JSON(http.StatusOK, bookResponse)
}

func (b BooksAPI) AddBookShelfHandler(context *gin.Context) {
	var request AddBookShelfRequest
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("bad request %s", err.Error())
		return
	}

	bookShelf := books.BookShelf{
		UserID: request.UserID,
		BookID: request.BookID,
		Score:  request.Score,
	}
	bookReview := books.BookReview{
		BookID: request.BookID,
		Score:  request.Score,
	}
	if err := b.Books.AddBookShelf(bookShelf, bookReview); err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		log.Printf("internal AddBookShelf %s", err.Error())
		return
	}
	context.Status(http.StatusCreated)
}

func (b BooksAPI) AddBookWishListHandler(context *gin.Context) {
	var request AddBookWishListRequest
	if err := context.BindJSON(&request); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		log.Printf("AddBookWishListHandler bad request %s", err.Error())
		return
	}
	bookWishList := books.BookWishList{
		UserID: request.BookID,
		BookID: request.UserID,
	}
	if err := b.Books.AddBookWishList(bookWishList); err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		log.Printf("internal AddBookShelf %s", err.Error())
		return
	}
	context.Status(http.StatusCreated)
}
