package affiliates

import (
	"strings"

	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/buckets"
	"github.com/jorgechato/api.jorgechato.com/api/tags"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func delete(p graphql.ResolveParams) (interface{}, error) {
	affiliate := getAffiliateByID(p.Args["id"].(string))

	if affiliate.ID != "" {
		deleteAffiliate(&affiliate)

		return affiliate, nil
	}

	return nil, nil
}

func get(p graphql.ResolveParams) (interface{}, error) {
	affiliate := getAffiliateByID(p.Args["id"].(string))

	return affiliate, nil
}

func getList(p graphql.ResolveParams) (interface{}, error) {
	articles := getAffiliates(
		p.Args["first"].(int),
		p.Args["offset"].(int),
	)

	return articles, nil
}

func create(p graphql.ResolveParams) (interface{}, error) {
	var affiliate Affiliate
	var bucket Bucket

	affiliate.Name = p.Args["name"].(string)
	affiliate.Url = p.Args["url"].(string)
	affiliate.Public = p.Args["public"].(bool)
	affiliate.Thumbmail = p.Args["thumbnail"].(string)

	if obj, ok := p.Args["preview"].(string); ok == true {
		affiliate.Preview = obj
	}

	if objs, ok := p.Args["tags"].([]interface{}); ok == true {
		for _, obj := range objs {
			tag := tags.GetTagByName(strings.ToLower(obj.(string)))
			affiliate.Tags = append(affiliate.Tags, &tag)
		}
	}

	bucket = buckets.GetBucketByName(strings.ToLower(p.Args["bucket"].(string)))
	bucket.Affiliates = append(bucket.Affiliates, &affiliate)

	createAffiliate(&affiliate)
	buckets.UpdateBucket(&bucket)

	return affiliate, nil
}

func update(p graphql.ResolveParams) (interface{}, error) {
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

	if obj, ok := p.Args["preview"].(string); ok == true {
		affiliate.Preview = obj
	}

	if objs, ok := p.Args["tags"].([]interface{}); ok == true {
		for _, obj := range objs {
			tag := tags.GetTagByName(strings.ToLower(obj.(string)))
			affiliate.Tags = append(affiliate.Tags, &tag)
		}
	}

	bucket = buckets.GetBucketByName(strings.ToLower(p.Args["bucket"].(string)))
	bucket.Affiliates = append(bucket.Affiliates, &affiliate)

	updateAffiliate(&affiliate)
	buckets.UpdateBucket(&bucket)

	return affiliate, nil
}
