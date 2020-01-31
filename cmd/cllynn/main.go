package main

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var version string

const appName = "cllynn"

const appDescription = "a mono-cli local developement strategy"

func init() {
	initLogs()
}

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Description = appDescription
	app.Version = version
	// cli.ErrWriter must implement the io.Writer interface
	// ref: https://github.com/urfave/cli/blob/8e01ec4cd3e2d84ab2fe90d8210528ffbb06d8ff/errors.go#L13-L15
	cli.ErrWriter = CliErrorWriter{}

	app.Commands = []*cli.Command{}

	err := app.Run(os.Args)
	if err != nil {
		err = errors.Wrap(err, app.Name)
		log.Error(err)
		os.Exit(1)
	}
}
