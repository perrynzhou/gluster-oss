package object

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/manage"
	"glusterfs-storage-gateway/meta"
	"os"
	"sync"
)
const (
	objectMetaFile = "object.meta"
	blockFileMinSize = 1024*1024*256
)
type ObjectManage struct {
	api     *fs_api.FsApi
	conn    *redis.Conn
	doneCh  chan struct{}
	wg      *sync.WaitGroup
	BucketInfoCache map[string]*meta.BucketInfo
	NotifyCh     chan *meta.BucketInfo
	ObjectMetaFile  map[string]*fs_api.FsFd

}

func NewObjectManage(api *fs_api.FsApi, BucketInfoCache map[string]*meta.BucketInfo,NotifyCh chan *meta.BucketInfo,conn *redis.Conn) (*ObjectManage, error) {
	if api == nil || conn == nil {
		return nil, errors.New("fsApi or conn is valid")
	}
	return &ObjectManage{
		api:     api,
		conn:    conn,
		doneCh:  make(chan struct{}),
		wg:      &sync.WaitGroup{},
		BucketInfoCache: BucketInfoCache,
		NotifyCh:NotifyCh,
	}, nil
}
func (objectManage *ObjectManage)createObjectMeta(bucketName string) error  {
	 var bucketInfo  *meta.BucketInfo
	 var err error
	 if objectManage.BucketInfoCache[bucketName]== nil {
		 bucketInfo,err = manage.FetchBucketInfo(objectManage.conn,bucketName)
		 if err == nil {
		 	return errors.New(fmt.Sprintf("%s not exeists",bucketName))
		 }
	 }
	 if err =manage.CheckBucketStatus(bucketInfo) ;err != nil  {
	 	return err
	 }
	 if objectManage.ObjectMetaFile[bucketName] == nil {
		 objectMetaFilePath := fmt.Sprintf("/%s/%s",bucketInfo.RealDirName,objectMetaFile)
		 if err=objectManage.api.Stat(objectMetaFilePath);err != nil {
			 objectManage.ObjectMetaFile[bucketName],err =objectManage.api.Creat(objectMetaFilePath,os.O_CREATE|os.O_RDWR|os.O_APPEND,os.ModePerm)
		 }else {
			 objectManage.ObjectMetaFile[bucketName],err =objectManage.api.Open(objectMetaFilePath,os.O_RDWR|os.O_APPEND)
		 }
	 }
	 return err
}