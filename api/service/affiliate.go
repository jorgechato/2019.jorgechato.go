package service

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func DeleteAffiliate(affiliate *Affiliate) {
	Session.
		Delete(affiliate)
}

func GetAffiliates(first, offset int) []Affiliate {
	var articles []Affiliate

	Session.
		Order("created_at desc").
		Offset(offset).
		Limit(first).
		Find(&articles)

	return articles
}

func GetAffiliateByID(id string) Affiliate {
	var affiliate Affiliate

	Session.First(&affiliate, "id = ?", id)

	return affiliate
}

func CreateAffiliate(affiliate *Affiliate) {
	Session.Create(affiliate)
}

func UpdateAffiliate(affiliate *Affiliate) {
	Session.
		Model(affiliate).
		Omit("created_at").
		Omit("id").
		Update(*affiliate)
}

func GetAffiliatesByBucket(bucket Bucket, first, offset int) []*Affiliate {
	var affiliates []*Affiliate

	Session.
		Offset(offset).
		Limit(first).
		Model(&bucket).
		Related(&affiliates, "Affiliates")

	return affiliates
}

func GetBestSellerByBucket(bucket Bucket) *Affiliate {
	var affiliates []*Affiliate

	Session.
		Order("score desc").
		Model(&bucket).
		Limit(1).
		Related(&affiliates, "Affiliates")

	return affiliates[0]
}
