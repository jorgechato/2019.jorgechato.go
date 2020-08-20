package application

import (
	"strconv"

	"jorgechato.com/pkg/books/domain"
)

func (b *BooksUseCase) GetBooksByScore(input *[]domain.Book, score int) *[]domain.Book {
	output := []domain.Book{}

	for _, val := range *input {
		if i, _ := strconv.Atoi(val.Score); i >= score {
			output = append(output, val)
		}
	}

	return &output
}
