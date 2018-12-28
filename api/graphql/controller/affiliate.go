package controller

import (
	"strings"

	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/service"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func DeleteAffiliate(p graphql.ResolveParams) (interface{}, error) {
	affiliate := service.GetAffiliateByID(p.Args["id"].(string))

	if affiliate.ID != "" {
		service.DeleteAffiliate(&affiliate)

		return affiliate, nil
	}

	return nil, nil
}

func GetAffiliate(p graphql.ResolveParams) (interface{}, error) {
	affiliate := service.GetAffiliateByID(p.Args["id"].(string))

	return affiliate, nil
}

func GetAffiliates(p graphql.ResolveParams) (interface{}, error) {
	if obj, ok := p.Source.(Bucket); ok == true {
		return service.GetAffiliatesByBucket(
			obj,
			p.Args["first"].(int),
			p.Args["offset"].(int),
		), nil
	}

	return service.GetAffiliates(
		p.Args["first"].(int),
		p.Args["offset"].(int),
	), nil
}

func CreateAffiliate(p graphql.ResolveParams) (interface{}, error) {
	var affiliate Affiliate
	var bucket Bucket

	affiliate.Name = p.Args["name"].(string)
	affiliate.Url = p.Args["url"].(string)
	affiliate.Public = p.Args["public"].(bool)
	affiliate.Thumbmail = p.Args["thumbnail"].(string)
	affiliate.Score = p.Args["score"].(int)

	if obj, ok := p.Args["preview"].(string); ok == true {
		affiliate.Preview = obj
	}

	if objs, ok := p.Args["tags"].([]interface{}); ok == true {
		for _, obj := range objs {
			tag := service.GetTagByName(strings.ToLower(obj.(string)))
			affiliate.Tags = append(affiliate.Tags, &tag)
		}
	}

	bucket = service.GetBucketByName(strings.ToLower(p.Args["bucket"].(string)))
	bucket.Affiliates = append(bucket.Affiliates, &affiliate)

	service.CreateAffiliate(&affiliate)
	service.UpdateBucket(&bucket)

	return affiliate, nil
}

func UpdateAffiliate(p graphql.ResolveParams) (interface{}, error) {
	var affiliate Affiliate
	var bucket Bucket

	affiliate.ID = p.Args["id"].(string)
	if affiliate.ID == "" {
		return nil, nil
	}

	affiliate.Name = p.Args["name"].(string)
	affiliate.Url = p.Args["url"].(string)
	affiliate.Public = p.Args["public"].(bool)
	affiliate.Thumbmail = p.Args["thumbnail"].(string)
	affiliate.Score = p.Args["score"].(int)

	if obj, ok := p.Args["preview"].(string); ok == true {
		affiliate.Preview = obj
	}

	if objs, ok := p.Args["tags"].([]interface{}); ok == true {
		for _, obj := range objs {
			tag := service.GetTagByName(strings.ToLower(obj.(string)))
			affiliate.Tags = append(affiliate.Tags, &tag)
		}
	}

	bucket = service.GetBucketByName(strings.ToLower(p.Args["bucket"].(string)))
	bucket.Affiliates = append(bucket.Affiliates, &affiliate)

	service.UpdateAffiliate(&affiliate)
	service.UpdateBucket(&bucket)

	return affiliate, nil
}
