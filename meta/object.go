package meta

import "time"

type ObjectInfo struct {
	// Name of the bucket.
	Bucket string

	// Name of the object.
	Name string

	// Date and time when the object was last modified.
	ModTime time.Time

	// Total object size.
	Size int64

	// Hex encoded unique entity tag of the object.
	ETag string

	// Version ID of this object.
	VersionID string

	// IsLatest indicates if this is the latest current version
	// latest can be true for delete marker or a version.
	IsLatest bool

	// A standard MIME type describing the format of the object.
	ContentType string

	// Date and time at which the object is no longer able to be cached
	Expires time.Time

	// User-Defined metadata
	UserDefined map[string]string

	// User-Defined object tags
	UserTags string

	// Date and time when the object was last accessed.
	AccTime time.Time

	// backendType indicates which backend filled this structure
	backendType string
}
