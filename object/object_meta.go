package object

const (
	DeleteStatus = iota
	UsedStatus
)

type ObjectMeta struct {
	Name       string
	Len        uint64
	Pos        uint64
	Stat       uint8
	VolumeName string
	BucketName string
	BucketFileName   string
}

func NewObjectMeta(name, addr string, len, pos uint64) (*ObjectMeta, error) {
	return &ObjectMeta{
		Name:       name,
		VolumeName: addr,
		Stat:       UsedStatus,
		Len:        len,
		Pos:        pos,
	}, nil
}
