package main

import (
	"flag"
	"gluster-storage-gateway/conf"
	fs_api "gluster-storage-gateway/fs-api"
	"gluster-storage-gateway/service"
	"gluster-storage-gateway/utils"
	"os"
	"os/signal"
	"strings"
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
	address := strings.Split(sc.Addr, ":")
	api, err := fs_api.NewFsApi(address[1], address[0], sc.StoreBackend.Port, true)
	if err != nil {
		log.Error("new metaApi failed")
		return nil, err
	}
	return api, nil
}
func main() {

	c, err := conf.NewServerConf(*confFile)
	if err != nil {
		log.Fatal(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	wg := &sync.WaitGroup{}
	utils.InitRedisClient(c.MetaBacked.Addr, c.MetaBacked.Port)
	fsApi, err := initStoreBackend(c)
	if err != nil {
		log.Fatal("init fsApi failed:", err)
	}
	bucketService := service.NewBucketSerivce(c, fsApi, *serviceName, wg)
	service := service.NewService(bucketService)
	service.Run()
	defer wg.Wait()
	for {
		select {
		case <-signals:
			bucketService.Stop()
			return
		}
	}

}
