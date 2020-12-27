package manage

import (
	"bufio"
	fs_api "glusterfs-storage-gateway/fs-api"
	"io"
	"os"
	"path/filepath"
)

func PutData(localFilePath string, api *fs_api.FsApi, obj *Object) (int64, error) {
	localFile, err := os.OpenFile(localFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return -1, err
	}
	obj.Meta.ContentType = filepath.Ext(localFilePath)
	// Reset the read pointer if necessary.
	localFile.Seek(0, 0)
	fileInfo, err := localFile.Stat()
	obj.Meta.Size = fileInfo.Size()
	nBytes := int64(0)
	r := bufio.NewReader(localFile)
	buf := make([]byte, 0, 4*1024*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
		}
		nBytes += int64(len(buf))
		obj.Write(api, buf)
		if err != nil && err != io.EOF {
			return -1, err
		}
	}
	return nBytes, nil

}
