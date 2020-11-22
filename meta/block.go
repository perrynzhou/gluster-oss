package meta

const (
	BlockMaxBytes = (1024 * 1024 * 64)
)
// blockname : bucket-id-uuid.bk
type BlockInfo struct {
	Bucket string
	Index  int64
	TotalSize   uint64
	FreeSize    uint64
	Objects map[string]*ObjectInfo
}

func NewBlockInfo(bucket string, index int64) (*BlockInfo, error) {
	return &BlockInfo{
		Bucket: bucket,
		Index:  index,
		TotalSize:   uint64(BlockMaxBytes),
		FreeSize:uint64(BlockMaxBytes),
		Objects:make(map[string]*ObjectInfo),
	}, nil
}
