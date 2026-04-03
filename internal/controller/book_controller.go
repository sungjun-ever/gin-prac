package controller

import (
	"slices"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

type Book struct {
	ID     string `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

var books = []Book{
	{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
}

func (b *BookController) GetBooks(c *gin.Context) {
	c.JSON(200, Response{
		Success: true,
		Data:    gin.H{"books": books},
	})
}

func (b *BookController) GetBook(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(200, Response{
				Success: true,
				Data:    gin.H{"book": book},
			})
			return
		}
	}
	c.JSON(404, Response{
		Success: false,
		Error:   gin.H{"message": "Book not found"},
	})
}

func (b *BookController) CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, Response{
			Success: false,
			Error:   gin.H{"message": err.Error()},
		})
		return
	}
	books = append(books, newBook)
	c.JSON(201, Response{
		Success: true,
		Data:    gin.H{"message": "Book Created", "book": newBook},
	})
}

func (b *BookController) UpdateBook(c *gin.Context) {
	var targetBook Book
	if err := c.ShouldBindJSON(&targetBook); err != nil {
		c.JSON(400, Response{
			Success: false,
			Error:   gin.H{"message": err.Error()},
		})
		return
	}

	for i, book := range books {
		if book.ID == targetBook.ID {
			books[i] = targetBook
			c.JSON(200, Response{
				Success: true,
				Data:    gin.H{"message": "Book Updated", "book": targetBook},
			})
			return
		}
	}

	c.JSON(404, Response{
		Success: false,
		Error:   gin.H{"message": "Book not found"},
	})
}

func (b *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID == id {
			slices.Delete(books, i, i+1)
			c.JSON(200, Response{
				Success: true,
				Data:    gin.H{"message": "Book Deleted", "books": books},
			})
			return
		}
	}

	c.JSON(404, Response{
		Success: false,
		Error:   gin.H{"message": "Book not found"},
	})
}
