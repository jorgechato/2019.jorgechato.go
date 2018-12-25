package controller

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/service"
)

func GetBuckets(p graphql.ResolveParams) (interface{}, error) {
	buckets := service.GetBuckets()

	return buckets, nil
}

func GetBucket(p graphql.ResolveParams) (interface{}, error) {
	bucket := service.GetBucketByID(p.Args["id"].(string))

	return bucket, nil
}
