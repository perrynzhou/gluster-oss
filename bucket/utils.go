package bucket

import (
	fs_api "gluster-gtw/fs-api"
)

const (
	MaxBucketBlockLen = (1024 * 1024 * 124)
)

const (
	UsedSatus = 0
	DelStatus = 1
)
func checkBucketExist(api *fs_api.FsApi, bucketPath string) error {
	if err := api.Stat(bucketPath); err != nil {
		return err
	}
	return nil
}

