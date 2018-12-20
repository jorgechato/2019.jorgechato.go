package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	. "github.com/jorgechato/api.jorgechato.com/utils"
)

var Session *gorm.DB

func init() {
	if Session, err := gorm.Open("postgres", DB); err == nil {
		Session.DB().SetConnMaxLifetime(time.Minute * 5)
		Session.DB().SetMaxIdleConns(5)
		Session.DB().SetMaxOpenConns(5)
	}
}

func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "go_" + defaultTableName
	}
}
