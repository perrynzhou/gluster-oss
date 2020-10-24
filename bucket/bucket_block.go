package bucket

import (
	"fmt"
	fs_api "fusion-storage-gateway/fs-api"
	"os"
)

type BucketBlock struct {
	api        *fs_api.FsApi
	bucketName string
	meta       *BucketBlockMeta
}

func NewBucketBlock(api *fs_api.FsApi, name string, index uint64) (*BucketBlock, error) {
	bucketBlockFilePath := fmt.Sprintf("/%s/%s.%d", name, name, index)
	if err := createBucketBlockFile(api, bucketBlockFilePath); err != nil {
		return nil, err
	}
	bucketBlockMeta := NewBucketBlockMeta(index, MaxBucketBlockLen, name)
	return &BucketBlock{
		api:        api,
		bucketName: name,
		meta:       bucketBlockMeta,
	}, nil

}
func (bucketBlock *BucketBlock) Delete(filepath string) (int64, error) {
	var err error
	bucketBlockFilePath := fmt.Sprintf("/%s/%s.%d", bucketBlock.bucketName, bucketBlock.bucketName)
	if err = bucketBlock.api.RmFile(bucketBlockFilePath); err != nil {
		return int64(-1), err
	}
	return int64(0), nil
}
func (bucketBlock *BucketBlock) Seek(index, pos uint64) (*fs_api.FsFd, error) {

	bucketBlockFile := fmt.Sprintf("/%s/%s.%d", bucketBlock.bucketName, bucketBlock.bucketName, index)
	fd, err := bucketBlock.api.Seek(bucketBlockFile, os.O_RDONLY, pos, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	return fd, nil
}
func (bucketBlock *BucketBlock) Read(fd *fs_api.FsFd, data []byte) (int64, error) {
	return bucketBlock.api.Read(fd, data)
}

func (bucketBlock *BucketBlock) Write(fd *fs_api.FsFd, data []byte) (int64, error) {
	return bucketBlock.api.Write(fd, data)
}
