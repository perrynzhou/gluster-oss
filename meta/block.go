package meta

import (
	"encoding/json"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"os"
	"sync"
	"sync/atomic"
)

const (
	BlockActive   = 0
	BlockInactive = 1
)
const (
	BlockFileLock   = true
	BlockFileUnlock = false
)
const (
	MaxBlockFileBytes = (1024*1024*256)
)

type BlockFile struct {
	File   *fs_api.FsFd
	Lock   *sync.Mutex
	IsLock bool
	MetaFile  *fs_api.FsFd
	IndexFile  *fs_api.FsFd
	Meta     *BlockFileInfo
}
type BlockFileInfo struct {
	BucketName   string
	BucketRefDir   string
	Index             int64
	TotalBytes        int64
	TotalObjectCount  int64
	ActiveObjectCount int64
	Status            uint8
}


func NewBlockFile(api *fs_api.FsApi,indexFile,metaFile *fs_api.FsFd,index int64,bucketName,bucketDirName string,isLoad bool) (*BlockFile,error) {
	var blockFile *fs_api.FsFd
	var err error
	blockFilePath := fmt.Sprintf("/%s/block/%s.block.%d", bucketDirName, bucketDirName, index)
	if isLoad {
		blockFile, err = api.Open(blockFilePath, os.O_RDWR|os.O_APPEND)
	}else{
		blockFile, err = api.Creat(blockFilePath, os.O_RDWR|os.O_APPEND, os.ModePerm)

	}
	if err != nil {
		return nil, err
	}
	blockInfo := &BlockFileInfo{
		Index:             int64(index),
		TotalObjectCount:  int64(0),
		ActiveObjectCount: int64(0),
		BucketName:bucketName,
		BucketRefDir  :bucketDirName,
		Status:            BlockActive,
	}
	return &BlockFile{
		MetaFile:metaFile,
		IndexFile:indexFile,
		File: blockFile,
		Lock: &sync.Mutex{},
		Meta: blockInfo,
	},nil
}

func (blockFile *BlockFile) ModifyObjectCount(opType bool) {
	value := int64(1)
	if !opType {
		value = -1
	}
	atomic.AddInt64(&blockFile.Meta.ActiveObjectCount, value)
	atomic.AddInt64(&blockFile.Meta.TotalObjectCount, value)
}
func (blockFile *BlockFile) AddBytes(n int64) {
	atomic.AddInt64(&blockFile.Meta.TotalBytes, n)
}
func (blockFile *BlockFile) ModifyStatusToInactive() {
	blockFile.Meta.Status =BlockInactive
}
func (blockFile *BlockFile) StoreMeta(api *fs_api.FsApi) error {
	b, err := json.Marshal(blockFile.Meta)
	if err != nil {
		return err
	}
	value := fmt.Sprintf("%s\n", string(b))
	api.Write(blockFile.MetaFile, []byte(value))
	return nil

}