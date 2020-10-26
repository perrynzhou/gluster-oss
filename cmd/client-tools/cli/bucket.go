package cli

import (
	"github.com/urfave/cli"
	"gluster-oss/cmd/client"
)

func BucketCommond() cli.Command {
	BucketFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "config, cfg",
			Usage:  "gluster_api 的配置文件",
			Value:  "./config.json",
			EnvVar: "FUSION_TOOLS_CONFIG",
		},
		cli.BoolFlag{
			Name:  "list, l",
			Usage: "查询所有的bucket信息",
		},
	}
	return cli.Command{
		Name:      "bucket",
		Usage:     "bucket operation",
		UsageText: "例子: fusion-tools bucket -l",
		Action:    Bucket,
		Flags:     BucketFlags,
	}
}

func Bucket(ctx *cli.Context) error {
	configFile := ctx.String("config")
	client, err := client.NewClient(configFile)
	if err != nil {
		return err
	}
	defer client.Close()

	list := ctx.Bool("list")
	if list {
		return client.ListBuckts()
	}
	return nil
}
