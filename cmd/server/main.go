package main

import (
	"flag"
	"glusterfs-storage-gateway/conf"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/manage"
	"glusterfs-storage-gateway/service"
	"glusterfs-storage-gateway/utils"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var (
	confFile    = flag.String("c", "conf.yaml", "glusterfs-storage-gateway conf file")
	serviceName = flag.String("n", "glusterfs-storage-gateway", "glusterfs-storage-gateway name")
)

func init() {
	flag.Parse()
	utils.InitLogFormat()
}
func initStoreBackend(sc *conf.ServerConfig) (*fs_api.FsApi, error) {

	api, err := fs_api.NewFsApi(sc.StoreBackend.Volume, sc.StoreBackend.Addr, sc.StoreBackend.Port, true)
	if err != nil {
		log.Errorln("new metaApi failed")
		return nil, err
	}
	return api, nil
}
func main() {

	serverConf, err := conf.NewServerConf(*confFile)
	log.Info("serverConf:", serverConf)
	if err != nil {
		log.Fatal(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	wg := &sync.WaitGroup{}
	if _,err =utils.InitRedisClient(serverConf.MetaBacked.Addr, serverConf.MetaBacked.Port);err != nil {
		log.Fatalln("conection redis erros:",err)
	}
	fsApi, err := initStoreBackend(serverConf)
	if err != nil {
		log.Fatal("init fsApi failed:", err)
	}
	log.Info("init glusterfs-storage-gateway success")
	bucketService := service.NewBucketSerivce(fsApi, manage.ServiceName, wg)
	service := service.NewService(serverConf, wg)
	service.RegisterService(bucketService.ServiceName, bucketService)
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
