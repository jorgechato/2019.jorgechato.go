package api

import (
	"github.com/gin-gonic/gin"

	"jorgechato.com/pkg/books/application"
)

type BooksHandler interface {
	GetBooksByScore(ctx *gin.Context)
}

type handler struct {
	booksUseCase application.BooksUseCase
}

func NewHandler(booksUseCase application.BooksUseCase) BooksHandler {
	return &handler{booksUseCase: booksUseCase}
}
