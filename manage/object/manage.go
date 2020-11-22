package object

import (
	"errors"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/manage"
	"glusterfs-storage-gateway/meta"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

const (
	objectMetaFile   = "object.meta"
	blockFileMinSize = 1024 * 1024 * 256

)

type ObjectManage struct {

	api             *fs_api.FsApi
	conn            *redis.Conn
	doneCh          chan struct{}
	wg              *sync.WaitGroup
	BucketInfoCache map[string]*meta.BucketInfo
	NotifyCh        chan *meta.BucketInfo
	// each bucket maintain object meta,that such {object_key object_name object_info}
	ObjectMetaFile map[string]*fs_api.FsFd
	//how many write to one bucket Concurreny
	writeConcurrent  int
	//each bucket current max block,each block can support writeConcurrent writer
	Block map[string][]*BlockInfo
	//each bucket maintains 128 block fd and mutex
	BlockFile map[string][]*BlockFd
}

func NewObjectManage(api *fs_api.FsApi, BucketInfoCache map[string]*meta.BucketInfo, NotifyCh chan *meta.BucketInfo, conn *redis.Conn) (*ObjectManage, error) {
	if api == nil || conn == nil {
		return nil, errors.New("fsApi or conn is valid")
	}
	return &ObjectManage{
		api:             api,
		conn:            conn,
		doneCh:          make(chan struct{}),
		wg:              &sync.WaitGroup{},
		BucketInfoCache: BucketInfoCache,
		NotifyCh:        NotifyCh,
	}, nil
}
func (objectManage *ObjectManage) createObjectBlock(bucketName string) error {
	return nil
}
func (objectManage *ObjectManage) createObjectMeta(bucketName string) (*meta.BucketInfo, error) {
	var bucketInfo *meta.BucketInfo
	var err error
	if objectManage.BucketInfoCache[bucketName] == nil {
		bucketInfo, err = manage.FetchBucketInfo(objectManage.conn, bucketName)
		if err == nil {
			return nil, errors.New(fmt.Sprintf("%s not exeists", bucketName))
		}
	}
	if err = manage.CheckBucketStatus(bucketInfo); err != nil {
		return nil, err
	}
	if objectManage.ObjectMetaFile[bucketName] == nil {
		objectMetaFilePath := fmt.Sprintf("/%s/%s", bucketInfo.RealDirName, objectMetaFile)
		if err = objectManage.api.Stat(objectMetaFilePath); err != nil {
			objectManage.ObjectMetaFile[bucketName], err = objectManage.api.Creat(objectMetaFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		} else {
			objectManage.ObjectMetaFile[bucketName], err = objectManage.api.Open(objectMetaFilePath, os.O_RDWR|os.O_APPEND)
		}
	}
	return bucketInfo, err
}
