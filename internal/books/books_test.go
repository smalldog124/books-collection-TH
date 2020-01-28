package books_test

import (
	"crawler/internal/books"
	"crawler/internal/database"
	"log"
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func TestPostgres(t *testing.T) {
	cfg := database.Config{
		User:       "smalldoc",
		Password:   "example",
		Host:       "localhost",
		Port:       5432,
		Name:       "books_collection_th",
		DisableTLS: true,
	}
	db, err := database.Open(cfg)
	if err != nil {
		log.Fatal("connect db error: ", err)
	}
	defer db.Close()
	postgresDB := books.BookDB{
		Connection: db,
	}
	t.Run("GetAllBooks_Should_Be_Internet_For_Bussinese", func(t *testing.T) {
		mockTime := time.Date(2020, 01, 28, 9, 12, 00, 00, time.UTC)
		expected := []books.Books{
			{
				ISBN:       "978-616-553-903-6",
				Name:       "อินเทอร์เน็ตเพื่องานธุรกิจ",
				Author:     "สุนทรีย์ โพธิ์อิ่ม, ไมตรี ฉลาดธรรม",
				Translator: "",
				Publisher:  "สำนักพิมพ์ศูนย์ส่งเสริมอาชีวะ",
				PrintYear:  "2562",
				Updated:    mockTime,
			},
		}

		actual, err := postgresDB.GetAllBooks()

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})
}