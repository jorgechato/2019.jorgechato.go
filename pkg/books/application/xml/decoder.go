package xml

import (
	"encoding/xml"

	"github.com/pkg/errors"

	"jorgechato.com/pkg/books/domain"
)

func (b *Books) Decode(input []byte) (*[]domain.Book, error) {
	res := &domain.GoodreadsResponse{}

	if err := xml.Unmarshal(input, &res); err != nil {
		return nil, errors.Wrap(err, "serializer.Books.Decode")
	}

	return &res.Books, nil
}
