/*************************************************************************
    > File Name: fs_api.h
  > Author:perrynzhou 
  > Mail:perrynzhou@gmail.com 
  > Created Time: Wednesday, September 16, 2020 PM03:10:58
 ************************************************************************/

#ifndef _FS_API_H
#define _FS_API_H
#include <glusterfs/api/glfs.h>
#include <glusterfs/api/glfs-handles.h>
typedef struct fs_api_t
{
  glfs_t *fs;
} fs_api;
typedef struct fs_fd_t
{
  int lfd;
  glfs_fd_t *gfd;
} fs_fd;
fs_api *fs_api_init(char *volume, char *addr, int port);
int fs_api_open(fs_api *fapi, fs_fd *fd, const char *pathname, int flags);
int fs_api_creat(fs_api *fapi, fs_fd *fd, const char *pathname, int flags, mode_t mode);
int fs_api_stat(fs_api *fapi,const char *pathname,struct stat *st);
off_t fs_api_seek(fs_api *fapi,fs_fd *fd,off_t offset, int whence);
ssize_t fs_api_read(fs_api *fapi, fs_fd *fd, void *buf, size_t count);
ssize_t fs_api_write(fs_api *fapi, fs_fd *fd, void *buf, size_t count);
int fs_api_fallocate(fs_fd *fd, int mode, off_t offset, off_t len);
int fs_api_mkdir(fs_api *fapi,const char *path,mode_t mode);
int fs_api_rmfile(fs_api *fapi,const char *path);
int fs_api_rmdir(fs_api *fapi,const char *path);
void fs_api_close(fs_fd *fd);
void fs_api_deinit(fs_api *fapi);
#endif
