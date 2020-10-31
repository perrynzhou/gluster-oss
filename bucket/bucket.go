package bucket

import (
	"github.com/google/uuid"
	meta "gluster-storage-gateway/meta"
	"gluster-storage-gateway/protocol/pb"
)

const (
	CreateBucketType = 1
	DeleteBucketType = 2
	UpdateBucketType = 3
)
const (
	BucketActiveStatus=1
	BucketInActiveStatus=2

)

type BucketInfoRequest struct {
	RequestType uint8
	Info *meta.BucketInfo
	Done chan *BucketInfoResponse
}
type BucketInfoResponse struct {
	Err  error
}

func NewCreateBucketInfoRequest(request *pb.CreateBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
			UsageInfo: &meta.BucketUsageInfo{
				ObjectsLimitCount:   request.ObjectsLimit,
				ObjectsCurrentCount: uint64(0),
			},
			Status:BucketActiveStatus,
			RealDirName: uuid.New().String(),
		},
		Done: make(chan *BucketInfoResponse),
		RequestType:CreateBucketType,
	}
}
func NewDeleteBucketInfoRequest(request *pb.CreateBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
		},
		Done: make(chan *BucketInfoResponse),
		RequestType:DeleteBucketType,
	}
}
func NewUpdateBucketInfoRequest(request *pb.CreateBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
			UsageInfo: &meta.BucketUsageInfo{
				ObjectsLimitCount:   request.ObjectsLimit,
				ObjectsCurrentCount: uint64(0),
			},
			Status:BucketActiveStatus,
		},
		Done: make(chan *BucketInfoResponse),
		RequestType:CreateBucketType,
	}
}
/*
func NewBuckManage(api *fs_api.FsApi)  (*BlockManage,error) {
	if api==nil {
		return nil,errors.New("api is nil")
	}
	return &BlockManage {
		api:api,
		ch:make(chan *meta.BlockInfo),
	},nil
}
func (mgr *BlockManage)Run() {
	for {
		select {
		case
		}
	}
}
func NewBlock(api *fs_api.FsApi, name string, index uint64) (*Block, error) {
	blockFilePath := fmt.Sprintf("/%s/%s.%d", name, name, index)
	if err := block.createBlockFile(api, blockFilePath); err != nil {
		return nil, err
	}
	blockMeta := meta.NewBlockMeta(index, block.MaxBlockLength, name)
	return &Block{
		api:        api,
		bucketName: name,
		meta:       blockMeta,
	}, nil

}
func (block *Block) Delete(filepath string) (int64, error) {
	var err error
	blockFilePath := fmt.Sprintf("/%s/%s.%d", block.bucketName, block.bucketName)
	if err = block.api.RmFile(blockFilePath); err != nil {
		return int64(-1), err
	}
	return int64(0), nil
}
func (block *Block) Seek(index, pos uint64) (*fs_api.FsFd, error) {

	blockFile := fmt.Sprintf("/%s/%s.%d", block.bucketName, block.bucketName, index)
	fd, err := block.api.Seek(blockFile, os.O_RDONLY, pos, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	return fd, nil
}
func (block *Block) Read(fd *fs_api.FsFd, data []byte) (int64, error) {
	return block.api.Read(fd, data)
}

func (block *Block) Write(fd *fs_api.FsFd, data []byte) (int64, error) {
	return block.api.Write(fd, data)
}


*/
