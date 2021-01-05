package xml

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"jorgechato.com/pkg/books/domain"
)

func (b *Books) Decode(input []byte) (*[]domain.Book, error) {
	res := &domain.GoodreadsResponse{}

	if err := xml.Unmarshal(input, &res); err != nil {
		return nil, errors.Wrap(err, "serializer.Books.Decode")
	}

	for key := range res.Books {
		if strings.Contains(res.Books[key].Metadata.Cover, NoPhoto) {
			res.Books[key].Metadata.Cover = fmt.Sprintf(
				BaseCover,
				res.Books[key].Metadata.ISBN,
			)
		}
	}

	return &res.Books, nil
}
