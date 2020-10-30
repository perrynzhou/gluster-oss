package bucket

import (
	"errors"
	"fmt"
	fs_api "gluster-gtw/fs-api"
	"gluster-gtw/meta"
	"gluster-gtw/utils"
)


type Bucket struct {
	Name    string
	Subffix string
	Meta    *meta.BucketMeta
	Index    uint64
	Count   uint64
}

func NewBucket(api *fs_api.FsApi, name, subffix string, limitObject, totalCapacity uint64) (*Bucket, error) {
	bucketPath := fmt.Sprintf("/%s", name)
	if err := checkBucketExist(api,bucketPath);err != nil {
		if err := api.Mkdir(name, 0664); err != nil {
			return nil, err
		}
	}else {
		return nil,errors.New(fmt.Sprintf("%s already exist"))
	}
	meta := meta.NewBucketMeta(limitObject,totalCapacity)
	bucket:= &Bucket{
		Name:    name,
		Subffix: subffix,
		Meta:    meta,
		Index:0,
		Count:0,
	}
	bucketBlockMetaFile := metaKey := fmt.Sprintf("%s.meta",meta.BucketName)
	return bucket,nil
}
func (bucket *Bucket) StoreMeta() error {
	if utils.RedisClient == nil {
		return errors.New("redisPool is init")
	}
	return bucket.Meta.Store(bucket.Name)
}
func (bucket *Bucket) GetMeta() (*meta.BucketMeta, error) {
	return meta.GetBucketMeta(bucket.Name)
}
func (bucket *Bucket)Delete() error {
	meta,err := meta.GetBucketMeta(bucket.Name)
	if err != nil {
		return err
	}
	meta.Status = DelStatus
	bucket.Meta = meta
	return bucket.StoreMeta()
}
