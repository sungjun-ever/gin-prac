package registry

import (
	"play/internal/controller"
	"play/internal/repository"
	"play/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	BookController *controller.BookController
}

func NewContainer(db *gorm.DB) *Container {
	bookRepo := repository.NewBookRepository(db)

	bookSvc := service.NewBookService(bookRepo)
	
	return &Container{
		BookController: controller.NewBookController(*bookSvc),
	}
}
