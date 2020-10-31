package bucket

import (
	"context"
	"fmt"
	fs_api "gluster-storage-gateway/fs-api"
	"gluster-storage-gateway/meta"
	"os"

	"github.com/go-redis/redis/v8"
	bson "go.mongodb.org/mongo-driver/bson"
)


const (
	createBucketDirType= 0
	deleteBucketDirType= 1
)

func checkBucketExist(conn redis.Conn, bucketName string) error {
	if _, err := conn.Get(context.Background(), bucketName).Result(); err != nil {
		return err
	}
	return nil
}
func HandleBucketDir(api *fs_api.FsApi, bucketDirName string,bucketDirType int) error {
	var err error
	bucketDir := fmt.Sprintf("/%s", bucketDirName)
	switch(bucketDirType) {
	case createBucketDirType:
		err = api.Mkdir(bucketDir, os.ModePerm)
		break;
	case deleteBucketDirType:
		api.RmDir(bucketDir)
		break;
	}
	return err
}

func storeBucketInfo(conn redis.Conn, bucketInfo *meta.BucketInfo) (string, error) {
	var result string
	b, err := bson.Marshal(bucketInfo)
	if err != nil {
		return result, err
	}
	return conn.Set(context.Background(), bucketInfo.Name, b, -1).Result()
}
func fetchBucketInfo(conn redis.Conn, bucetInfoCache map[string]*meta.BucketInfo, bucket string) (*meta.BucketInfo, error) {
	if bucetInfoCache[bucket] == nil {
		binstr, err := conn.Get(context.Background(), bucket).Result()
		if err != nil {
			return nil, err
		}
		bucketInfo := &meta.BucketInfo{}
		if err := bson.Unmarshal([]byte(binstr), bucketInfo); err != nil {
			return nil, err
		}
		bucetInfoCache[bucket] = bucketInfo
	}
	return bucetInfoCache[bucket], nil
}
func delBucketInfoAndBucketData(api *fs_api.FsApi,bucketInfoRequest *BucketInfoRequest,bucketDir string) error{
	bucketInfoResponse :=&BucketInfoResponse{}
  if err := api.RmAllFileFromPath(bucketDir);err != nil {
	  bucketInfoResponse.Err = nil
	  return err
  }
  bucketInfoRequest.Done <-bucketInfoResponse
  return nil
}
