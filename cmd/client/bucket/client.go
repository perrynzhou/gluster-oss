package main

import (
	"context"
	"flag"
	"fmt"
	"glusterfs-storage-gateway/conf"
	"glusterfs-storage-gateway/protocol/pb"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)
const (
	bucketInfoFile = "/tmp/bucket.json"
)
var (
	requestBucketType = flag.String("op", "create", "create bucket")
	confFile =flag.String("c", "./conf.yaml", "default conf is ./conf.yaml")
)

type ClientRequest struct {
	Request  []*pb.CreateBucketRequest
}
type Client struct {
	glusterStorageGatewayClient pb.GlusterStorageGatewayClient
	timeout  int
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
		timeout:config.Conf.TimeOut,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func CreateBucket(c *Client) (*ClientRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.timeout)*time.Second)
	defer cancel()

	clientRequest :=  &ClientRequest {
		Request:make([]*pb.CreateBucketRequest, 10),
	}
	for i := 0; i < 10; i++ {
		clientRequest.Request[i] = &pb.CreateBucketRequest{
			Name:     fmt.Sprintf("bucket%d", i),
			Capacity: 100,
			ObjectsLimit: 100,
		}
		resp, err := c.glusterStorageGatewayClient.CreateBucket(ctx, clientRequest.Request[i])
		if err != nil {
			log.Errorf("CreateBucket err:%v", err)
			return nil,err
		}
		log.Info("resp:", resp)
	}
	b, err := bson.Marshal(&clientRequest)
	if err!= nil {
		log.Errorln("err:",err)
	}
	log.Infoln("valiue:",string(b))
	ioutil.WriteFile(bucketInfoFile, b, os.ModePerm)
	return clientRequest, nil
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
   var clientRequest *ClientRequest
	if err != nil {
		log.Error("NewClient:", err)
		return
	}
	if *requestBucketType == "create" {
		clientRequest, err = CreateBucket(c)
		if err != nil {
			log.Error("CreateBucket:", err)
		}
		fmt.Printf("finish %v request\n", clientRequest)
		return
	}
	clientRequest = &ClientRequest{
		Request:make([]*pb.CreateBucketRequest,0),
	}
	b,err := ioutil.ReadFile(bucketInfoFile)
	if err != nil {
		log.Errorln(err)
		return
	}
	if err = bson.Unmarshal(b, clientRequest);err!=nil {
		log.Errorln(err)
		return
	}
	if *requestBucketType == "delete" {

		for i:=0;i<len(clientRequest.Request);i++ {
			delBucketRequest := &pb.DeleteBucketRequest{
				Name: clientRequest.Request[i].Name,
			}
			DeleteBucket(c, delBucketRequest)
		}

	} else {
		for i:=0;i<len(clientRequest.Request);i++ {
			updateBucketRequest := &pb.UpdateBucketRequest{
				Name:         clientRequest.Request[i].Name,
				Capacity:     clientRequest.Request[i].Capacity + 100,
				ObjectsLimit: clientRequest.Request[i].ObjectsLimit + 1024,
			}
			UpdateBucket(c, updateBucketRequest)
		}
	}
}
