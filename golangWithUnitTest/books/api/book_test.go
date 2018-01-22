package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Name: "Book A", Author: "A", ISBN: "1"}
	json := book.ToJSON()

	assert.Equal(t, `{"name":"Book A","author":"A","isbn":"1"}`, string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"name":"Book A","author":"A","isbn":"1"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Name: "Book A", Author: "A", ISBN: "1"}, book, "Book JSON unmarshalling wrong.")
}

func TestAllBooks(t *testing.T) {
	books := AllBooks()
	assert.Len(t, books, 3, "Wrong number of books.")
}

func TestCreateNewBook(t *testing.T) {
	book := Book{Name: "Book E", Author: "E", ISBN: "5"}
	isbn, created := CreateBook(book)
	assert.True(t, created, "Book was not created.")
	assert.Equal(t, "5", isbn, "Wrong ISBN.")
}

func TestDoNotCreateExistingBook(t *testing.T) {
	book := Book{ISBN: "1"}
	_, created := CreateBook(book)
	assert.False(t, created, "Book was created.")
}

func TestUpdateExistingBook(t *testing.T) {
	book := Book{Name: "Book B", Author: "B", ISBN: "1", Description: "Updated description"}
	updated := UpdateBook("1", book)
	assert.True(t, updated, "Contact not updated.")

	book, _ = GetBook("1")
	assert.Equal(t, "B", book.Author, "Author not updated.")
	assert.Equal(t, "Book B", book.Name, "Name not updated.")
	assert.Equal(t, "Updated description", book.Description, "Description not updated.")
}

func TestDeleteBook(t *testing.T) {
	DeleteBook("1")
	assert.Len(t, AllBooks(), 3, "Wrong number of books after delete.")
}
