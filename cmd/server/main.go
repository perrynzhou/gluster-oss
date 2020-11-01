package main

import (
	"flag"
	"gluster-storage-gateway/conf"
	fs_api "gluster-storage-gateway/fs-api"
	"gluster-storage-gateway/service"
	"gluster-storage-gateway/utils"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var (
	confFile    = flag.String("c", "server_conf.yaml", "gluster-storage-gateway conf file")
	serviceName = flag.String("n", "gluster-storage-gateway", "gluster-storage-gateway name")
)

func init() {
	flag.Parse()
	utils.InitLogFormat()
}
func initStoreBackend(sc *conf.ServerConfig) (*fs_api.FsApi, error) {

	api, err := fs_api.NewFsApi(sc.StoreBackend.Volume, sc.StoreBackend.Addr, sc.StoreBackend.Port, true)
	if err != nil {
		log.Error("new metaApi failed")
		return nil, err
	}
	return api, nil
}
func main() {

	serverConf, err := conf.NewServerConf(*confFile)
	log.Info("serverConf:",serverConf)
	if err != nil {
		log.Fatal(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	wg := &sync.WaitGroup{}
	utils.InitRedisClient(serverConf.MetaBacked.Addr, serverConf.MetaBacked.Port)
	fsApi, err := initStoreBackend(serverConf)
	if err != nil {
		log.Fatal("init fsApi failed:", err)
	}
	log.Info("init gluster-storage-gateway success")
	bucketService := service.NewBucketSerivce(fsApi, "bucket", wg)
	service := service.NewService(serverConf,wg)
	service.RegisterService(bucketService.ServiceName,bucketService)
	service.Run()
	defer wg.Wait()
	for {
		select {
		case <-signals:
			service.Stop()
			return
		}
	}

}
