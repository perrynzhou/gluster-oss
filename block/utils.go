package block

import (
	fs_api "gluster-gtw/fs-api"
	"os"
)

const (
	MaxBlockLength = (1024 * 1024 * 124)
)


func createBlockFile(api *fs_api.FsApi,path string) error {
	fd,err :=api.Creat(path,os.O_RDWR|os.O_APPEND,os.ModePerm)
	if err != nil {
		return nil
	}
	api.Close(fd)
	return nil
}
