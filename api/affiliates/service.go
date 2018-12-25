package affiliates

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func deleteAffiliate(affiliate *Affiliate) {
	Session.
		Delete(affiliate)
}

func getAffiliates(first, offset int) []Affiliate {
	var articles []Affiliate

	Session.
		Order("created_at desc").
		Offset(offset).
		Limit(first).
		Find(&articles)

	return articles
}

func getAffiliateByID(id string) Affiliate {
	var affiliate Affiliate

	Session.First(&affiliate, "id = ?", id)

	return affiliate
}

func createAffiliate(affiliate *Affiliate) {
	Session.Create(affiliate)
}

func updateAffiliate(affiliate *Affiliate) {
	Session.
		Model(affiliate).
		Omit("created_at").
		Omit("id").
		Update(*affiliate)
}

func getTags(affiliate Affiliate) []*Tag {
	var tags []*Tag

	Session.
		Model(&affiliate).
		Related(&tags, "Tags")

	return tags
}
