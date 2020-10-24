package main

import (
	"context"
	"fmt"
	"fusion-storage-gateway/conf"
	"fusion-storage-gateway/protocol/pb"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.FusionStorageGatewayClient
	conn   *grpc.ClientConn
}

func NewClient(path string) (*Client, error) {
	config, err := conf.NewClientConf(path)
	if err != nil {
		log.Fatal("Load Config failed:", err)
		return nil, err
	}
	diaOpt := grpc.WithDefaultCallOptions()
	cnn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Addr, config.Port), grpc.WithInsecure(), diaOpt)
	if err != nil {
		log.Error("grpc.Dial Failed, err:", err)
		return nil, nil
	}
	client := pb.NewFusionStorageGatewayClient(cnn)

	return &Client{
		client: client,
		conn:   cnn,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func CreateBucket(c *Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	for i := 0; i < 100; i++ {
		req := &pb.CreateBucketRequest{
			Name: fmt.Sprintf("bucket-%d", i),
		}
		resp, err := c.client.CreateBucket(ctx, req)
		if err != nil {
			log.Errorf("CreateBucket err:%v", err)
			return err
		}
		log.Info("resp:", resp)
	}
	return nil
}
func main() {
	c,err:= NewClient("./conf.yaml")
	if err != nil {
		log.Error("NewClient:",err)
		return
	}
	if err := CreateBucket(c); err != nil {
		log.Error("CreateBucket:", err)
	}
}
