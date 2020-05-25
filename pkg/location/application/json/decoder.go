package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"jorgechato.com/pkg/location/domain"
)

func (l *Location) Decode(input []byte) (*domain.Location, error) {
	if l.Tag == "" {
		l.Tag = TAG
	}

	json := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 l.Tag,
	}.Froze()

	location := &domain.Location{}

	if err := json.Unmarshal(input, &location); err != nil {
		return nil, errors.Wrap(err, "serializer.Location.Decode")
	}

	return location, nil
}

func (t *Trip) Decode(input []byte) (*domain.Trip, error) {
	if t.Tag == "" {
		t.Tag = TAG
	}

	json := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 t.Tag,
	}.Froze()

	trip := &domain.Trip{}

	if err := json.Unmarshal(input, &trip); err != nil {
		return nil, errors.Wrap(err, "serializer.Trip.Decode")
	}

	return trip, nil
}
