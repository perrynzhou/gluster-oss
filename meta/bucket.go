package meta

const (
	ActiveBucket   = 0
	InactiveBucket = 1
)

// BucketInfo - represents bucket-client metadata.
type BucketInfo struct {
	// Name of the bucket-client.
	Name                string
	RealDirName         string
	Status              uint8
	MaxStorageBytes     int64
	CurrentStorageBytes int64
	MaxObjectCount      int64
	CurrentObjectCount  int64
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
