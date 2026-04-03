package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"play/internal/controller"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func DummyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token != "secret-token" {
			c.JSON(401, gin.H{"message": "인증 실패"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	b := controller.BookController{}

	v1 := r.Group("/api/v1")
	v1.Use(DummyAuthMiddleware())
	{
		v1.GET("/books", b.GetBooks)
		v1.GET("/books/:id", b.GetBook)
		v1.POST("/books", b.CreateBook)
		v1.PUT("/books/:id", b.UpdateBook)
		v1.DELETE("/books/:id", b.DeleteBook)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server gracefully stopped")
}
