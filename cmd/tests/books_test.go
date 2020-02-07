package test

import (
	"bytes"
	"crawler/cmd/handlers"
	"crawler/internal/books"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_BookScanHandler_Input_ISBN_978_616_18_2996_4_Shold_Be_Book_ID_1(t *testing.T) {
	expected := `{"id":1,"isbn":"978-616-18-2996-4","name":"ทำไม Netflix ถึงมีแต่คนโตครเก่ง","writer":"แพตตี้ แมคคอร์ด","translator":"วิกันดา จันทร์ทองสุข","publisher":"บริษัทอมรินทร์พริ้นติ้งแอนด์พับลิชซิ่ง จำกัด (มหาชน)","edition_note":"","print_year":"2558","no_of_page":"","updated":"2020-01-28"}
`
	isbn := "978-616-18-2996-4"
	request := httptest.NewRequest("GET", "/api/v1/scan/978-616-18-2996-4", nil)
	write := httptest.NewRecorder()
	mockTime := time.Date(2020, 01, 28, 9, 12, 00, 00, time.UTC)
	mockBooksDB := new(mockBooksDB)
	mockBooksDB.On("GetBookBy", isbn).Return(books.Books{
		ID:         1,
		ISBN:       "978-616-18-2996-4",
		Name:       "ทำไม Netflix ถึงมีแต่คนโตครเก่ง",
		Writer:     "แพตตี้ แมคคอร์ด",
		Translator: "วิกันดา จันทร์ทองสุข",
		Publisher:  "บริษัทอมรินทร์พริ้นติ้งแอนด์พับลิชซิ่ง จำกัด (มหาชน)",
		PrintYear:  "2558",
		Updated:    mockTime,
	}, nil)
	booksAPI := handlers.BooksAPI{
		Books: mockBooksDB,
	}

	mockRoute := gin.Default()
	mockRoute.GET("/api/v1/scan/:isbn", booksAPI.BookScanHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}

func Test_CollectionHandler_Input_User_ID_137499732_Shold_Be_Book_Shelf_1_And_Whislist_1(t *testing.T) {
	expected := `{"shelf":[{"id":1,"isbn":"978-616-18-2996-4","name":"ทำไม Netflix ถึงมีแต่คนโตครเก่ง","writer":"แพตตี้ แมคคอร์ด","translator":"วิกันดา จันทร์ทองสุข","publisher":"บริษัทอมรินทร์พริ้นติ้งแอนด์พับลิชซิ่ง จำกัด (มหาชน)","edition_note":"","print_year":"2558","score":4}],"wishlist":[{"id":2,"isbn":"978-616-553-903-6","name":"อินเทอร์เน็ตเพื่องานธุรกิจ","writer":"สุนทรีย์ โพธิ์อิ่ม, ไมตรี ฉลาดธรรม","translator":"","publisher":"สำนักพิมพ์ศูนย์ส่งเสริมอาชีวะ","edition_note":"","print_year":"2562","no_of_page":"","updated":"2020-01-28"}]}
`
	UserID := "137499732"
	request := httptest.NewRequest("GET", "/api/v1/collection/137499732", nil)
	write := httptest.NewRecorder()
	mockTime := time.Date(2020, 01, 28, 9, 12, 00, 00, time.UTC)
	mockBooksDB := new(mockBooksDB)
	mockBooksDB.On("GetBookCollectionBy", UserID).Return(books.BooksCollection{
		BooksSelf: []books.BooksShelf{
			{
				ID:         1,
				ISBN:       "978-616-18-2996-4",
				Name:       "ทำไม Netflix ถึงมีแต่คนโตครเก่ง",
				Writer:     "แพตตี้ แมคคอร์ด",
				Translator: "วิกันดา จันทร์ทองสุข",
				Publisher:  "บริษัทอมรินทร์พริ้นติ้งแอนด์พับลิชซิ่ง จำกัด (มหาชน)",
				PrintYear:  "2558",
				Updated:    mockTime,
				Score:      4,
			},
		},
		BooksWishList: []books.Books{
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
		},
	}, nil)
	booksAPI := handlers.BooksAPI{
		Books: mockBooksDB,
	}

	mockRoute := gin.Default()
	mockRoute.GET("/api/v1/collection/:user_id", booksAPI.CollectionHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}
func Test_AddBookShelfHandler_Handler_Input_ISBN_978_616_18_2996_4_Shold_Htpp_Status_201(t *testing.T) {
	requestBody := handlers.AddBookShelfRequest{
		UserID: 129494830394,
		BookID: 1,
		Score:  4,
	}
	jsonRequest, _ := json.Marshal(requestBody)
	request := httptest.NewRequest("POST", "/api/v1/book/shelf", bytes.NewBuffer(jsonRequest))
	write := httptest.NewRecorder()
	bookShelf := books.Shelf{
		UserID: 129494830394,
		BookID: 1,
		Score:  4,
	}
	bookReview := books.BookReview{
		BookID: 1,
		Score:  4,
	}
	mockBooksDB := new(mockBooksDB)
	mockBooksDB.On("AddBookShelf", bookShelf, bookReview).Return(nil)
	booksAPI := handlers.BooksAPI{
		Books: mockBooksDB,
	}

	mockRoute := gin.Default()
	mockRoute.POST("/api/v1/book/shelf", booksAPI.AddBookShelfHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()

	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func Test_AddBookWishListHandler_Handler_Input_Book_ID_1_And_User_ID_129494830394_Shold_Htpp_Status_201(t *testing.T) {
	requestBody := handlers.AddBookWishListRequest{
		UserID: 129494830394,
		BookID: 1,
	}
	jsonRequest, _ := json.Marshal(requestBody)
	request := httptest.NewRequest("POST", "/api/v1/book/wishlist", bytes.NewBuffer(jsonRequest))
	write := httptest.NewRecorder()
	bookWishList := books.BookWishList{
		UserID: 129494830394,
		BookID: 1,
	}
	mockBooksDB := new(mockBooksDB)
	mockBooksDB.On("AddBookWishList", bookWishList).Return(nil)
	booksAPI := handlers.BooksAPI{
		Books: mockBooksDB,
	}

	mockRoute := gin.Default()
	mockRoute.POST("/api/v1/book/wishlist", booksAPI.AddBookWishListHandler)
	mockRoute.ServeHTTP(write, request)
	response := write.Result()

	assert.Equal(t, http.StatusCreated, response.StatusCode)
}
