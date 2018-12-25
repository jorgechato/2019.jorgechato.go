package service

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func GetMiscList(first, offset int) []Misc {
	var miscList []Misc

	Session.Order("published_at desc").
		Offset(offset).
		Limit(first).
		Find(&miscList)

	return miscList
}

func GetMiscByID(id string) Misc {
	var misc Misc

	Session.First(&misc, "id = ?", id)

	return misc
}

func CreateMisc(misc *Misc) {
	Session.Create(misc)
}

func UpdateMisc(misc *Misc) {
	Session.Model(misc).
		Omit("created_at").
		Omit("id").
		Update(*misc)
}
