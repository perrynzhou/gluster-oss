package main

import (
	fs_api "glusterfs-storage-gateway/fs-api"
	"log"
	"os"
)

func ReadData() {
	fsApi, err := fs_api.NewFsApi("dht-vol", "10.211.55.3", 24007, true)
	if err != nil {
		log.Fatal(err)
	}
	fd, err := fsApi.Open("test.1", os.O_RDONLY)
	if err != nil {
		log.Println("fsOpen err:",err)
		fd,err =  fsApi.Creat("test.1", os.O_RDONLY,os.ModePerm)
		if err == nil {
			fsApi.Write(fd,[]byte("test my data"))
		}
		log.Println("fsCreat err:",err)
	}
	defer fsApi.Close(fd)
	log.Println("fsOpen err:",err)
	buf := make([]byte, 4096)
	_, err = fsApi.Read(fd, buf)
	log.Println("read buf:", string(buf))
	fsApi.Releae()
}
func main() {
	ReadData()
}
