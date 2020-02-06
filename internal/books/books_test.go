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
				ID:         1,
				ISBN:       "978-616-18-2996-4",
				Name:       "ทำไม Netflix ถึงมีแต่คนโตครเก่ง",
				Writer:     "แพตตี้ แมคคอร์ด",
				Translator: "วิกันดา จันทร์ทองสุข",
				Publisher:  "บริษัทอมรินทร์พริ้นติ้งแอนด์พับลิชซิ่ง จำกัด (มหาชน)",
				PrintYear:  "2558",
				Updated:    mockTime,
			},
			{
				ID:         2,
				ISBN:       "978-616-553-903-6",
				Name:       "อินเทอร์เน็ตเพื่องานธุรกิจ",
				Writer:     "สุนทรีย์ โพธิ์อิ่ม, ไมตรี ฉลาดธรรม",
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
	t.Run("GetBookBy_Input_ISBN_978-61618-2996-4_Should_Be_Netflix", func(t *testing.T) {
		mockTime := time.Date(2020, 01, 28, 9, 12, 00, 00, time.UTC)
		expected := books.Books{
			ID:         1,
			ISBN:       "978-616-18-2996-4",
			Name:       "ทำไม Netflix ถึงมีแต่คนโตครเก่ง",
			Writer:     "แพตตี้ แมคคอร์ด",
			Translator: "วิกันดา จันทร์ทองสุข",
			Publisher:  "บริษัทอมรินทร์พริ้นติ้งแอนด์พับลิชซิ่ง จำกัด (มหาชน)",
			PrintYear:  "2558",
			Updated:    mockTime,
		}

		actual, err := postgresDB.GetBookBy("978-616-18-2996-4")

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("AddBookShelf_Input_Book_Shelf_And_Book_Review_Should_Be_Ceated", func(t *testing.T) {
		book_shelf := books.BookShelf{
			UserID: 137499732,
			BookID: 1,
			Score:  4,
		}
		book_review := books.BookReview{
			BookID: 1,
			Score:  4,
		}

		actual := postgresDB.AddBookShelf(book_shelf, book_review)

		assert.Equal(t, nil, actual)
	})
	t.Run("AddBookWishList_Input_Book_ID_1_And_User_ID_137499732_Should_Be_Ceated", func(t *testing.T) {
		book_wishlist := books.BookWishList{
			UserID: 137499732,
			BookID: 1,
		}

		actual := postgresDB.AddBookWishList(book_wishlist)

		assert.Equal(t, nil, actual)
	})
}
