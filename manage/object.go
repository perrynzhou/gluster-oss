package manage

import (
	"encoding/json"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
)

type Object struct {
	BlockFile  *meta.BlockFile
	TotalBytes uint64
	bucket     *Bucket
	Meta       *meta.ObjectInfo
}

func NewObject(api *fs_api.FsApi, blockFile *meta.BlockFile, bucket *Bucket, objMeta *meta.ObjectInfo) (*Object, error) {
	return &Object{
		BlockFile:  blockFile,
		TotalBytes: objMeta.Size,
		bucket:     bucket,
		Meta:       objMeta,
	}, nil
}
func (obj *Object) Write(api *fs_api.FsApi, data []byte) (int64, error) {
	wbytes, err := api.Write(obj.BlockFile.File, data)
	if err != nil {
		return wbytes, err
	}
	obj.BlockFile.AddBytes(uint64(wbytes))
	if obj.TotalBytes == obj.Meta.Size {
		obj.BlockFile.ModifyObjectCount(true)
		b, err := json.Marshal(obj.Meta)
		if err != nil {
			return -1, err
		}
		objValue := fmt.Sprintf("%s\n", string(b))
		api.Write(obj.bucket.ObjectMetaFile, []byte(objValue))
		if obj.BlockFile.Meta.TotalBytes >= meta.MaxBlockFileBytes {
			obj.BlockFile.ModifyStatusToInactive()
			obj.bucket.AllocBlockFile()
		}
	}
	return wbytes, nil
}

func (obj *Object) Read(api *fs_api.FsApi, data []byte) (int64, error) {
	return -1, nil
}
