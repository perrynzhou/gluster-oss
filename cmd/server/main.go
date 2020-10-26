package main

import (
	"flag"
	"gluster-oss/conf"
	fs_api "gluster-oss/fs-api"
	ser "gluster-oss/grpc-service"
	"gluster-oss/utils"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var (
	confFile    = flag.String("c", "conf.yaml", "gluster-oss conf file")
	serviceName = flag.String("n", "gluster-oss-service", "gluster-oss name")
)

func init() {
	flag.Parse()
	utils.InitLogFormat()
}
func initStoreServer(addr string, port int) (*fs_api.FsApi, error) {
	address := strings.Split(addr, ":")
	metaApi, err := fs_api.NewFsApi(address[1], address[0], port, true)
	if err != nil {
		log.Error("new metaApi failed")
		return nil, err
	}

	if err := metaApi.Stat(utils.GlobalObjectMetaSavePath); err != nil {
		if err = metaApi.Mkdir(utils.GlobalObjectMetaSavePath, 0644); err != nil {
			log.Error("new metaApi failed",err)
			return nil, err
		}
	}
	if err := metaApi.Stat(utils.GlobalBucketMetaSavePath); err != nil {
		if err = metaApi.Mkdir(utils.GlobalBucketMetaSavePath, 0644); err != nil {
			log.Error("new metaApi failed",err)
			return nil, err
		}
	}
	return metaApi, nil
}
func main() {

	c, err := conf.NewServiceConf(*confFile)
	if err != nil {
		log.Fatal(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	wg := &sync.WaitGroup{}
	utils.InitRedisClient(c.MetaSrvAddr.Addr, c.MetaSrvAddr.Port)
	metaApi, err := initStoreServer(c.StoreSrvAddr.Addr, c.StoreSrvAddr.Port)
	if err != nil {
		log.Fatal("init metaApi failed:", err)
	}
	dataApi := make(map[string]*fs_api.FsApi)
	dataApi["ssd_vol"] = metaApi
	service := ser.NewGrpcSerivce(c, dataApi, metaApi, *serviceName, wg)
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
