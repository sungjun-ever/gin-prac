package router

import (
	"play/internal/middlware"
	"play/registry"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, ct *registry.Container) {
	v1 := r.Group("/api/v1")
	v1.Use(middlware.DummyAuthMiddleware())
	{
		v1.GET("/books", ct.BookController.GetBooks)
		v1.GET("/books/:id", ct.BookController.GetBook)
		v1.POST("/books", ct.BookController.CreateBook)
		v1.PUT("/books/:id", ct.BookController.UpdateBook)
		v1.DELETE("/books/:id", ct.BookController.DeleteBook)
	}
}
