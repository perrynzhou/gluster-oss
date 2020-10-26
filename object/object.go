package object

import (
	"errors"
	fs_api "gluster-oss/fs-api"
	"github.com/google/uuid"
)


type Object struct {
    Key  string
    Meta *ObjectMeta
}

func NewObject(meta *ObjectMeta) (*Object,error) {
	if meta == nil {
		return nil ,errors.New("invald object length")
	}
	uid,err:=uuid.NewUUID()
	if err != nil {
		return nil,err
	}
	return &Object {
		Key:uid.String(),
		Meta:meta,
	},nil
}
func (obj *Object)Write(targetFile string,api *fs_api.FsApi,bytes []byte) (int64,error) {
	return -1,nil
}
func (obj *Object)Read() (int64,error) {
	return -1,nil
}




