package cli

import (
	"github.com/urfave/cli"
	"gluster-oss/cmd/client"
)

func VolumeCommond() cli.Command {
	VolumeFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "config, cfg",
			Usage:  "gluster_api 的配置文件",
			Value:  "./config.json",
			EnvVar: "FUSION_TOOLS_CONFIG",
		},
		cli.BoolFlag{
			Name:  "list, l",
			Usage: "查询所有的volume信息",
		},
	}
	return cli.Command{
		Name:      "volume",
		Usage:     "Volume operation",
		UsageText: "例子: fusion-tools volume -l",
		Action:    Volume,
		Flags:     VolumeFlags,
	}
}

func Volume(ctx *cli.Context) error {
	configFile := ctx.String("config")
	client, err := client.NewClient(configFile)
	if err != nil {
		return err
	}
	defer client.Close()

	list := ctx.Bool("list")
	if list {
		return client.ListVolumes()
	}
	return nil
}