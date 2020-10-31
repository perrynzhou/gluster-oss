package put_service

import (
	"fmt"
	"gluster-storage-gateway/bucket"
	"gluster-storage-gateway/conf"
	fs_api "gluster-storage-gateway/fs-api"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"gluster-storage-gateway/protocol/pb"
)

type GrpcService struct {
	addr        string
	grpcPort    int
	httpPort    int
	stopGrpcCh  chan struct{}
	stopHttpCh  chan struct{}
	fsApi       *fs_api.FsApi
	serviceName string
	wg          *sync.WaitGroup
	bucketRequestCh  chan *bucket.BucketInfoRequest
}

//putservice init
//func NewGrpcSerivce(addr string, grpcPort, httpPort int, wg *sync.WaitGroup) *GrpcService {
func NewGrpcSerivce(c *conf.ServerConfig, api *fs_api.FsApi,  serviceName string, wg *sync.WaitGroup) *GrpcService {
	service := &GrpcService{
		addr:        c.Addr,
		grpcPort:    c.GrpcPort,
		httpPort:    c.HttpPort,
		stopGrpcCh:  make(chan struct{}),
		stopHttpCh:  make(chan struct{}),
		fsApi:     api,
		serviceName: serviceName,
		wg:          wg,
		bucketRequestCh:make(chan *bucket.BucketInfoRequest),
	}

	return service
}
func (s *GrpcService) Put(context.Context, *pb.PutObjectRequest) (*pb.PutObjectResponse, error) {
	log.Info("test put function")
	return &pb.PutObjectResponse{}, nil

}

func (s *GrpcService) CreateBucket(ctx context.Context, createBucketRequest *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	req := bucket.NewCreateBucketInfoRequest(createBucketRequest)
	s.bucketRequestCh <-req
	bucketResponse :=<- req.Done

	createBucketResponse := &pb.CreateBucketResponse{
		Requst:createBucketRequest,
		Message: "SUCCESS",
	}
	if bucketResponse.Err !=nil {
		createBucketResponse.Message = bucketResponse.Err.Error()
	}
	return createBucketResponse, bucketResponse.Err
}

func (s *GrpcService) DeleteBucket(context.Context, *pb.DeleteBucketRequest) (*pb.DeleteBucketResponse, error) {
	return &pb.DeleteBucketResponse{}, nil
}

func (s *GrpcService) ListBuckets(context.Context, *pb.ListBucketsRequest) (*pb.ListBucketsResponse, error) {
	return &pb.ListBucketsResponse{}, nil
}

func (s *GrpcService) UpdateBucket(context.Context, *pb.UpdateBucketRequest) (*pb.UpdateBucketResponse, error) {
	return &pb.UpdateBucketResponse{}, nil
}
func (s *GrpcService) AddVolume(context.Context, *pb.AddVolumeRequest) (*pb.AddVolumeResponse, error) {
	return &pb.AddVolumeResponse{}, nil
}

func (s *GrpcService) DeleteVolume(context.Context, *pb.DeleteVolumeRequest) (*pb.DeleteVolumeResponse, error) {
	return &pb.DeleteVolumeResponse{}, nil
}

func (s *GrpcService) ListVolumes(context.Context, *pb.ListVolumesRequest) (*pb.ListVolumesResponse, error) {
	return &pb.ListVolumesResponse{}, nil
}

func (s *GrpcService) Stop() {
	s.stopGrpcCh <- struct{}{}
	s.stopHttpCh <- struct{}{}
}
func (s *GrpcService) Run() {
	s.wg.Add(2)
	go s.runGrpc()
	go s.runHttp()
}
func (s *GrpcService) runHttp() {
	defer s.wg.Done()
	//http gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterFusionStorageGatewayHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", s.grpcPort), dialOptions); err != nil {
		log.Fatal("register http FusionSorageGatewayService failed:", err)
	}
	go func(port int) {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
			log.Infof("start  http FusionSorageGatewayService on %s:%d failed,err:%v", s.addr, s.httpPort, err)
		}
	}(s.httpPort)
	log.Infof("start  http FusionSorageGatewayService on %s:%d  success", s.addr, s.httpPort)
	for {
		select {
		case <-s.stopHttpCh:
			log.Infof("stop http FusionSorageGatewayService on %s:%d success\n", s.addr, s.httpPort)
			return
		}
	}
}
func (s *GrpcService) runGrpc() {
	defer s.wg.Done()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on %d,err: %v", s.grpcPort, err)
	}
	srv := grpc.NewServer()
	pb.RegisterFusionStorageGatewayServer(srv, s)
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
