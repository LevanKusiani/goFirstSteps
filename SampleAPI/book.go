package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Book One", Author: "Arthur Morgan", Quantity: 2},
	{ID: "2", Title: "Book Two", Author: "John Marston", Quantity: 3},
	{ID: "3", Title: "Book Three", Author: "Finn the Human", Quantity: 5},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context) {
	book, err := getBookById(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	book, getErr := getBookById(c.Query("id"))

	if getErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": getErr.Error()})
	}

	decErr := decrementBookQuantity(book)

	if decErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": decErr.Error()})

		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkinBook(c *gin.Context) {
	book, getErr := getBookById(c.Query("id"))

	if getErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": getErr.Error()})
	}

	incrementBookQuantity(book)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book returned, thank you!"})
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func decrementBookQuantity(b *book) error {
	if b.Quantity <= 0 {
		return errors.New("Out of books for the title: " + b.Title)
	}

	b.Quantity -= 1

	return nil
}

func incrementBookQuantity(b *book) {
	b.Quantity += 1
}
