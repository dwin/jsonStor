package model

import (
	"github.com/rs/xid"
	"time"
)

func GenerateBucketID() string {
	return xid.New().String()
}

type Bucket struct {
	ID                  string
	UserID              uint32
	Private             bool
	SecretKeyGenerated  time.Time
	SecretKeyExpiration time.Time
	InternalNotes       string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
}

func CreateBucket(bucket Bucket) (newBucket Bucket, err error) {
	return
}

func ReadBucket(id string) (bucket Bucket, err error) {
	err = db.Model(&bucket).Where("id=?", id).Select()
	if err != nil {
		return
	}
	return
}


