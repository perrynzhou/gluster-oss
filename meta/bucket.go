package meta


const (
	ActiveBucket = 0
	InactiveBucket=1
)

// BucketInfo - represents bucket metadata.
type BucketInfo struct {
	// Name of the bucket.
	Name        string
	RealDirName string
	Status      uint8
	LimitSize   uint64 `json:"capacityLimitSize"`
	CurrentSize   uint64 `json:"capacityCurrentSize"`
	LimitCount   uint64 `json:"objLimitCount"`
	CurrentCount uint64 `json:"objCurrentCount"`
}


func NewBucketInfo(limitBytes,limitCount uint64,bucketName,refDirName string) *BucketInfo {
	return &BucketInfo{
		Name:bucketName,
		RealDirName:refDirName,
		Status:ActiveBucket,
			LimitSize:limitBytes,
			CurrentSize:0,
			LimitCount:limitCount,
			CurrentCount:0,
	}
}
