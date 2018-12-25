package service

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

func GetTagsByAffiliate(affiliate Affiliate) []*Tag {
	var tags []*Tag

	Session.
		Model(&affiliate).
		Related(&tags, "Tags")

	return tags
}

func GetTagsByArticle(article Article) []*Tag {
	var tags []*Tag

	Session.
		Model(&article).
		Related(&tags, "Tags")

	return tags
}
