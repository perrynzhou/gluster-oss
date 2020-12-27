package meta

const (
	ActiveBucket   = 0
	InactiveBucket = 1
)

// BucketInfo - represents bucket metadata.
type BucketInfo struct {
	// Name of the bucket.
	Name                string
	RealDirName         string
	Status              uint8
	MaxStorageBytes     uint64 `json:"LimitBytesMaxStorageBytes"`
	CurrentStorageBytes uint64 `json:"LimitBytesCurrentStorageBytes"`
	MaxObjectCount      uint64 `json:"objMaxObjectCount"`
	CurrentObjectCount  uint64 `json:"objCurrentObjectCount"`
}

func NewBucketInfo(limitBytes, MaxObjectCount uint64, bucketName, refDirName string) *BucketInfo {
	return &BucketInfo{
		Name:                bucketName,
		RealDirName:         refDirName,
		Status:              ActiveBucket,
		MaxStorageBytes:     limitBytes,
		CurrentStorageBytes: 0,
		MaxObjectCount:      MaxObjectCount,
		CurrentObjectCount:  0,
	}
}
