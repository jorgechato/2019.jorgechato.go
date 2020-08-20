package domain

type BooksSerializer interface {
	Decode(input []byte) (*[]Book, error)
}
