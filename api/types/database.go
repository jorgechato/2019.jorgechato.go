package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	. "github.com/jorgechato/api.jorgechato.com/utils"
)

var Session *gorm.DB

func init() {
	Session, _ = gorm.Open("postgres", DB)

	Session.DB().SetConnMaxLifetime(time.Minute * 5)
	Session.DB().SetMaxIdleConns(5)
	Session.DB().SetMaxOpenConns(5)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "go_" + defaultTableName
	}
}
