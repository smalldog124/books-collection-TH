package books

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type BookDB struct {
	Connection *sqlx.DB
}

type Books struct {
	ISBN       string    `db:"id"`
	Name       string    `db:"name"`
	Author     string    `db:"author"`
	Translator string    `db:"translator"`
	Publisher  string    `db:"publisher"`
	PrintYear  string    `db:"print_year"`
	Updated    time.Time `db:"date_updated"`
}

func (postgres BookDB) GetAllBooks() ([]Books, error) {
	var book []Books
	const query = `SELECT id,name,author,translator,publisher,print_year,date_updated FROM books`
	err := postgres.Connection.Select(&book, query)
	if err != nil {
		return book, err
	}
	return book, err
}

func (postgres BookDB) CreateBook(book Books) error {
	const query = `INSERT INTO books (id,name,author,translator,publisher,print_year)VALUES($1,$2,$3,$4,$5,$6);`
	tx := postgres.Connection.MustBegin()
	tx.MustExec(query, book.ISBN, book.Name, book.Author, book.Translator, book.Publisher, book.PrintYear)
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
