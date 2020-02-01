package test

import (
	"crawler/cmd/handlers"
	"crawler/internal/books"
	"io/ioutil"
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
