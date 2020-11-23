package bucket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"glusterfs-storage-gateway/meta"
	"go.mongodb.org/mongo-driver/bson"
"errors"
	log "github.com/sirupsen/logrus"
)

const (
	createBucketDirType = 0
	deleteBucketDirType = 1
)

func (bucketManage *BucketManage) checkBucketExist(bucketName string) bool {
	return CheckBucketExist(bucketManage.conn, bucketName)
}
func (bucketManage *BucketManage) handleBucketDir(bucketName, bucketDirName string, bucketDirType int) error {
	var err error
	bucketDir := fmt.Sprintf("%s-%s", bucketName, bucketDirName)
	switch bucketDirType {
	case createBucketDirType:
		err = bucketManage.api.Mkdir(bucketDir, 0755)
		break
	case deleteBucketDirType:
		bucketManage.api.RmDir(bucketDir)
		break
	}
	return err
}
func (bucketManage *BucketManage) persistenceBucketInfoToDisk(bucektName string, b []byte) error {
	s := fmt.Sprintf("%s\t%s\n", bucektName, string(b))
	if bucketManage.bucketInfoFile == nil {
		log.Errorln("bucketInfoFile is nil ")
	}
	if _, err := bucketManage.api.Write(bucketManage.bucketInfoFile, []byte(s)); err != nil {
		return err
	}

	return nil
}
func (bucketManage *BucketManage) storeBucketInfo(bucketInfo *meta.BucketInfo, OpType uint8) ([]byte, error) {
	b, err := json.Marshal(bucketInfo)
	if err != nil {
		return nil, err
	}
	defer func(err error) {
		if err != nil {
			log.Errorln("redis:", err)
		}
	}(err)
	if err = bucketManage.persistenceBucketInfoToDisk(bucketInfo.Name, b); err != nil {
		return nil, err
	}
	if OpType == DeleteBucketType {
		//remove  origin bucketinfo
		if _, err = bucketManage.conn.Del(context.Background(), bucketInfo.Name).Result(); err != nil {
			return nil, err
		}
	} else {
		if _, err = bucketManage.conn.Set(context.Background(), bucketInfo.Name, b, -1).Result(); err != nil {
			return nil, err
		}
	}
	return b, nil
}

func (bucketManage *BucketManage) fetchBucketInfo(bucket string) (*meta.BucketInfo, error) {
	if bucketManage.BucketInfoCache[bucket] == nil {
		bucketInfo, err := FetchBucketInfo(bucketManage.conn, bucket)
		if err != nil {
			return nil, err
		}
		bucketManage.BucketInfoCache[bucket] = bucketInfo
	}
	return bucketManage.BucketInfoCache[bucket], nil
}
func (bucketManage *BucketManage) delBucketInfoAndBucketData(bucketInfoRequest *BucketInfoRequest, bucketInfo *meta.BucketInfo) error {
	bucketInfoResponse := &BucketInfoResponse{}
	if err := bucketManage.api.RmAllFileFromPath(bucketInfo.RealDirName); err != nil {
		bucketInfoResponse.Err = nil
		return err
	}
	bucketManage.conn.Del(context.Background(), bucketInfoRequest.Info.Name)
	bucketInfoRequest.Done <- bucketInfoResponse
	return nil
}
func CheckBucketExist( conn *redis.Conn,bucketName string) bool {
	//defer log.Errorln("checkBucketExist err:",err)
	ret, err := conn.Exists(context.Background(), bucketName).Result()
	if err != nil || ret > 0 {
		log.Errorln("the bucket:%v is exists", bucketName)
		return true
	}
	return false
}
func FetchBucketInfo(conn *redis.Conn,bucket string) (*meta.BucketInfo, error) {
	binstr, err := conn.Get(context.Background(), bucket).Result()
	if err != nil {
		return nil, err
	}
	bucketInfo := &meta.BucketInfo{}
	if err := bson.Unmarshal([]byte(binstr), bucketInfo); err != nil {
		return nil, err
	}
	return bucketInfo,nil
}
func CheckBucketStatus(bucketInfo *meta.BucketInfo) error {
	if bucketInfo.Status==BucketInActiveStatus {
		return errors.New(fmt.Sprintf("%s is inactive",bucketInfo.Name))
	}
	if bucketInfo.UsageInfo.ObjectsLimitCount> 0 && bucketInfo.UsageInfo.ObjectsCurrentCount > bucketInfo.UsageInfo.ObjectsLimitCount {
		return errors.New(fmt.Sprintf("%d over %d objects",bucketInfo.UsageInfo.ObjectsCurrentCount,bucketInfo.UsageInfo.ObjectsLimitCount ))
	}
	if bucketInfo.UsageInfo.CapacityLimitSize>0 &&bucketInfo.UsageInfo.CapacityCurrentSize > bucketInfo.UsageInfo.CapacityLimitSize {
		return errors.New(fmt.Sprintf("%d over %d bytes",bucketInfo.UsageInfo.CapacityCurrentSize,bucketInfo.UsageInfo.CapacityLimitSize ))
	}
	return nil
}
