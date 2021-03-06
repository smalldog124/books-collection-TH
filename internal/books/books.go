package books

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type BookDB struct {
	Connection *sqlx.DB
}

type BooksInserface interface {
	GetBookBy(isbn string) (Books, error)
	AddBookShelf(bookShelf Shelf, bookReview BookReview) error
	AddBookWishList(bookWishList BookWishList) error
	GetBookCollectionBy(userID string) (BooksCollection, error)
}

type Books struct {
	ID         int       `db:"id"`
	ISBN       string    `db:"isbn"`
	Name       string    `db:"name"`
	Writer     string    `db:"writer"`
	Translator string    `db:"translator"`
	Publisher  string    `db:"publisher"`
	PrintYear  string    `db:"print_year"`
	Updated    time.Time `db:"date_updated"`
}

type BooksCollection struct {
	BooksSelf     []BooksShelf
	BooksWishList []Books
}
type BooksShelf struct {
	ID         int       `db:"id"`
	ISBN       string    `db:"isbn"`
	Name       string    `db:"name"`
	Writer     string    `db:"writer"`
	Translator string    `db:"translator"`
	Publisher  string    `db:"publisher"`
	PrintYear  string    `db:"print_year"`
	Updated    time.Time `db:"date_updated"`
	Score      int       `db:"score"`
}
type Shelf struct {
	UserID int `db:"user_id"`
	BookID int `db:"book_id"`
	Score  int `db:"score"`
}
type BookReview struct {
	BookID int `db:"book_id"`
	Score  int `db:"score"`
}
type BookWishList struct {
	UserID int `db:"user_id"`
	BookID int `db:"book_id"`
}

func (postgres BookDB) GetAllBooks() ([]Books, error) {
	var book []Books
	const query = `SELECT id,isbn,name,writer,translator,publisher,print_year,date_updated FROM books`
	err := postgres.Connection.Select(&book, query)
	if err != nil {
		return book, err
	}
	return book, err
}

func (postgres BookDB) GetBookBy(isbn string) (Books, error) {
	var book Books
	const query = `SELECT id,isbn,name,writer,translator,publisher,print_year,date_updated FROM books WHERE isbn=$1`
	err := postgres.Connection.Get(&book, query, isbn)
	if err != nil {
		return book, err
	}
	return book, err
}

func (postgres BookDB) CreateBook(book Books) error {
	const query = `INSERT INTO books (isbn,name,writer,translator,publisher,print_year)VALUES(:isbn,:name,:writer,:translator,:publisher,:print_year);`
	tx := postgres.Connection.MustBegin()
	_, err := tx.NamedExec(query, &book)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (postgres BookDB) AddBookShelf(bookShelf Shelf, bookReview BookReview) error {
	const queryShelf = `INSERT INTO shelf (user_id,book_id,score) VALUES (:user_id,:book_id,:score)`
	const queryReview = `INSERT INTO review (book_id,score) VALUES (:book_id,:score)`
	tx := postgres.Connection.MustBegin()
	_, err := tx.NamedExec(queryShelf, &bookShelf)
	if err != nil {
		return err
	}
	_, err = tx.NamedExec(queryReview, &bookReview)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (postgres BookDB) AddBookWishList(bookWishList BookWishList) error {
	const queryWishList = `INSERT INTO wish_list (user_id,book_id) VALUES (:user_id,:book_id)`
	tx := postgres.Connection.MustBegin()
	_, err := tx.NamedExec(queryWishList, &bookWishList)
	if err != nil {
		return err
	}
	return nil
}

func (postgres BookDB) GetBookCollectionBy(userID string) (BooksCollection, error) {
	var booksCollection BooksCollection
	const queryShelf = `SELECT books.id,isbn,name,writer,translator,publisher,print_year,date_updated,score 
	FROM books
	LEFT JOIN shelf
	ON books.id = shelf.book_id
	WHERE shelf.user_id=$1`
	var shelf []BooksShelf
	if err := postgres.Connection.Select(&shelf, queryShelf, userID); err != nil {
		return booksCollection, err
	}
	const queryWishList = `SELECT books.id,isbn,name,writer,translator,publisher,print_year,date_updated 
	FROM books
	LEFT JOIN wish_list
	ON books.id = wish_list.book_id
	WHERE wish_list.user_id=$1`
	var wishList []Books
	if err := postgres.Connection.Select(&wishList, queryWishList, userID); err != nil {
		return booksCollection, err
	}
	return BooksCollection{BooksSelf: shelf, BooksWishList: wishList}, nil
}
