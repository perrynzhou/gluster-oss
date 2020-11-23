package object

import (
	"context"
	"encoding/json"
	"fmt"
	"glusterfs-storage-gateway/meta"
	"os"
)

func (objectManage *ObjectManage) initBlockMetaAndIndexFile(bucketInfo *meta.BucketInfo) error {
	var err error
	if objectManage.BlockMetaFile[bucketInfo.Name] == nil {
		objectManage.BlockMetaFile[bucketInfo.Name], err = objectManage.api.Creat(fmt.Sprintf("/%s/%s", bucketInfo.RealDirName, blockMetaFileName), os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}
		blockIndexFileName := fmt.Sprintf("/%s/%s", bucketInfo.RealDirName, blockIndexFileName)
		objectManage.BlockIndexFile[bucketInfo.Name], err = objectManage.api.Creat(blockIndexFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
func (objectManage *ObjectManage) initBlockCache(bucketInfo *meta.BucketInfo,startIndex int) error {
	var err error
	if objectManage.BlockCache[bucketInfo.Name] == nil {
		blockIndex := fmt.Sprintf("%d\n", 0)
		if _, err = objectManage.api.Write(objectManage.BlockIndexFile[bucketInfo.Name], []byte(blockIndex)); err != nil {
			return err
		}
		objectManage.BlockCache[bucketInfo.Name] = make([]*BlockFd, objectManage.writeConcurrent)
		blockIndexKey := fmt.Sprintf("%s.bid", bucketInfo.Name)
		objectManage.conn.IncrBy(context.Background(), blockIndexKey, int64(objectManage.writeConcurrent))
		for i := startIndex; i < (startIndex+objectManage.writeConcurrent); i++ {
			fd,err := objectManage.api.Creat(fmt.Sprintf("/%s/%d.block", bucketInfo.RealDirName, i), os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
			if err != nil {
				return err
			}
			bfd:= &BlockFd{
				Index:uint64(i),
				Fd:fd,
				Info:&BlockInfo{
					Index :uint64(i),
					Size :0,
					DelObject:0,
					DelSize:0,
					Status:ObjectActiveStatus,
				},
			}
			objectManage.BlockCache[bucketInfo.Name][i] = bfd
			b,err := json.Marshal(bfd)
			if err != nil {
				return err
			}
			line := fmt.Sprintf("%d\t%s\t%s\n",bfd.Index,bucketInfo.Name, string(b))
			if _,err =objectManage.api.Write(objectManage.BlockMetaFile[bucketInfo.Name],[]byte(line));err != nil {
				return err
			}
			objectManage.conn.Set(context.Background(), fmt.Sprintf("%s.%d.b",bucketInfo.Name,bfd.Index), string(b),-1)
		}
	}
	return nil
}
