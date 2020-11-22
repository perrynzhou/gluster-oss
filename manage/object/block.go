package object

import (
	fs_api "glusterfs-storage-gateway/fs-api"
	"sync"
)

const (
	BlockKey = ""
)
type BlockFd struct {
	Fd *fs_api.FsFd
	Index  uint64
	Lock sync.Mutex
}
type BlockInfo struct {
	Index uint64
	Size  uint32
	Free  uint32
	DelObject  uint32
	DelSize   uint32
	Status uint8
}
