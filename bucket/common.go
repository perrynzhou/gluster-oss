package bucket

import (
	fs_api "fusion-storage-gateway/fs-api"
	"os"
)

const (
	MaxBucketBlockLen = (1024 * 1024 * 124)
)

const (
	BucketUsedSatus = 0
	BucketDelStatus = 1
)
func checkBucketExist(api *fs_api.FsApi, bucketPath string) error {
	if err := api.Stat(bucketPath); err != nil {
		return err
	}
	return nil
}
func createBucketBlockFile(api *fs_api.FsApi,path string) error {
	fd,err :=api.Creat(path,os.O_RDWR|os.O_APPEND,os.ModePerm)
	if err != nil {
		return nil
	}
	api.Close(fd)
	return nil
}
