package controller

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/service"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func GetBuckets(p graphql.ResolveParams) (interface{}, error) {
	buckets := service.GetBuckets(
		p.Args["first"].(int),
		p.Args["offset"].(int),
	)

	return buckets, nil
}

func GetBucket(p graphql.ResolveParams) (interface{}, error) {
	bucket := service.GetBucketByID(p.Args["id"].(string))

	return bucket, nil
}

func DeleteBucket(p graphql.ResolveParams) (interface{}, error) {
	bucket := service.GetBucketByID(p.Args["id"].(string))

	if bucket.ID != "" {
		service.DeleteBucket(&bucket)

		return bucket, nil
	}

	return bucket, nil
}

func CleanBucket(p graphql.ResolveParams) (interface{}, error) {
	var bucket Bucket

	bucket.ID = p.Args["id"].(string)
	if bucket.ID == "" {
		return nil, nil
	}

	bucket = service.GetBucketByID(bucket.ID)
	bucket.Affiliates = []*Affiliate{}

	service.UpdateBucket(&bucket)

	return bucket, nil
}

func CreateBucket(p graphql.ResolveParams) (interface{}, error) {
	// TODO
}
