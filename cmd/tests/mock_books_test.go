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

func (b *mockBooksDB) GetBookCollectionBy(userID string) (books.BooksCollection, error) {
	argument := b.Called(userID)
	return argument.Get(0).(books.BooksCollection), argument.Error(1)
}

func (b *mockBooksDB) AddBookShelf(bookShelf books.Shelf, bookReview books.BookReview) error {
	argument := b.Called(mock.Anything, mock.Anything)
	return argument.Error(0)
}

func (b *mockBooksDB) AddBookWishList(bookWishList books.BookWishList) error {
	argument := b.Called(mock.Anything)
	return argument.Error(0)
}
