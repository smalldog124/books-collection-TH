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
