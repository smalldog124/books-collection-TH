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
		books := []db.Book{
			{
				Name:      "ศึกยุทธหัตถี (ปกแข็ง)",
				Writer:    "สุภฤกษ์ บุญกอง",
				Publisher: "บริษัท สกายบุ๊กส์ จำกัด",
				Updatated: mock_date,
			},
			{
				Name:      "ยิหวายาใจ",
				Writer:    "แอลลี่",
				Publisher: "บริษัท มันดี กรุ๊ป จำกัด",
				Updatated: mock_date,
			},
		}

		err := bookMongo.SaveBooks(books)

		assert.Equal(t, nil, err)
	})
}
