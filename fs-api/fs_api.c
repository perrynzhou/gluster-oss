/*************************************************************************
    > File Name: fs_api.c
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Wednesday, September 16, 2020 PM03:11:03
 ************************************************************************/

#include "fs_api.h"
#include <fcntl.h>
#include <glusterfs/api/glfs-handles.h>
#include <glusterfs/api/glfs.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <stdio.h>
fs_api *fs_api_init(char *volume, char *addr, int port)
{
  fs_api *fapi = calloc(1, sizeof(fs_api *));
  char buf[1024] = {'\0'};
  snprintf((char *)&buf, 128, "/tmp/gluster-storage-gateway.log.%d", getpid());
  glfs_t *fs = glfs_new(volume);
  glfs_set_volfile_server(fs, "tcp", addr, port);
  glfs_set_logging(fs, (char *)&buf, 7);
  glfs_init(fs);
  fapi->fs = fs;
  return fapi;
}
off_t fs_api_seek(fs_api *fapi, fs_fd *fd, off_t offset, int whence)
{
  if (fd->lfd == -1 || fd->gfd == NULL)
  {
    return -1;
  }
  if (fapi == NULL)
  {
    return lseek(fd->lfd, offset, whence);
  }
  return glfs_lseek(fd->gfd, offset, whence);
}
int fs_api_stat(fs_api *fapi, const char *pathname, struct stat *st)
{
  if (pathname == NULL)
  {
    return -1;
  }
  if (fapi == NULL)
  {
    return stat(pathname, st);
  }
  return glfs_stat(fapi->fs, pathname, st);
}
int fs_api_open(fs_api *fapi, fs_fd *fd, const char *pathname, int flags)
{
  fd->lfd = -1;
  fd->gfd = NULL;
  if (fapi == NULL)
  {
    fd->lfd = open(pathname, flags);
    return fd->lfd;
  }
  //not support create file
  //flags only O_RDWR
  fd->gfd = glfs_open(fapi->fs, pathname, flags);
  if (fd->gfd == NULL)
  {
    return -1;
  }
  return 0;
}
int fs_api_rm_file_from_path(fs_api *fapi, const char *path)
{
  char buffer[4096] = {'\0'};
  char dirent_buffer[512] = {'\0'};
  struct dirent *dt = NULL;
  glfs_fd_t *gfd = glfs_opendir(fapi->fs, path);
  while (glfs_readdir_r(gfd, (struct dirent *)dirent_buffer, &dt), dt)
  {
    size_t len = strlen(dt->d_name);
    snprintf((char *)&buffer, 4096, "/%s/%s", path, dt->d_name);
    if (dt->d_type == DT_REG)
    {
      glfs_unlink(fapi->fs, (char *)&buffer);
    }
  }
  return 0;
}
ssize_t fs_api_read(fs_api *fapi, fs_fd *fd, void *buf, size_t count)
{
  if (fapi == NULL && fd->lfd != -1)
  {
    return read(fd->lfd, buf, count);
  }
  return glfs_read(fd->gfd, buf, count, 0);
}
ssize_t fs_api_write(fs_api *fapi, fs_fd *fd, void *buf, size_t count)
{
  if (fapi == NULL && fd->lfd != -1)
  {
    return write(fd->lfd, buf, count);
  }
  return glfs_write(fd->gfd, buf, count, 0);
}
void fs_api_close(fs_fd *fd)
{
  if (fd != NULL && fd->gfd != NULL)
  {
    glfs_close(fd->gfd);
    return;
  }
  if (fd != NULL && fd->lfd != -1)
  {
    close(fd->lfd);
  }
}
int fs_api_creat(fs_api *fapi, fs_fd *fd, const char *pathname, int flags, mode_t mode)
{
  if (fd == NULL || pathname == NULL)
  {
    return -1;
  }
  if (fapi == NULL)
  {
    fd->lfd = creat(pathname, mode);
    return fd->lfd;
  }
  if ((fd->gfd = glfs_creat(fapi->fs, pathname, flags, mode)) == NULL)
  {
    return -1;
  }
  return 0;
}
int fs_api_fallocate(fs_fd *fd, int mode, off_t offset, off_t len)
{
  if (fd == NULL)
  {
    return -1;
  }
  if (fd->lfd != -1)
  {
    return fallocate(fd->lfd, mode, offset, len);
  }
  return glfs_fallocate(fd->gfd, mode, offset, len);
}
int fs_api_mkdir(fs_api *fapi, const char *path, mode_t mode)
{
  if (fapi != NULL)
  {
    return glfs_mkdir(fapi->fs, path, mode);
  }
  return mkdir(path, mode);
}
int fs_api_rmfile(fs_api *fapi, const char *path)
{
  if (fapi == NULL)
  {
    return remove(path);
  }
  return glfs_unlink(fapi->fs, path);
}
int fs_api_rmdir(fs_api *fapi, const char *path)
{
  if (fapi == NULL)
  {
    return rmdir(path);
  }
  return glfs_rmdir(fapi->fs, path);
}
void fs_api_deinit(fs_api *fapi)
{
  if (fapi != NULL)
  {
    glfs_fini(fapi->fs);
    free(fapi);
    fapi = NULL;
  }
}

#ifdef FS_API_TEST
int main(void)
{
  //172.25.78.19:/train_vol
  //172.25.78.11:rep_ssd_vol
  fs_api *fapi = fs_api_init("ssd_vol", "172.25.78.11", 24007);
  if (fapi == NULL)
  {
    fprintf(stdout, "inint failed\n");
    return -1;
  }

  fs_fd fd;
  char buf[64] = {'\0'};
  snprintf((char *)&buf, 64, "test.%d", rand() % 10);
  char *test_file = strdup((char *)&buf);
  fprintf(stdout, "fs_api_open:%d\n", fs_api_open(fapi, &fd, "test.data", O_RDWR));
  fs_api_close(&fd);
  fprintf(stdout, "fs_api_creat:%d\n", fs_api_creat(fapi, &fd, (char *)&buf, O_CREAT | O_RDWR | O_APPEND, 0644));
  size_t count = 0;
  for (int i = 0; i < 100; i++)
  {
    snprintf((char *)&buf, 64, "test data=%d\n", i);
    size_t blen = strlen((char *)&buf);
    count += blen;
    fprintf(stdout, "write ret:%ld\n", fs_api_write(fapi, &fd, (char *)&buf, blen));
  }
  fs_api_close(&fd);
  fprintf(stdout, "fs_api_open:%d\n", fs_api_open(fapi, &fd, test_file, O_RDWR));
  char rb[4096] = {'\0'};
  fprintf(stdout, "read ret:%ld\n", fs_api_read(fapi, &fd, (char *)&rb, 4096));
  fprintf(stdout, "buf=%s\n", (char *)&rb);

  fs_api_close(&fd);
  fprintf(stdout, "fs_api_open:%d\n", fs_api_creat(fapi, &fd, "test.data.fallocate2", O_CREAT | O_RDWR, 0644));
  int ret = fs_api_fallocate(&fd, FALLOC_FL_KEEP_SIZE, 1024 * 1024 * 1024, 1024 * 1024 * 1024);
  fprintf(stdout, "fs_api_fallocate:%d,err:%s\n", ret, strerror(ret));
  fs_api_close(&fd);
  free(test_file);
  fs_api_deinit(fapi);
  return 0;
}
#endif