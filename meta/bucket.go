package meta

type BucketUsageInfo struct {
	CapacityLimitSize   uint64
	ObjectsLimitCount   uint64 `json:"objectsLimitCount"`
	ObjectsCurrentCount uint64 `json:"objectsCurrentCount"`
	//conatains all blocks for this bucket
	Blocks              map[int]*BlockInfo `json:"blockInfo"`
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
