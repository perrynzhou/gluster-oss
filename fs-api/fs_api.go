package fs_api

//#cgo CFLAGS: -I/usr/include/glusterfs -I./
//#cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu/ -lgfapi
/*
#define _GNU_SOURCE
#include "fs_api.h"
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <glusterfs/api/glfs.h>
#include <glusterfs/api/glfs-handles.h>
*/
import (
	"C"
)
import (
	"errors"
	"fmt"
	"os"
	"unsafe"
)

type FsApi struct {
	api     *C.fs_api
	address string
}
type FsFd struct {
	Fd *C.fs_fd
}

//create glusterfs handler base on fstype
func NewFsApi(volume, addr string, port int, fstype bool) (*FsApi, error) {
	var fsApi *FsApi
	var err error
	if fstype {
		cvolume := C.CString(volume)
		caddr := C.CString(addr)
		defer C.free(unsafe.Pointer(cvolume))
		defer C.free(unsafe.Pointer(caddr))
		api := C.fs_api_init(cvolume, caddr, C.int(port))
		if api == nil {
			err = errors.New("init glfs failed")
		}
		fsApi = &FsApi{
			address: fmt.Sprintf("%s:%d:%s", addr, port, volume),
			api:     api,
		}
	}
	return fsApi, err
}

//check file state
func (fsApi *FsApi) Stat(name string) error {
	var st C.struct_stat
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret, err := C.fs_api_stat(fsApi.api, cname, &st)
	if int(ret) != 0 || err != nil {
		return err
	}
	return nil
}
func (fsApi *FsApi) RmFile(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	ret, err := C.fs_api_rmfile(fsApi.api, cpath)
	if int(ret) != 0 || err != nil {
		return err
	}
	return nil
}
func (fsApi *FsApi) RmDir(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	ret, err := C.fs_api_rmdir(fsApi.api, cpath)
	if int(ret) != 0 || err != nil {
		return err
	}
	return nil
}
func (fsApi *FsApi) Open(filename string, flags int) (*FsFd, error) {
	var err error

	fd := &FsFd{}
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	if _, err := C.fs_api_open(fsApi.api, fd.Fd.gfd, cfilename, C.int(flags)); err != nil {
		return nil, err
	}
	return fd, err
}
func (fsApi *FsApi) Creat(filename string, flags int, mode os.FileMode) (*FsFd, error) {
	fd := &FsFd{}
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	ret, err := C.fs_api_creat(fsApi.api, fd.Fd.gfd, cfilename, C.int(flags), C.mode_t(mode))
	if int(ret) != 0 {
		return nil, err
	}
	return fd, nil
}

func (fsApi *FsApi) Mkdir(path string, mode os.FileMode) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	ret, err := C.fs_api_mkdir(fsApi.api, cpath, C.mode_t(mode))
	if int(ret) != 0 {
		return err
	}
	return nil
}

func (fsApi *FsApi) Seek(filename string, flags int, offset uint64, whence int) (*FsFd,error) {
	coffset := C.off_t(offset)
	fd, err := fsApi.Open(filename, flags)
	if fd == nil {
		return nil,err
	}
	ret, err := C.fs_api_seek(fsApi.api, fd, coffset, C.int(whence))
	if int(ret) != 0 {
		C.fs_api_close(fd)
		return nil,err
	}
	return fd,nil
}
func (fsApi *FsApi) Write(fd *FsFd, data []byte) (int64, error) {

	sz, err := C.fs_api_write(fsApi.api, fd.Fd.gfd, unsafe.Pointer(&data[0]), (C.size_t)(len(data)))
	if err != nil {
		return int64(-1), err
	}
	return int64(sz), nil
}

func (fsApi *FsApi) Read(fd *FsFd, data []byte) (int64, error) {
	sz, err := C.fs_api_read(fsApi.api, fd.Fd.gfd, unsafe.Pointer(&data[0]), (C.size_t)(len(data)))
	if err != nil {
		return int64(-1), err
	}
	return int64(sz), nil
}
func (fsApi *FsApi) Close(fd *FsFd) error {
	var err error
	if fd != nil {
		_, err = C.fs_api_close(fd.Fd.gfd)
	}
	return err
}

func (api *FsApi) Fallocate(path string, flag int, size uint64) (*FsFd, error) {
	cfd, err := api.Open(path, flag)
	if err != nil {
		return nil, err
	}
	ret, err := C.glfs_fallocate(cfd, C.FALLOC_FL_KEEP_SIZE, 0, C.uint64_t(size))
	if int(ret) != 0 {
		_, err = C.glfs_close(cfd)
		return nil, err
	}
	return &FsFd{
		Fd: cfd,
	}, nil
}

func (fsApi *FsApi) Releae() {
	if fsApi != nil {
		C.fs_api_deinit(fsApi.api)
	}
}
