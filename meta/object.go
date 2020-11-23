package meta

import "time"


type ObjectInfo struct {
	Id uint64
	// Name of the bucket.
	Bucket string
    BlockId uint64
	StartPos uint64
	// Name of the object.
	Name string
	// Date and time when the object was last modified.
	ModTime time.Time
	CreatTime time.Time
	AccessTime time.Time
	// Total object size.
	Size uint64
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
