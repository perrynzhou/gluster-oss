package manage

import (
	"errors"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
	"sync"

	"github.com/google/uuid"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type ObjectManage struct {
	api            *fs_api.FsApi
	conn           *redis.Conn
	ReqCh          chan *ObjectRequest
	doneCh         chan struct{}
	wg             *sync.WaitGroup
	BucketCache    map[string]*Bucket
	ObjectCache    map[string]*Object
	notifyObjectCh chan *Object
	notifyBucketCh chan *Bucket
}

func NewObjectManage(api *fs_api.FsApi, bucketCache map[string]*Bucket, notifyBucketCh chan *Bucket, objectRequestCh chan *ObjectRequest, wg *sync.WaitGroup) *ObjectManage {
	return &ObjectManage{
		api:            api,
		ReqCh:          objectRequestCh,
		wg:             wg,
		notifyBucketCh: notifyBucketCh,
		doneCh:         make(chan struct{}),
		notifyObjectCh: make(chan *Object),
		ObjectCache:    make(map[string]*Object),
		BucketCache:    bucketCache,
	}
}

func (objectManage *ObjectManage) handlePutObjectRequest(request *ObjectRequest) error {
	var err error
	var bucket *Bucket
	var ok bool
	response := &ObjectResponse{
		Reply: request.Info,
	}
	defer func(request *ObjectRequest, response *ObjectResponse) {
		request.Done <- response
	}(request, response)
	log.Errorln("create object:", request.Info.Name)
	if bucket, ok = objectManage.BucketCache[request.Info.Bucket]; !ok {
		err = errors.New(fmt.Sprintf("bucket %s not  exists", request.Info.Name))
		response.Err = err
		return err
	}
	objectKey := uuid.New().String()
	blockFile := bucket.GetBlockFile()
	objectInfo, err := NewObject(objectManage.api, blockFile, bucket, request.Info)
	objectInfo.Meta.BlockId = blockFile.Meta.Index
	objectInfo.Meta.StartPos, err = blockFile.File.Size()
	wByets, err := PutData(request.LocalFile, objectManage.api, objectInfo)
	if err != nil {
		response.Err = err
		return err
	}
	objectInfo.Meta.Size = wByets
	objectInfo.Meta.Key = objectKey
	objectInfo.Meta.Status = meta.ActiveObjectStatus
	objectManage.notifyObjectCh <- objectInfo
	objectManage.notifyBucketCh <- bucket
	log.Warningln("bucketManage.notifyCh <- bucket:", bucket)
	return nil

}
func (objectManage *ObjectManage) handleError(request *ObjectRequest, response *ObjectResponse, err error) {
	response.Reply = request.Info
	response.Err = err
	if err != nil {
		log.Errorln("happen err:", err)
	}
	request.Done <- response
}
func (objectManage *ObjectManage) handleGetBucketRequest(request *ObjectRequest) error {

	response := &ObjectResponse{
		Reply: request.Info,
	}
	defer func(request *ObjectRequest, response *ObjectResponse) {
		request.Done <- response
	}(request, response)

	return nil
}
func (objectManage *ObjectManage) handleDeleteObjectRequest(request *ObjectRequest) error {
	var ok bool
	var object *Object
	var err error
	response := &ObjectResponse{
		Reply: request.Info,
	}
	defer func(request *ObjectRequest, response *ObjectResponse) {
		request.Done <- response
	}(request, response)
	if object, ok = objectManage.ObjectCache[request.Info.Key]; !ok {
		err = errors.New(fmt.Sprintf("object %s not exists in bucket %s", request.Info.Key, request.Info.Bucket))
		response.Err = err
		return err
	}
	object.Meta.Status = meta.InactiveObjectStatus
	objectManage.notifyBucketCh <- object.bucket
	objectManage.notifyObjectCh <- object
	return nil
}
func (objectManage *ObjectManage) Run() {
	objectManage.wg.Add(2)
	go objectManage.handleObjectRequest()
	go objectManage.handleObjectMeta()

}
func (objectManage *ObjectManage) handleObjectMeta() {
	log.Infoln("run ObjectService handleObjectMeta")
	defer objectManage.wg.Done()
	for {
		select {
		case object := <-objectManage.notifyObjectCh:
			sign := int64(1)
			if object.Meta.Status == meta.InactiveObjectStatus {
				delete(objectManage.ObjectCache, object.Meta.Key)
				object.bucket.SubObjectCounter()
				sign = int64(-1)
			} else {
				object.bucket.AddObjectCounter()
			}
			object.StoreMeta(objectManage.api, object.bucket.ObjectMetaFile)
			object.bucket.ModifyObjectBytes(sign * object.Meta.Size)
			break
		case <-objectManage.doneCh:
			return
		}
	}
}
func (objectManage *ObjectManage) handleObjectRequest() {
	log.Infoln("run ObjectService handleObjectRequest")
	objectManage.wg.Done()
	for {
		select {
		case req := <-objectManage.ReqCh:
			log.Infoln("recive request:", req)
			switch req.RequestType {
			case PutObjectType:
				objectManage.handlePutObjectRequest(req)
				break
			case GetObjectType:
				objectManage.handleGetBucketRequest(req)
				break
			case DelObjectType:
				objectManage.handleDeleteObjectRequest(req)
				break
			}
		case <-objectManage.doneCh:
			return
		}
	}
}
func (objectManage *ObjectManage) Stop() {
	for i := 0; i < 2; i++ {
		objectManage.doneCh <- struct{}{}
	}
}
