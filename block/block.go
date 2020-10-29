package block

import (
	"fmt"
	fs_api "gluster-gtw/fs-api"
	"os"
)

type Block struct {
	api        *fs_api.FsApi
	bucketName string
	meta       *BlockMeta
}

func NewBlock(api *fs_api.FsApi, name string, index uint64) (*Block, error) {
	blockFilePath := fmt.Sprintf("/%s/%s.%d", name, name, index)
	if err := createBlockFile(api, blockFilePath); err != nil {
		return nil, err
	}
	blockMeta := NewBlockMeta(index, MaxBlockLength, name)
	return &Block{
		api:        api,
		bucketName: name,
		meta:       blockMeta,
	}, nil

}
func (block *Block) Delete(filepath string) (int64, error) {
	var err error
	blockFilePath := fmt.Sprintf("/%s/%s.%d", block.bucketName, block.bucketName)
	if err = block.api.RmFile(blockFilePath); err != nil {
		return int64(-1), err
	}
	return int64(0), nil
}
func (block *Block) Seek(index, pos uint64) (*fs_api.FsFd, error) {

	blockFile := fmt.Sprintf("/%s/%s.%d", block.bucketName, block.bucketName, index)
	fd, err := block.api.Seek(blockFile, os.O_RDONLY, pos, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	return fd, nil
}
func (block *Block) Read(fd *fs_api.FsFd, data []byte) (int64, error) {
	return block.api.Read(fd, data)
}

func (block *Block) Write(fd *fs_api.FsFd, data []byte) (int64, error) {
	return block.api.Write(fd, data)
}
