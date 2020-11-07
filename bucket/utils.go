package bucket

import (
	"context"
	"fmt"
	"glusterfs-storage-gateway/meta"
	"go.mongodb.org/mongo-driver/bson"
	log "github.com/sirupsen/logrus"
)

const (
	createBucketDirType = 0
	deleteBucketDirType = 1
)

func (manage *BucketManage) checkBucketExist(bucketName string) bool {
	//defer log.Errorln("checkBucketExist err:",err)
	ret, err := manage.conn.Exists(context.Background(), bucketName).Result();
	if err != nil || ret > 0 {
		log.Debugf("the bucket:%v is exists", bucketName)
		return true
	}
	return false
}
func (manage *BucketManage) handleBucketDir(bucketName,bucketDirName string, bucketDirType int) error {
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

func (manage *BucketManage) storeBucketInfo(bucketInfo *meta.BucketInfo) (string, error) {
	var result string
	b, err := bson.Marshal(bucketInfo)
	if err != nil {
		return result, err
	}
	return manage.conn.Set(context.Background(), bucketInfo.Name, b, -1).Result()
}
func (manage *BucketManage) fetchBucketInfo(bucket string) (*meta.BucketInfo, error) {
	if manage.bucketInfoCache[bucket] == nil {
		binstr, err := manage.conn.Get(context.Background(), bucket).Result()
		if err != nil {
			return nil, err
		}
		bucketInfo := &meta.BucketInfo{}
		if err := bson.Unmarshal([]byte(binstr), bucketInfo); err != nil {
			return nil, err
		}
		manage.bucketInfoCache[bucket] = bucketInfo
	}
	return manage.bucketInfoCache[bucket], nil
}
func (manage *BucketManage) delBucketInfoAndBucketData(bucketInfoRequest *BucketInfoRequest, bucketDir string) error {
	bucketInfoResponse := &BucketInfoResponse{}
	if err := manage.api.RmAllFileFromPath(bucketDir); err != nil {
		bucketInfoResponse.Err = nil
		return err
	}
	manage.conn.Del(context.Background(), bucketInfoRequest.Info.Name)
	bucketInfoRequest.Done <- bucketInfoResponse
	return nil
}
