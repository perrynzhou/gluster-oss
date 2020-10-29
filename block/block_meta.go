package block

import (
	"context"
	"errors"
	"fmt"
	fs_api "gluster-gtw/fs-api"
	"gluster-gtw/utils"
	"go.mongodb.org/mongo-driver/bson"
	"sync/atomic"
)

type BlockMeta struct {
	BucketName  string
	Index       uint64
	StartOffset uint64
	FreeLen     uint64
}

func NewBlockMeta(index, freeLen uint64,bucketName string) *BlockMeta {
	return  &BlockMeta{
		BucketName:bucketName,
		Index:       index,
		StartOffset: uint64(0),
		FreeLen:     freeLen,
	}
}

func (meta *BlockMeta) Store() error {
	b, err := bson.Marshal(meta)
	if err != nil {
		return err
	}
	ctx := context.Background()

	metaKey := fmt.Sprintf("%s.meta",meta.BucketName)
	if _, err := utils.RedisClient.Set(ctx, metaKey, string(b), -1).Result(); err != nil {
		return err
	}
	atomic.AddUint64(&meta.Index,1)
	return nil
}

func (meta *BlockMeta) Delete(api *fs_api.FsApi) error {
	metaKey := fmt.Sprintf("%s.meta",meta.BucketName)
	ctx := context.Background()
	if _, err := utils.RedisClient.Del(ctx, metaKey).Result(); err != nil {
		return err
	}
	return nil
}
func GetBucketBlockMeta(bucketName string) (*BlockMeta,error) {
	redisClient := utils.RedisClient
	if redisClient == nil {
		return nil, errors.New("redisClient is nil")
	}
	ctx := context.Background()
	metaKey := fmt.Sprintf("%s.meta",bucketName)
	b, err := redisClient.Get(ctx, metaKey).Bytes()
	if err != nil {
		return nil,err
	}
	meta :=&BlockMeta{}
	if err :=bson.Unmarshal(b,meta);err != nil {
		return nil,err
	}
	return meta,nil
}