package meta

import "time"

const (
	ActiveObjectStatus =0
	InactiveObjectStatus=1
)
type ObjectInfo struct {
	Key string
	// Name of the bucket.
	Bucket string
    BlockId int64
	StartPos int64
	// Name of the object.
	Name string
	// Date and time when the object was last modified.
	ModTime time.Time
	CreatTime time.Time
	AccessTime time.Time
	// Total object size.
	Size int64
	// A standard MIME type describing the format of the object.
	ContentType string
	// Date and time at which the object is no longer able to be cached
	Expires time.Time
	// User-Defined metadata
	UserDefined map[string]string
	// User-Defined object tags
	UserTags string
	Status uint8
}
