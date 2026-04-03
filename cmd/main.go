package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"play/config"
	"play/database"
	"play/internal/model"
	"play/registry"
	"play/router"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := database.ConnectDB()
	err := db.AutoMigrate(&model.Book{})

	if err != nil {
		log.Println(err.Error())
	}

	container := registry.NewContainer(db)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	r := gin.Default()
	router.SetupRouter(r, container)

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
