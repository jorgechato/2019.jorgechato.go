package service

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func DeleteBucket(bucket *Bucket) {
	Session.
		Delete(bucket)
}

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

func CreateBucket(bucket *Bucket) {
	Session.Create(bucket)
}

func UpdateBucket(bucket *Bucket) {
	Session.
		Model(bucket).
		Omit("created_at").
		Omit("id").
		Update(*bucket)
}

func GetBuckets(first, offset int) []Bucket {
	var buckets []Bucket

	Session.
		Order("created_at desc").
		Offset(offset).
		Limit(first).
		Find(&buckets)

	return buckets
}
