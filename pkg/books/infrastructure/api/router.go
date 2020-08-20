package api

import (
	"github.com/gin-gonic/gin"

	"jorgechato.com/pkg/books/application"
)

func Router(BASE string, router *gin.Engine) {
	var service application.BooksUseCase
	handler := NewHandler(service)

	v := router.Group(BASE + "/books")
	{
		v.GET("/", handler.GetBooksByScore)
	}
}
