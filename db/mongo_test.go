package db_test

import (
	"crawler/db"
	"testing"
	"time"

	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/assert"
)

func Test_Books_Repository(t *testing.T) {
	DBConnection, err := mgo.Dial(db.MongoEndPoint)
	if err != nil {
		t.Fatalf("Cannot connect database %s\n", err.Error())
	}
	bookMongo := db.BookMongo{
		DBConnection: DBConnection,
		Database:     db.Database,
	}

	t.Run("SaveBooks_Input_Books_Should_Be_Success", func(t *testing.T) {
		mock_date, _ := time.Parse(time.RFC3339, "2020-01-02T15:04:05Z")
		books := db.Book{
			ISBNumber:    "978-616-411-106-6",
			Name:         "ศึกยุทธหัตถี (ปกแข็ง)",
			Writer:       "สุภฤกษ์ บุญกอง",
			Translator:   "",
			Publisher:    "บริษัท สกายบุ๊กส์ จำกัด",
			EditionNote:  "1",
			PrintYear:    "2563",
			NumberOfPage: "192",
			Updatated:    mock_date,
		}

		err := bookMongo.SaveBooks(books)

		assert.Equal(t, nil, err)
	})
}
