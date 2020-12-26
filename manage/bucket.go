package manage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
	"os"
	"strconv"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

const (
	BucketBlockIndexFileSubffix  = "block.index"
	BucketObjectIndexFileSubffix = "object.index"
	BucketDefauleBlockCount      = 32
)

type Bucket struct {
	MaxIndex           uint64
	BlockCount         int
	Meta               *meta.BucketInfo
	BlockFiles         map[uint64]*meta.BlockFile
	ObjectMetaFile     *fs_api.FsFd
	BlockIndexMetaFile *fs_api.FsFd
}

func (bucket *Bucket) Load(fsApi *fs_api.FsApi) *Bucket {
	var err error
	defer func(err error) {
		if err != nil {
			log.Error(err)
		}
	}(err)
	bucket.BlockCount = BucketDefauleBlockCount
	blockIndexFile, err := fsApi.Open(fmt.Sprintf("/%s/%s.%s", bucket.Meta.Name, bucket.Meta.Name, BucketBlockIndexFileSubffix), os.O_RDWR|os.O_APPEND)
	buf := make([]byte, 4096)
	fsApi.Read(blockIndexFile, buf)
	bucket.MaxIndex, err = strconv.ParseUint(string(buf), 10, 64)

	for i := uint64(0); i < bucket.MaxIndex; i++ {
		blockFilePath := fmt.Sprintf("/%s/block/%s.block.%d", bucket.Meta.Name, bucket.Meta.Name, i)
		blockFile, err := fsApi.Open(blockFilePath, os.O_RDWR|os.O_APPEND)
		if err != nil {
			return nil
		}
		if file := meta.NewBlockFile(uint64(i), blockFile); file != nil {
			file.LockFlag = false
			bucket.BlockFiles[uint64(i)] = file
		}
	}
	// create object meta file
	objMetaFile, err := fsApi.Open(fmt.Sprintf("/%s/%s.%s", bucket.Meta.Name, bucket.Meta.Name, BucketObjectIndexFileSubffix), os.O_RDWR|os.O_APPEND)
	if err != nil {
		return nil
	}
	data := make([]byte, 1024*1024*256)
	fsApi.Read(objMetaFile, data)
	reader := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		buketMeta := scanner.Text()
		buketMeta = buketMeta[0 : len(buketMeta)-1]
		meta := meta.BucketInfo{}
		if err := json.Unmarshal([]byte(buketMeta), &meta); err != nil {
			return nil
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return nil
	}
	bucket.ObjectMetaFile = objMetaFile
	bucket.BlockIndexMetaFile = blockIndexFile
	return bucket
}

func NewBucket(fsApi *fs_api.FsApi, cap, limit uint64, bucketName, refDirName string) *Bucket {
	var err error
	defer func(err error) {
		if err != nil {
			log.Error(err)
		}
	}(err)
	bucket := &Bucket{
		BlockCount: BucketDefauleBlockCount,
		Meta:       meta.NewBucketInfo(cap, limit, bucketName, refDirName),
		BlockFiles: make(map[uint64]*meta.BlockFile),
	}
	// create bucket path
	err = fsApi.Mkdir(fmt.Sprintf("/%s", bucketName), os.ModePerm)
	if err != nil {
		return nil
	}

	// create bucket block path
	err = fsApi.Mkdir(fmt.Sprintf("/%s/block", bucketName), os.ModePerm)
	if err != nil {
		return nil
	}
	// create block index
	blockIndexFile, err := fsApi.Creat(fmt.Sprintf("/%s/%s.%s", bucketName, bucketName, BucketBlockIndexFileSubffix), os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil
	}
	// fsApi.Close(blockIndexFile)

	// create object meta file
	objMetaFile, err := fsApi.Creat(fmt.Sprintf("/%s/%s.%s", bucketName, bucketName, BucketObjectIndexFileSubffix), os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil
	}
	// fsApi.Close(objIndexFile)
	bucket.ObjectMetaFile = objMetaFile
	bucket.BlockIndexMetaFile = blockIndexFile
	//
	for i := 0; i < bucket.BlockCount; i++ {
		blockFilePath := fmt.Sprintf("/%s/block/%s.block.%d", bucketName, bucketName, i)
		blockFile, err := fsApi.Creat(blockFilePath, os.O_RDWR|os.O_APPEND, os.ModePerm)
		if err != nil {
			return nil
		}
		if file := meta.NewBlockFile(uint64(i), blockFile); file != nil {
			file.LockFlag = false
			bucket.BlockFiles[uint64(i)] = file
		}

	}
	index := []byte(fmt.Sprintf("%d\n", bucket.BlockCount-1))
	fsApi.Write(bucket.BlockIndexMetaFile, index)
	return bucket
}
func (bucket *Bucket) IsPermit() error {
	if bucket.Meta.CurrentCount >= bucket.Meta.LimitCount {
		return errors.New(fmt.Sprintf("current count %d over limit count %d", bucket.Meta.CurrentCount, bucket.Meta.LimitCount))
	}
	if bucket.Meta.CurrentSize >= bucket.Meta.LimitSize {
		return errors.New(fmt.Sprintf("current size %d over limit size %d", bucket.Meta.CurrentSize, bucket.Meta.LimitSize))
	}
	if bucket.Meta.Status == BucketInActiveStatus {
		return errors.New(fmt.Sprintf("bucket is invalid", bucket.Meta.Name))
	}
	return nil
}
func (bucket *Bucket) ModifyObjectMeta(api *fs_api.FsApi, obj *meta.ObjectInfo) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	objStr := fmt.Sprintf("%s\n", string(b))
	api.Write(bucket.ObjectMetaFile, []byte(objStr))
	return nil
}
func (bucket *Bucket) AllocStorage() *meta.BlockFile {
	for _, blockFile := range bucket.BlockFiles {
		if !blockFile.LockFlag {
			return blockFile
		}
	}
	return nil
}
func StoreBucket(api *fs_api.FsApi, fd *fs_api.FsFd, bucket *Bucket) error {
	b, err := json.Marshal(bucket)
	if err != nil {
		return err
	}
	value := fmt.Sprintf("%s\n", string(b))
	api.Write(fd, []byte(value))
	return nil

}
func (bucket *Bucket) SummarySize(size uint64) {
	atomic.AddUint64(&bucket.Meta.CurrentSize, size)
}
func (bucket *Bucket) SummaryCount() {
	atomic.AddUint64(&bucket.Meta.CurrentCount, 1)
}
func ReleaseBucket(bucket *Bucket) {
	bucket.Meta.Status = meta.InactiveBucket
}
