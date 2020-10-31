package meta

const (
	BlockMaxBytes = (1024*1024*64)
)
type BlockInfo struct {
	Bucket string
	Index  int64
	Id     string
	Size   int64
}

func NewBlockInfo(bucket string,index int64,id string) (*BlockInfo,error) {
	return  &BlockInfo{
		Bucket:bucket,
		Index:index,
		Id:id,
		Size:int64(0),
	},nil
}
