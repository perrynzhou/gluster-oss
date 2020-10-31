package main

import (
	"context"
	"flag"
	"fmt"
	"gluster-storage-gateway/conf"
	"gluster-storage-gateway/protocol/pb"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	requestBucketType = flag.String("op", "create", "create bucket")
)

type ClientRequest struct {
	Requests []*pb.CreateBucketRequest
}
type Client struct {
	glusterStorageGatewayClient pb.GlusterStorageGatewayBucketClient
	conn                        *grpc.ClientConn
}

func NewClient(path string) (*Client, error) {
	config, err := conf.NewClientConf(path)
	if err != nil {
		log.Fatal("Load Config failed:", err)
		return nil, err
	}
	diaOpt := grpc.WithDefaultCallOptions()
	cnn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.StorageGatewayAddr, config.StorageGatewayPort), grpc.WithInsecure(), diaOpt)
	if err != nil {
		log.Error("grpc.Dial Failed, err:", err)
		return nil, nil
	}
	glusterStorageGatewayClient := pb.NewGlusterStorageGatewayBucketClient(cnn)

	return &Client{
		glusterStorageGatewayClient: glusterStorageGatewayClient,
		conn:                        cnn,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func CreateBucket(c *Client) (*ClientRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	clientRequest := &ClientRequest{
		Requests: make([]*pb.CreateBucketRequest, 0),
	}
	for i := 0; i < 100; i++ {
		req := &pb.CreateBucketRequest{
			Name:     fmt.Sprintf("bucket%d", i),
			Capacity: 10234,
			//obejcts limits
			ObjectsLimit: 100,
		}
		resp, err := c.glusterStorageGatewayClient.CreateBucket(ctx, req)
		if err != nil {
			log.Errorf("CreateBucket err:%v", err)
			return nil, err
		}
		clientRequest.Requests = append(clientRequest.Requests, req)
		log.Info("resp:", resp)
	}
	b, _ := bson.Marshal(clientRequest)
	ioutil.WriteFile("/tmp/clientRequest.bson", b, os.ModePerm)
	return clientRequest, nil
}
func UpdateBucket(c *Client, request *pb.UpdateBucketRequest) error {
	resp, err := c.glusterStorageGatewayClient.UpdateBucket(context.Background(), request)
	if err != nil {
		fmt.Println("UpdateBucket:%v", err)
		return err
	}
	fmt.Println("UpdateBucket:%v", resp)
	return nil
}
func DeleteBucket(c *Client, request *pb.DeleteBucketRequest) error {
	resp, err := c.glusterStorageGatewayClient.DeleteBucket(context.Background(), request)
	if err != nil {
		fmt.Println("DeleteBucket:%v", err)
		return err
	}
	fmt.Println("DeleteBucket:%v", resp)
	return nil
}
func main() {
	c, err := NewClient("../server_conf.yaml")
	flag.Parse()
	if err != nil {
		log.Error("NewClient:", err)
		return
	}
	var clientRequest *ClientRequest
	if *requestBucketType == "create" {
		clientRequest, err = CreateBucket(c)
		if err != nil {
			log.Error("CreateBucket:", err)
		}
		fmt.Printf("finish %v request\n", clientRequest)
	} else if *requestBucketType == "delete" {
		b, err := ioutil.ReadFile("/tmp/clientRequest.bson")
		if err == nil {
			clientRequest = &ClientRequest{}
			err = bson.Unmarshal(b, clientRequest)
			if err == nil {
				fmt.Println("read %v", clientRequest)
				i := rand.Intn(len(clientRequest.Requests) - 1)
				request := clientRequest.Requests[i]
				delBucketRequest := &pb.DeleteBucketRequest{
					Name: request.Name,
				}
				request.ObjectsLimit = 8901
				request.Capacity = 102400

				DeleteBucket(c, delBucketRequest)
			}

		}
	} else if *requestBucketType == "update" {
		b, err := ioutil.ReadFile("/tmp/clientRequest.bson")
		if err == nil {
			clientRequest = &ClientRequest{}
			if err = bson.Unmarshal(b, clientRequest); err != nil {
				fmt.Println("UpdateBucket:%", err)
				return
			}
		}
		for _, req := range clientRequest.Requests {
			updateBucketRequest := &pb.UpdateBucketRequest{
				Name:         req.Name,
				Capacity:     req.Capacity + 100,
				ObjectsLimit: req.ObjectsLimit + 1024,
			}
			UpdateBucket(c, updateBucketRequest)
		}
	}
}
