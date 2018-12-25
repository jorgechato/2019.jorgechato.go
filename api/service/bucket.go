package service

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func GetBucketByID(id string) Bucket {
	var bucket Bucket

	Session.First(&bucket, "id = ?", id)

	return bucket
}

func GetBucketByName(name string) Bucket {
	var bucket Bucket

	Session.
		Where(Bucket{Name: name}).
		FirstOrCreate(&bucket)

	return bucket
}

func UpdateBucket(bucket *Bucket) {
	Session.
		Model(bucket).
		Omit("created_at").
		Omit("id").
		Update(*bucket)
}

func GetBuckets() []Bucket {
	var buckets []Bucket

	Session.
		Order("created_at desc").
		Find(&buckets)

	return buckets
}
