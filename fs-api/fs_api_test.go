package fs_api

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"log"
)
func TestNewFsApi(t *testing.T) {
	//172.25.78.11:rep_ssd_vol
	fsApi,err := NewFsApi("rep_vol","172.25.78.11",24007,true)
	if err != nil {
		log.Fatal(err)
	}
	assert.NotNil(t,fsApi.api)
	assert.Nil(t,err)
	fsApi.Releae()
}
func TestOpen(t *testing.T) {
	fsApi,err := NewFsApi("rep_vol","172.25.78.11",24007,true)
	if err != nil {
		log.Fatal(err)
	}
	fd,err :=fsApi.Open("test.1024",os.O_RDONLY)
	assert.NotNil(t,err)
	assert.Nil(t,fd)

	fd,err =fsApi.Open("test.9",os.O_RDONLY)
	assert.NotNil(t,fd)
	assert.Nil(t,err)
	fsApi.Releae()
}
func TestCreat(t *testing.T) {
	fsApi,err := NewFsApi("rep_vol","172.25.78.11",24007,true)
	if err != nil {
		log.Fatal(err)
	}
	fd,err :=fsApi.Creat("test.1023",os.O_RDWR|os.O_APPEND,os.ModePerm)
	assert.NotNil(t,fd)
	assert.Nil(t,err)
	fsApi.Releae()
}
func TestRead(t *testing.T) {
	fsApi,err := NewFsApi("rep_vol","172.25.78.11",24007,true)
	if err != nil {
		log.Fatal(err)
	}
	fd,err :=fsApi.Open("test.1",os.O_RDONLY)
	assert.NotNil(t,fd)
	assert.Nil(t,err)
	buf := make([]byte,4096)
	size,err :=fsApi.Read(fd,buf)
    assert.NotEqual(t,int64(size),-1)
	assert.Nil(t,err)
	log.Println("read buf:",string(buf))
	fsApi.Releae()
}



