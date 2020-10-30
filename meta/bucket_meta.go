package meta

import (
	"context"
	"errors"
	"gluster-gtw/bucket"
	"gluster-gtw/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type BucketMeta struct {
	LimitObject   uint64
	ObjectCount   uint64
	UsedCapacity  uint64
	TotalCapacity uint64
	Status        uint8
	BlockCount    uint64
}

func NewBucketMeta(limitObject, totalCapacity uint64) *BucketMeta {
	return &BucketMeta{
		LimitObject:   limitObject,
		ObjectCount:   0,
		UsedCapacity:  0,
		TotalCapacity: totalCapacity,
		Status:        bucket.UsedSatus,
	}
}

func GetBucketMeta(bucketName string) (*BucketMeta, error) {
	var b []byte
	var err error
	redisClient := utils.RedisClient
	if redisClient == nil {
		return nil, errors.New("redisClient is nil")
	}
	b, err = redisClient.Get(context.Background(), bucketName).Bytes()
	if err != nil {
		return nil, err
	}
	var bucketMata BucketMeta
	if err = bson.Unmarshal(b, &bucketMata); err != nil {
		return nil, err
	}
	return &bucketMata, nil
}
func (meta *BucketMeta) Store(bucketName string) error {
	b, err := bson.Marshal(meta)
	if err != nil {
		return err
	}
	ctx := context.Background()
	if _, err := utils.RedisClient.Set(ctx, bucketName, string(b), -1).Result(); err != nil {
		return err
	}
	return nil
}
func (meta *BucketMeta) Delete(bucketName string) error {
	meta.Status = bucket.DelStatus
	return meta.Store(bucketName)
}
