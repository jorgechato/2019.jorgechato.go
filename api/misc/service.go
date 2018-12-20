package misc

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func getMiscList(first, offset int) []Misc {
	var miscList []Misc

	Session.Order("published_at desc").
		Offset(offset).
		Limit(first).
		Find(&miscList)

	return miscList
}

func createMisc(misc *Misc) {
	Session.Create(misc)
}

func updateMisc(misc *Misc) {
	Session.Model(misc).
		Omit("created_at").
		Omit("id").
		Update(*misc)
}

func getMiscByID(id string) Misc {
	var misc Misc

	Session.First(&misc, "id = ?", id)

	return misc
}
