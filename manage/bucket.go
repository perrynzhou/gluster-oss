package manage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
	"io"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

const (
	BucketBlockIndexFileSuffix  = "block.index"
	BucketObjectMetaFileSuffix = "object.meta"
	BucketBlockMetaFileSuffix = "block.meta"
	BucketDefauleBlockCount      = 32
)

type Bucket struct {
	MaxIndex           int64
	BlockCount         int
	Meta               *meta.BucketInfo
	BucketMetaFile       *fs_api.FsFd
	BlockFiles         map[int64]*meta.BlockFile
	ObjectMetaFile     *fs_api.FsFd
	BlockIndexFile *fs_api.FsFd
	BlockMetaFile        *fs_api.FsFd
	Locker       *sync.Mutex
}

func (bucket *Bucket) Load(fsApi *fs_api.FsApi) (*Bucket, error) {
	var err error
	var blockFile *meta.BlockFile
	bucket.BlockCount = BucketDefauleBlockCount
	bucketFile, err := fsApi.Open(bucketMetaFilePath, os.O_RDONLY)


	blockIndexFile, err := fsApi.Open(fmt.Sprintf("/%s/%s.%s", bucket.Meta.Name, bucket.Meta.Name, BucketBlockIndexFileSuffix), os.O_RDONLY)
	buf := make([]byte, 4096)
	fsApi.Read(blockIndexFile, buf)
	bucket.MaxIndex, err = strconv.ParseInt(string(buf), 10, 64)

	// load  object meta file
	objectMetaFile, err := fsApi.Open(fmt.Sprintf("/%s/%s.%s", bucket.Meta.Name, bucket.Meta.Name, BucketObjectMetaFileSuffix), os.O_RDONLY)
	if err != nil {
		return nil, err
	}
	// load  object meta file
	blockMetaFile, err := fsApi.Open(fmt.Sprintf("/%s/%s.%s", bucket.Meta.Name, bucket.Meta.Name, BucketBlockMetaFileSuffix), os.O_RDWR|os.O_APPEND)
	if err != nil {
		return nil, err
	}

	for i := int64(0); i <= bucket.MaxIndex; i++ {
		if blockFile, err = meta.NewBlockFile(fsApi, blockIndexFile,blockMetaFile,i,bucket.Meta.Name, bucket.Meta.RealDirName, true); err != nil {
			log.Errorln("load ", bucket.Meta.RealDirName, ",index %d", i, ",err:", err)
			continue
		}
		if blockFile.Meta.Status != meta.BlockInactive {
			blockFile.IsLock = false
			bucket.BlockFiles[i] = blockFile
		}

	}

	data := make([]byte, 1024*1024*256)
	fsApi.Read(objectMetaFile, data)
	buffer := bytes.NewBuffer(data)
	reader := bufio.NewReader(buffer)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		meta := meta.BucketInfo{}
		if err := json.Unmarshal([]byte(line), &meta); err != nil {
			return nil, err
		}
	}
	if err != io.EOF {
		log.Printf(" > Failed with error: %v\n", err)
		return nil, err
	}
	bucket.ObjectMetaFile = objectMetaFile
	bucket.BlockIndexFile = blockIndexFile
	bucket.BucketMetaFile =bucketFile
	//to do
	bucket.BlockMetaFile = blockMetaFile
	return bucket, nil
}

