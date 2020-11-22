package meta

type BucketUsageInfo struct {
	CapacityLimitSize   uint64 `json:"capacityLimitSize"`
	CapacityCurrentSize   uint64 `json:"capacityCurrentSize"`
	ObjectsLimitCount   uint64 `json:"objLimitCount"`
	ObjectsCurrentCount uint64 `json:"objCurrentCount"`
	//conatains all blocks for this bucket
	//conatinas all object for this bucket
	Objects             map[string]*ObjectInfo `json:"objectInfo"`
}

// BucketInfo - represents bucket metadata.
type BucketInfo struct {
	// Name of the bucket.
	Name        string
	RealDirName string
	UsageInfo   *BucketUsageInfo
	Status      uint8
}
