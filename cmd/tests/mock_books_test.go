package test

import (
	"github.com/stretchr/testify/mock"

	"crawler/internal/books"
)

type mockBooksDB struct {
	mock.Mock
}

func (b *mockBooksDB) GetBookBy(isbn string) (books.Books, error) {
	argument := b.Called(isbn)
	return argument.Get(0).(books.Books), argument.Error(1)
}

func (b *mockBooksDB) AddBookShelf(bookShelf books.BookShelf, bookReview books.BookReview) error {
	argument := b.Called(mock.Anything, mock.Anything)
	return argument.Error(0)
}

func (b *mockBooksDB) AddBookWishList(bookWishList books.BookWishList) error {
	argument := b.Called(mock.Anything)
	return argument.Error(0)
}
