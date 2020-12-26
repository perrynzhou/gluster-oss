package service

import (
	"fmt"
	"glusterfs-storage-gateway/conf"
	"glusterfs-storage-gateway/manage"
	"glusterfs-storage-gateway/protocol/pb"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type IService interface {
	Run()
	Stop()
}
type Service struct {
	addr          string
	grpcPort      int
	httpPort      int
	services      map[string]IService
	stopGrpcCh    chan struct{}
	stopHttpCh    chan struct{}
	wg            *sync.WaitGroup
}

func NewService(c *conf.ServerConfig, wg *sync.WaitGroup) *Service {
	return &Service{
		addr:          c.ServiceBackend.BackendAddr,
		grpcPort:      c.ServiceBackend.GrpcPort,
		httpPort:      c.ServiceBackend.HttpPort,
		stopGrpcCh:    make(chan struct{}),
		stopHttpCh:    make(chan struct{}),
		services:make(map[string]IService),
		wg:            wg,
	}
}
func (s *Service)RegisterService(serviceName string,service IService) {
	if _, ok := s.services[serviceName]; !ok {
		s.services[serviceName]=service
	}}
func (s *Service) CreateBucket(ctx context.Context, createBucketRequest *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	bucketService := s.services[manage.ServiceName].(*BucketService)
	return bucketService.CreateBucket(ctx, createBucketRequest)
}
func (s *Service) DeleteBucket(ctx context.Context, deleteBucketRequest *pb.DeleteBucketRequest) (*pb.DeleteBucketResponse, error) {
	bucketService := s.services[manage.ServiceName].(*BucketService)
	return bucketService.DeleteBucket(ctx, deleteBucketRequest)
}
func (s *Service) UpdateBucket(ctx context.Context, updateBucketRequest *pb.UpdateBucketRequest) (*pb.UpdateBucketResponse, error) {
	bucketService := s.services[manage.ServiceName].(*BucketService)
	return bucketService.UpdateBucket(ctx, updateBucketRequest)

}
func (s *Service) Stop() {
	s.stopGrpcCh <- struct{}{}
	s.stopHttpCh <- struct{}{}
}
func (s *Service) Run() {
	s.wg.Add(2)
	for k,v := range s.services{
		log.Info("load ",k," service")
		v.Run()
	}
	go s.runGrpc()
	go s.runHttp()
}
func (s *Service) runHttp() {
	log.Info("start Service Http")
	defer s.wg.Done()
	//http gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	grpcServerEndpoint :=fmt.Sprintf("%s:%d",s.addr, s.grpcPort)
	if err := pb.RegisterGlusterStorageGatewayHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, dialOptions); err != nil {
		log.Fatal("register http service failed:", err)
	}
	log.Infoln("start grpcService on ",grpcServerEndpoint)
	go func(port int) {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
			log.Infof("start  %s:%d failed,err:%v", s.addr, s.httpPort, err)
		}
	}(s.httpPort)
	for {
		select {
		case <-s.stopHttpCh:
			log.Infof("stop http service on %s:%d success\n", s.addr, s.httpPort)
			return
		}
	}
}
func (s *Service) runGrpc() {
	log.Info("start Service GRPC")
	defer s.wg.Done()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on %d,err: %v", s.grpcPort, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGlusterStorageGatewayServer(grpcServer, s)
	go func(srv *grpc.Server) {
		if err := srv.Serve(listen); err != nil {
			log.Fatal("start  grpc service on %s:%d failed:%v ", s.addr, s.grpcPort, err)
		}
	}(grpcServer)
	log.Infof("start  grpc service on %s:%d  success", s.addr, s.grpcPort)
	for {
		select {
		case <-s.stopGrpcCh:
			for k,v := range s.services{
				log.Info("stop ",k," service")
				v.Stop()
			}
			grpcServer.Stop()
			log.Infof("stop grpc service on %s:%d success", s.addr, s.grpcPort)
			return
		}
	}
}
