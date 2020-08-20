package domain

type BooksUseCase interface {
	GetBooksByScore(*[]Book, int) *[]Book
}
