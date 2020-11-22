package bucket

import (
	"context"
	"encoding/json"
	"fmt"
	"glusterfs-storage-gateway/meta"
	mgr "glusterfs-storage-gateway/manage"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	createBucketDirType = 0
	deleteBucketDirType = 1
)

func (manage *BucketManage) checkBucketExist(bucketName string) bool {
	return mgr.CheckBucketExist(manage.conn,bucketName)
}
func (manage *BucketManage) handleBucketDir(bucketName, bucketDirName string, bucketDirType int) error {
	var err error
	bucketDir := fmt.Sprintf("%s-%s", bucketName, bucketDirName)
	switch bucketDirType {
	case createBucketDirType:
		err = manage.api.Mkdir(bucketDir, 0755)
		break
	case deleteBucketDirType:
		manage.api.RmDir(bucketDir)
		break
	}
	return err
}
func (manage *BucketManage) persistenceBucketInfoToDisk(bucektName string, b []byte) error {
	s := fmt.Sprintf("%s\t%s\n", bucektName, string(b))
	if manage.bucketInfoFile == nil {
		log.Errorln("bucketInfoFile is nil ")
	}
	if _, err := manage.api.Write(manage.bucketInfoFile, []byte(s)); err != nil {
		return err
	}

	return nil
}
func (manage *BucketManage) storeBucketInfo(bucketInfo *meta.BucketInfo, OpType uint8) ([]byte, error) {
	b, err := json.Marshal(bucketInfo)
	if err != nil {
		return nil, err
	}
	defer func(err error) {
		if err != nil {
			log.Errorln("redis:",err)
		}
	}(err)
	if err = manage.persistenceBucketInfoToDisk(bucketInfo.Name, b); err != nil {
		return nil, err
	}
	if OpType == DeleteBucketType {
		//remove  origin bucketinfo
		if _, err = manage.conn.Del(context.Background(), bucketInfo.Name).Result(); err != nil {
			return nil, err
		}
	} else {
		if _, err = manage.conn.Set(context.Background(), bucketInfo.Name, b, -1).Result(); err != nil {
			return nil, err
		}
	}
	return b, nil
}

func (manage *BucketManage) fetchBucketInfo(bucket string) (*meta.BucketInfo, error) {
	if manage.bucketInfoCache[bucket] == nil {
		bucketInfo,err := mgr.FetchBucketInfo(manage.conn,bucket)
		if err != nil {
			return nil, err
		}
		manage.bucketInfoCache[bucket] = bucketInfo
	}
	return manage.bucketInfoCache[bucket], nil
}
func (manage *BucketManage) delBucketInfoAndBucketData(bucketInfoRequest *BucketInfoRequest, bucketInfo *meta.BucketInfo) error {
	bucketInfoResponse := &BucketInfoResponse{}
	if err := manage.api.RmAllFileFromPath(bucketInfo.RealDirName); err != nil {
		bucketInfoResponse.Err = nil
		return err
	}
	manage.conn.Del(context.Background(), bucketInfoRequest.Info.Name)
	bucketInfoRequest.Done <- bucketInfoResponse
	return nil
}