func NewBucket(fsApi *fs_api.FsApi, bucketFile *fs_api.FsFd,cap, limit int64, bucketName, refDirName string) (*Bucket, error) {
	var err error
	defer func(err error) {
		if err != nil {
			log.Error(err)
		}
	}(err)
	bucket := &Bucket{
		BlockCount: BucketDefauleBlockCount,
		Meta:       meta.NewBucketInfo(cap, limit, bucketName, refDirName),
		BlockFiles: make(map[int64]*meta.BlockFile),
		Locker:&sync.Mutex{},
		BucketMetaFile:bucketFile,
	}
	// create bucket path
	err = fsApi.Mkdir(fmt.Sprintf("/%s", bucket.Meta.RealDirName), os.ModePerm)
	if err != nil {
		return nil, err
	}

	// create bucket block path
	err = fsApi.Mkdir(fmt.Sprintf("/%s/block", bucket.Meta.RealDirName), os.ModePerm)
	if err != nil {
		return nil, err
	}
	// create block index
	blockIndexFile, err := fsApi.Creat(fmt.Sprintf("/%s/%s.%s", bucket.Meta.RealDirName, bucket.Meta.RealDirName, BucketBlockIndexFileSuffix), os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	// create block index
	blockMetaFile, err := fsApi.Creat(fmt.Sprintf("/%s/%s.%s", bucket.Meta.RealDirName, bucket.Meta.RealDirName, BucketBlockMetaFileSuffix), os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	// fsApi.Close(blockIndexFile)

	// create object meta file
	objectMetaFile, err := fsApi.Creat(fmt.Sprintf("/%s/%s.%s", bucket.Meta.RealDirName, bucket.Meta.RealDirName, BucketObjectMetaFileSuffix), os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	// fsApi.Close(objIndexFile)
	bucket.ObjectMetaFile = objectMetaFile
	bucket.BlockIndexFile = blockIndexFile
	bucket.BlockMetaFile = blockMetaFile

	//
	for i := 0; i < bucket.BlockCount; i++ {
		if file, err := meta.NewBlockFile(fsApi, blockIndexFile,blockMetaFile,int64(i), bucket.Meta.Name,bucket.Meta.RealDirName, false); err == nil {
			file.IsLock = meta.BlockFileUnlock
			file.StoreMeta(fsApi)
			bucket.BlockFiles[int64(i)] = file
		}
	}
	index := []byte(fmt.Sprintf("%d\n", bucket.BlockCount-1))
	fsApi.Write(bucket.BlockIndexFile, index)
	return bucket, nil
}
func (bucket *Bucket) checkStatus() error {
	if bucket.Meta.Status == meta.InactiveBucket {
		return errors.New(fmt.Sprintf("bucket  %s   inactive", bucket.Meta.Name))
	}
	return nil
}
func (bucket *Bucket) CheckLimit() error {
	if bucket.Meta.CurrentObjectCount >= bucket.Meta.MaxObjectCount {
		return errors.New(fmt.Sprintf("current count %d over limit count %d", bucket.Meta.CurrentObjectCount, bucket.Meta.MaxObjectCount))
	}
	if bucket.Meta.CurrentStorageBytes >= bucket.Meta.MaxStorageBytes {
		return errors.New(fmt.Sprintf("current size %d over limit size %d", bucket.Meta.CurrentStorageBytes, bucket.Meta.MaxStorageBytes))
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
func (bucket *Bucket) GetBlockFile() *meta.BlockFile {
	for _, blockFile := range bucket.BlockFiles {
		if  !blockFile.IsLock {
			return blockFile
		}
	}
	return nil
}
func (bucket *Bucket) StoreMeta(api *fs_api.FsApi, fd *fs_api.FsFd) error {
	b, err := json.Marshal(bucket.Meta)
	if err != nil {
		return err
	}
	value := fmt.Sprintf("%s\n", string(b))
	api.Write(bucket.BlockMetaFile, []byte(value))
	return nil

}
func (bucket *Bucket) AllocBlockFile() error {
	bucket.Locker.Lock()
	// create tmp file
	defer bucket.Locker.Unlock()
	return nil
}
func (bucket *Bucket) ModifyObjectBytes(size int64) {
		atomic.AddInt64(&bucket.Meta.CurrentStorageBytes, size)
}
func (bucket *Bucket) AddObjectCounter() {
	atomic.AddInt64(&bucket.Meta.CurrentObjectCount, 1)
}
func (bucket *Bucket) SubObjectCounter() {
	atomic.AddInt64(&bucket.Meta.CurrentObjectCount, -1)
}
func ReleaseBucket(bucket *Bucket) {
	bucket.Meta.Status = meta.InactiveBucket
}
