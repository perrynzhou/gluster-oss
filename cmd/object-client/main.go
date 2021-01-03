package main

import (
	"context"
	"flag"
	"fmt"
	"glusterfs-storage-gateway/conf"
	"glusterfs-storage-gateway/protocol/pb"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	objectInfoFile = "/tmp/object-client.json"
)

var (
	requestObjectType = flag.String("o", "c", "p-put,g-get,d-delete object-client")
	confFile          = flag.String("c", "./conf.yaml", "default conf is ./conf.yaml")
	requestObjectFile = flag.String("f", "./test", "default upload file is ./test")
	requestBucketName = flag.String("b", "test", "default bucket  is test")
)

type Client struct {
	glusterStorageGatewayClient pb.GlusterStorageGatewayClient
	timeout                     int
	conn                        *grpc.ClientConn
}

func NewClient(path string) (*Client, error) {
	config, err := conf.NewClientConf(path)
	if err != nil {
		log.Fatal("Load Config failed:", err)
		return nil, err
	}
	diaOpt := grpc.WithDefaultCallOptions()
	cnn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Conf.Addr, config.Conf.Port), grpc.WithInsecure(), diaOpt)
	if err != nil {
		log.Error("grpc.Dial Failed, err:", err)
		return nil, nil
	}
	glusterStorageGatewayClient := pb.NewGlusterStorageGatewayClient(cnn)
	return &Client{
		glusterStorageGatewayClient: glusterStorageGatewayClient,
		conn:                        cnn,
		timeout:                     config.Conf.TimeOut,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func PutObject(c *Client, bucketName, localFile string) (*pb.PutObjectResponse, error) {
	var fileInfo os.FileInfo
	fileInfo, err := os.Stat(localFile)
	if err != nil {
		return &pb.PutObjectResponse{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.timeout)*time.Second)
	defer cancel()

	clientObjectRequest := &pb.PutObjectRequest{
		BucketName:  bucketName,
		LocalFile:   localFile,
		ObjectName:  filepath.Base(localFile),
		ObjectsSize: fileInfo.Size(),
		ContentType: filepath.Ext(localFile),
	}
	resp, err := c.glusterStorageGatewayClient.PutObject(ctx, clientObjectRequest)
	if err != nil {
		log.Errorf("putObject err:%v", err)
		return nil, err
	}
	log.Info("resp:", resp)
	return resp, nil

}

func main() {
	flag.Parse()
	c, err := NewClient(*confFile)
	if err != nil {
		log.Error("NewClient:", err)
		return
	}
	defer c.Close()

	if *requestObjectType == "p" {

	} else if *requestObjectType == "d" {

	} else {
		resp, err := PutObject(c, *requestBucketName, *requestObjectFile)
		if err != nil {
			log.Println(resp)
			return
		}
	}

}
