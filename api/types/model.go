package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        string    `gorm:"primary_key;type:char(32);column:id" json:"id"`
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
	hasher := md5.New()
	hasher.Write(
		[]byte(
			fmt.Sprintf(
				"%v:%v",
				model.CreatedAt,
				time.Now().UnixNano(),
			),
		),
	)

	scope.SetColumn(
		"id",
		hex.EncodeToString(hasher.Sum(nil)),
	)

	return nil
}
