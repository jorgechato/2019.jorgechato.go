package misc

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
	. "github.com/jorgechato/api.jorgechato.com/utils"
)

func getMiscList(first, offset int) []Misc {
	db, _ := gorm.Open("postgres", DB)
	defer db.Close()

	var miscList []Misc
	db.Order("published_at desc").
		Offset(offset).
		Limit(first).
		Find(&miscList)

	return miscList
}

func createMisc(misc *Misc) {
	db, _ := gorm.Open("postgres", DB)
	defer db.Close()

	db.Create(misc)
}

func updateMisc(misc *Misc) {
	db, _ := gorm.Open("postgres", DB)
	defer db.Close()

	db.Model(misc).
		Omit("created_at").
		Omit("id").
		Update(*misc)
}

func getMiscByID(id string) Misc {
	db, _ := gorm.Open("postgres", DB)
	defer db.Close()

	var misc Misc
	db.First(&misc, "id = ?", id)

	return misc
}
