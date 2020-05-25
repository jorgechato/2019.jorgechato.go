package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"jorgechato.com/pkg/location/domain"
)

func (l *Location) Encode(input *domain.Location) ([]byte, error) {
	if l.Tag == "" {
		l.Tag = TAG
	}

	json := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 l.Tag,
	}.Froze()

	rawMsg, err := json.Marshal(input)

	if err != nil {
		return nil, errors.Wrap(err, "serializer.Location.Encode")
	}
	return rawMsg, nil
}

func (t *Trip) Encode(input *domain.Trip) ([]byte, error) {
	if t.Tag == "" {
		t.Tag = TAG
	}

	json := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 t.Tag,
	}.Froze()

	rawMsg, err := json.Marshal(input)

	if err != nil {
		return nil, errors.Wrap(err, "serializer.Trip.Encode")
	}
	return rawMsg, nil
}
