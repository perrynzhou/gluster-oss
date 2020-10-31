package service

import (
	"fmt"
	"gluster-storage-gateway/bucket"
	"gluster-storage-gateway/conf"
	fs_api "gluster-storage-gateway/fs-api"
	"gluster-storage-gateway/meta"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"gluster-storage-gateway/protocol/pb"
)

type BucketService struct {
	addr            string
	grpcPort        int
	httpPort        int
	stopGrpcCh      chan struct{}
	stopHttpCh      chan struct{}
	fsApi           *fs_api.FsApi
	serviceName     string
	wg              *sync.WaitGroup
	bucketRequestCh chan *bucket.BucketInfoRequest
}

func NewBucketSerivce(c *conf.ServerConfig, api *fs_api.FsApi, serviceName string, wg *sync.WaitGroup) *BucketService {
	service := &BucketService{
		addr:            c.Addr,
		grpcPort:        c.GrpcPort,
		httpPort:        c.HttpPort,
		stopGrpcCh:      make(chan struct{}),
		stopHttpCh:      make(chan struct{}),
		fsApi:           api,
		serviceName:     serviceName,
		wg:              wg,
		bucketRequestCh: make(chan *bucket.BucketInfoRequest),
	}

	return service
}

func (s *BucketService) CreateBucket(ctx context.Context, createBucketRequest *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	req := bucket.NewCreateBucketInfoRequest(createBucketRequest)
	s.bucketRequestCh <- req
	resp := <-req.Done
	bucketInfo := resp.Reply.(*meta.BucketInfo)
	createBucketResponse := &pb.CreateBucketResponse{
		Name: bucketInfo.Name,
		//request storage capacity
		Capacity: bucketInfo.UsageInfo.CapacityLimitSize,
		//obejcts limits
		ObjectsLimit: bucketInfo.UsageInfo.ObjectsLimitCount,
		BucketDir:    bucketInfo.RealDirName,
		Message:      "success",
	}
	if resp.Err != nil {
		createBucketResponse.Message = resp.Err.Error()
	}
	return createBucketResponse, resp.Err
}

func (s *BucketService) DeleteBucket(ctx context.Context, deleteBucketRequest *pb.DeleteBucketRequest) (*pb.DeleteBucketResponse, error) {
	req := bucket.NewDeleteBucketInfoRequest(deleteBucketRequest)
	s.bucketRequestCh <- req
	resp := <-req.Done
	bucketInfo := resp.Reply.(*meta.BucketInfo)
	deleteBucketResponse := &pb.DeleteBucketResponse{
		Name:         bucketInfo.Name,
		ObjectsLimit: bucketInfo.UsageInfo.ObjectsLimitCount,
		Capacity:     bucketInfo.UsageInfo.CapacityLimitSize,
		ObjectCount:  bucketInfo.UsageInfo.ObjectsCurrentCount,
	}
	if resp.Err != nil {
		deleteBucketResponse.Message = resp.Err.Error()
	}
	return deleteBucketResponse, resp.Err
}
func (s *BucketService) UpdateBucket(ctx context.Context, updateBucketRequest *pb.UpdateBucketRequest) (*pb.UpdateBucketResponse, error) {
	req := bucket.NewUpdateBucketInfoRequest(updateBucketRequest)
	s.bucketRequestCh <- req
	resp := <-req.Done
	bucketInfo := resp.Reply.(*meta.BucketInfo)
	updateBucketResponse := &pb.UpdateBucketResponse{
		Name:         bucketInfo.Name,
		ObjectsLimit: bucketInfo.UsageInfo.ObjectsLimitCount,
		Capacity:     bucketInfo.UsageInfo.CapacityLimitSize,
		ObjectCount:  bucketInfo.UsageInfo.ObjectsCurrentCount,
		BucketDir:    bucketInfo.RealDirName,
	}
	updateBucketResponse.Message = "success"
	if resp.Err != nil {
		updateBucketResponse.Message = resp.Err.Error()
	}
	return updateBucketResponse, resp.Err
}
func (s *BucketService) Stop() {
	s.stopGrpcCh <- struct{}{}
	s.stopHttpCh <- struct{}{}
}
func (s *BucketService) Run() {
	s.wg.Add(2)
	go s.runGrpc()
	go s.runHttp()
}
func (s *BucketService) runHttp() {
	defer s.wg.Done()
	//http gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterGlusterStorageGatewayBucketHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", s.grpcPort), dialOptions); err != nil {
		log.Fatal("register http GlusterStorageGatewayBucket failed:", err)
	}
	go func(port int) {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
			log.Infof("start  http GlusterStorageGatewayBucket on %s:%d failed,err:%v", s.addr, s.httpPort, err)
		}
	}(s.httpPort)
	log.Infof("start  http GlusterStorageGatewayBucket on %s:%d  success", s.addr, s.httpPort)
	for {
		select {
		case <-s.stopHttpCh:
			log.Infof("stop http GlusterStorageGatewayBucket on %s:%d success\n", s.addr, s.httpPort)
			return
		}
	}
}
func (s *BucketService) runGrpc() {
	defer s.wg.Done()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on %d,err: %v", s.grpcPort, err)
	}
	srv := grpc.NewServer()
	pb.RegisterGlusterStorageGatewayBucketServer(srv, s)
	go func(srv *grpc.Server) {
		if err := srv.Serve(listen); err != nil {
			log.Fatal("start  grpc FusionSorageGatewayService on %s:%d failed:%v ", s.addr, s.grpcPort, err)
		}
	}(srv)
	log.Infof("start  grpc FusionSorageGatewayService on %s:%d  success", s.addr, s.grpcPort)
	for {
		select {
		case <-s.stopGrpcCh:
			srv.Stop()
			log.Infof("stop grpc FusionSorageGatewayService on %s:%d success", s.addr, s.grpcPort)
			return
		}
	}
}
