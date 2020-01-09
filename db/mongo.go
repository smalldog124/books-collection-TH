package db

import (
	"github.com/globalsign/mgo"
	"time"
)

const (
	MongoEndPoint = "mongodb://localhost:27017"
	Database      = "books"
)

type BookMongo struct {
	DBConnection *mgo.Session
	Database     string
}

type Book struct {
	ISBNumber    string    `bson:"ISBNumber" json:"ISBNumber"`
	Name         string    `bson:"naem" json:"name"`
	Writer       string    `bson:"writer" json:"writer"`
	Translator   string    `bson:"translator" json:"translator"`
	Publisher    string    `bson:"publisher" json:"publisher"`
	EditionNote  string    `bson:"edition_note" json:"edition_note"`
	PrintYear    string    `bson:"print_year" json:"print_year"`
	NumberOfPage string    `bson:"no_of_page" json:"no_of_page"`
	Updatated    time.Time `bson:"updated" json:"updated"`
}

func (db BookMongo) SaveBooks(books Book) error {
	return db.DBConnection.DB(db.Database).C("detail").Insert(books)
}

func (db BookMongo) GetBooks() ([]Book, error) {
	var books []Book
	err := db.DBConnection.DB(db.Database).C("detail").Find(nil).All(&books)
	return books, err
}
