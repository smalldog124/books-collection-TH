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
	Name      string    `bson:"naem" json:"name"`
	Writer    string    `bson:"writer" json:"writer"`
	Publisher string    `bsom:"publisher" json:"publisher"`
	Updatated time.Time `bson:"updated" json:"updated"`
}

func (db BookMongo) SaveBooks(books []Book) error {
	bulk := db.DBConnection.DB(db.Database).C("detail").Bulk()
	var err error
	for _, book := range books {
		bulk.Insert(book)
		err = db.DBConnection.DB(db.Database).C("detail").Insert(&book)
	}
	return err
}

func (db BookMongo) GetBooks() ([]Book, error) {
	var books []Book
	err := db.DBConnection.DB(db.Database).C("detail").Find(nil).All(&books)
	return books, err
}
