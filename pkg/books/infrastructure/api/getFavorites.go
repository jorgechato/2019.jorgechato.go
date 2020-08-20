package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jorgechato.com/pkg/books/application/xml"
	"jorgechato.com/pkg/books/infrastructure/goodreads"
)

type bookQuery struct {
	Score int `form:"score"`
}

func (h handler) GetBooksByScore(ctx *gin.Context) {
	book := bookQuery{
		Score: 4,
	}
	ctx.ShouldBind(&book)

	raw, err := goodreads.FetchBooks()

	if err != nil {
		ctx.AbortWithError(
			http.StatusInternalServerError,
			err,
		)
	}

	xmlBooks := &xml.Books{}
	books, _ := xmlBooks.Decode(raw)

	ctx.JSON(
		http.StatusOK,
		h.booksUseCase.GetBooksByScore(books, book.Score),
	)
}
