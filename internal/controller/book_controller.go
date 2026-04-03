package controller

import (
	"play/internal/model"
	"play/internal/service"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	Service service.BookService
}

type BookUriParams struct {
	ID int `uri:"id" binding:"required"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{Service: service}
}

func (b *BookController) GetBooks(c *gin.Context) {
	books, err := b.Service.GetAllBooks()

	if err != nil {
		c.JSON(404, Response{
			Success: false,
			Error:   gin.H{"message": "Books not found"},
		})
		return
	}

	c.JSON(200, Response{
		Success: true,
		Data:    gin.H{"books": books},
	})
}

func (b *BookController) GetBook(c *gin.Context) {
	var uri BookUriParams
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, Response{
			Success: false,
			Error:   gin.H{"message": "Invalid URI parameters"},
		})
		return
	}

	book, err := b.Service.GetBook(uri.ID)

	if err != nil {
		c.JSON(404, Response{
			Success: false,
			Error:   gin.H{"message": "Book not found"},
		})
		return
	}

	c.JSON(200, Response{
		Success: true,
		Data:    gin.H{"book": book},
	})
}

func (b *BookController) CreateBook(c *gin.Context) {
	var newBook model.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, Response{
			Success: false,
			Error:   gin.H{"message": err.Error()},
		})
		return
	}

	if err := b.Service.CreateBook(&newBook); err != nil {
		c.JSON(500, Response{
			Success: false,
			Error:   gin.H{"message": err.Error()},
		})
		return
	}

	c.JSON(201, Response{
		Success: true,
		Data:    gin.H{"message": "Book Created", "book": newBook},
	})
}

func (b *BookController) UpdateBook(c *gin.Context) {
	var targetBook model.Book

	if err := c.ShouldBindJSON(&targetBook); err != nil {
		c.JSON(400, Response{
			Success: false,
			Error:   gin.H{"message": err.Error()},
		})
		return
	}

	err := b.Service.UpdateBook(&targetBook)

	if err != nil {
		c.JSON(500, Response{
			Success: false,
			Error:   gin.H{"message": "Internal Server Error"},
		})
		return
	}

	c.JSON(201, Response{
		Success: true,
		Data:    gin.H{"message": "Book Updated", "book": targetBook},
	})

}

func (b *BookController) DeleteBook(c *gin.Context) {
	var uri BookUriParams

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, Response{
			Success: false,
			Error:   gin.H{"message": "Invalid URI parameters"},
		})
		return
	}

	err := b.Service.DeleteBook(uri.ID)

	if err != nil {
		c.JSON(500, Response{
			Success: false,
			Error:   gin.H{"message": "Internal Server Error"},
		})

		return
	}

	c.JSON(200, Response{
		Success: true,
		Data:    gin.H{"message": "Book Deleted"},
	})
}
