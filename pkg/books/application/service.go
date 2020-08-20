package application

import "jorgechato.com/pkg/books/domain"

type BooksUseCase struct {
	booksUseCase domain.BooksUseCase
}

func NewBooksService(service domain.BooksUseCase) BooksUseCase {
	return BooksUseCase{
		service,
	}
}
