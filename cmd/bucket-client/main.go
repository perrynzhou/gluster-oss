package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"glusterfs-storage-gateway/conf"
	"glusterfs-storage-gateway/protocol/pb"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	bucketInfoFile = "/tmp/bucket-client.json"
)

var (
	requestBucketType = flag.String("o", "c", "c-create,u-upate,d-delete bucket-client")
	confFile          = flag.String("c", "./conf.yaml", "default conf is ./conf.yaml")
	count             = flag.Int("n", 1, "defaiult count is 1")
)

type ClientBucketRequest struct {
	Request []*pb.CreateBucketRequest
}
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

func CreateBucket(c *Client, n int) (*ClientBucketRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.timeout)*time.Second)
	defer cancel()

	ClientBucketRequest := &ClientBucketRequest{
		Request: make([]*pb.CreateBucketRequest, n),
	}
	for i := 0; i < n; i++ {
		ClientBucketRequest.Request[i] = &pb.CreateBucketRequest{
			Name:         fmt.Sprintf("bucket-client-%d", i),
			MaxStorageBytes:     100,
			MaxObjectCount: 100,
		}
		resp, err := c.glusterStorageGatewayClient.CreateBucket(ctx, ClientBucketRequest.Request[i])
		if err != nil {
			log.Errorf("CreateBucket err:%v", err)
			return nil, err
		}
		log.Info("resp:", resp)
	}
	b, err := json.Marshal(&ClientBucketRequest)
	if err != nil {
		log.Errorln("err:", err)
	}
	var file *os.File
	file, err = os.OpenFile(bucketInfoFile,os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Errorln("err:",err)
		return nil,err
	}
	file.WriteString(string(b))
	return ClientBucketRequest, nil
}
func UpdateBucket(c *Client, request *pb.UpdateBucketRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.timeout)*time.Second)
	defer cancel()
	resp, err := c.glusterStorageGatewayClient.UpdateBucket(ctx, request)
	if err != nil {
		fmt.Println("UpdateBucket:", err)
		return err
	}
	fmt.Println("UpdateBucket:", resp)

	return nil
}
func DeleteBucket(c *Client, request *pb.DeleteBucketRequest) error {
	resp, err := c.glusterStorageGatewayClient.DeleteBucket(context.Background(), request)
	if err != nil {
		fmt.Println("DeleteBucket:", err)
		return err
	}
	fmt.Println("DeleteBucket:", resp)
	return nil
}
func main() {
	flag.Parse()
	c, err := NewClient(*confFile)
	var clientBucketRequest *ClientBucketRequest
	if err != nil {
		log.Error("NewClient:", err)
		return
	}
	if *requestBucketType == "c" {
		clientBucketRequest, err = CreateBucket(c, *count)
		if err != nil {
			log.Error("CreateBucket:", err)
		}
		fmt.Printf("finish %v request\n", clientBucketRequest)
		return
	}
	clientBucketRequest = &ClientBucketRequest{}
	b, err := ioutil.ReadFile(bucketInfoFile)
	if err != nil {
		log.Errorln(err)
		return
	}
	if err = json.Unmarshal(b, clientBucketRequest); err != nil {
		log.Errorln(err)
		return
	}
	if *requestBucketType == "d" {
		for i := 0; i < *count; i++ {
			delBucketRequest := &pb.DeleteBucketRequest{
				Name: clientBucketRequest.Request[i].Name,
			}
			DeleteBucket(c, delBucketRequest)
		}

	} else {
		for i := 0; i < *count; i++ {
			updateBucketRequest := &pb.UpdateBucketRequest{
				Name:         clientBucketRequest.Request[i].Name,
				MaxStorageBytes:     clientBucketRequest.Request[i].MaxStorageBytes + rand.Int63(),
				MaxObjectCount: clientBucketRequest.Request[i].MaxObjectCount + rand.Int63()%1024,
			}
			UpdateBucket(c, updateBucketRequest)
		}
	}
}
