package tags

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func GetTagByName(name string) Tag {
	var tag Tag

	Session.
		Where(Tag{Name: name}).
		FirstOrCreate(&tag)

	return tag
}
