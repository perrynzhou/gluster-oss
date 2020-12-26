package meta

import (
	fs_api "glusterfs-storage-gateway/fs-api"
	"sync"
)

const (
	BlockActive = 0
	BlockInactive =1
)
type BlockFile struct {
	File *fs_api.FsFd
	Lock *sync.Mutex
	LockFlag bool
	Meta *BlockFileMeta
}
type BlockFileMeta struct {
	Index uint64
	TotalObjectSize uint32
	ActiveObjectSize  uint32
	InactiveObjectSize  uint32
	Status uint8
}

func NewBlockFile(index uint64,file *fs_api.FsFd) *BlockFile {
	blockInfo := &BlockFileMeta{
		Index:uint64(index),
		TotalObjectSize:uint32(0),
		ActiveObjectSize :uint32(0),
		InactiveObjectSize:uint32(0),
		Status:BlockActive,
	}
	return &BlockFile{
		File:file,
		Lock:&sync.Mutex{},
		Meta:blockInfo,
	}
}
