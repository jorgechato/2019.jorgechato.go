package utils

import (
	"regexp"
	"strings"
)

type Model interface {
	Save()
}

var re = regexp.MustCompile("[^a-z0-9]+")

func slugify(s string) string {
	return strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
