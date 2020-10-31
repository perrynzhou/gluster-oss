package meta

type BucketUsageInfo struct {
	CapacityLimitSize   uint64
	ObjectsLimitCount   uint64 `json:"objectsLimitCount"`
	ObjectsCurrentCount uint64 `json:"objectsCurrentCount"`
}

// BucketInfo - represents bucket metadata.
type BucketInfo struct {
	// Name of the bucket.
	Name        string
	RealDirName string
	UsageInfo   *BucketUsageInfo
	Status      uint8
}
