package meta

import "time"

const (
	ActiveObjectStatus =0
	InactiveObjectStatus=1
)
type ObjectInfo struct {
	Key string
	// Name of the bucket-client.
	Bucket string
    BlockId int64
	StartPos int64
	// Name of the object-client.
	Name string
	// Date and time when the object-client was last modified.
	ModTime time.Time
	CreatTime time.Time
	AccessTime time.Time
	// Total object-client size.
	Size int64
	// A standard MIME type describing the format of the object-client.
	ContentType string
	// Date and time at which the object-client is no longer able to be cached
	Expires time.Time
	// User-Defined metadata
	UserDefined map[string]string
	// User-Defined object-client tags
	UserTags string
	Status uint8
}
