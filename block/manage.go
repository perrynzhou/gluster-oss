package block

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
	"golang.org/x/net/context"
	"strconv"
	"sync"
)

type BlockManage struct {
	api     *fs_api.FsApi
	conn    *redis.Conn
	doneCh  chan struct{}
	wg      *sync.WaitGroup
	buckets map[string]*meta.BucketInfo
}

func NewBlockManage(api *fs_api.FsApi, conn *redis.Conn) (*BlockManage, error) {
	if api == nil || conn == nil {
		return nil, errors.New("fsApi or conn is valid")
	}
	return &BlockManage{
		api:     api,
		conn:    conn,
		doneCh:  make(chan struct{}),
		wg:      &sync.WaitGroup{},
		buckets: make(map[string]*meta.BucketInfo),
	}, nil
}
func (blockManage *BlockManage) fetchBucketBlockIndex(bucketName string) (uint64, error) {
	//bucket block key:bucket.idx
	blockIdxKey := fmt.Sprintf("%s.bid", bucketName)
	blockIdxVal, err := blockManage.conn.Get(context.Background(), blockIdxKey).Result()
	if err != nil {
		return 0, err
	}
	blockId, err := strconv.ParseUint(blockIdxVal, 10, 64)
	if err != nil {
		return 0, err
	}
	return blockId, nil
}
func (blockManage *BlockManage) creatBlock(bucketName string) error {
	return nil

}
func (blockManage *BlockManage) Run() {

}
