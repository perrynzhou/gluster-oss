package main

import (
	"os"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	cmd "gluster-gtw/cmd/client-tools/cli"
    "github.com/stretchr/testify/assert"
)

func main() {
	logrus.SetOutput(colorable.NewColorableStdout())
	if err := mainErr(); err != nil {
		logrus.Fatal(err)
	}
}
type Test struct {
	Name string
	Value int
}
func mainErr() error {

	app := cli.NewApp()
	assert.NotNil(app, nil)
	app.Name = "fusion-tools"
	app.Version = "v1.1"
	app.Usage = "fusion-tools 是 gluster-gtw 的客户端工具"
	app.Author = "周琳(11095119)，孙伟(11102494)"
	app.Commands = []cli.Command{
		cmd.BucketCommond(),
		cmd.VolumeCommond(),
	}
	return app.Run(os.Args)


}
