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
	MaxStorageBytes     int64 `json:"LimitBytesMaxStorageBytes"`
	CurrentStorageBytes int64 `json:"LimitBytesCurrentStorageBytes"`
	MaxObjectCount      int64 `json:"objMaxObjectCount"`
	CurrentObjectCount  int64 `json:"objCurrentObjectCount"`
}

func NewBucketInfo(limitBytes, MaxObjectCount int64, bucketName, refDirName string) *BucketInfo {
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
