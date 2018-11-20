package types

import (
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Model struct {
	ID        string    `gorm:"primary_key;type:char(36);column:id" json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

var re = regexp.MustCompile("[^a-z0-9]+")

func slugify(s string) string {
	return strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "go_" + defaultTableName
	}
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", uuid.Must(uuid.NewV4()).String())
	return nil
}
