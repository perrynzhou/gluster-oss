package manage

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"glusterfs-storage-gateway/manage/bucket"
	"glusterfs-storage-gateway/meta"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"errors"
)

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
	if bucketInfo.Status==bucket.BucketInActiveStatus {
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