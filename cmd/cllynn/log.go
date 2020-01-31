package main

import (
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type customFormatter struct {
	textFormatter log.TextFormatter
}

func (f *customFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := entry.Message

	switch entry.Level {
	case log.InfoLevel:
		msg = color.BlueString(msg)
	case log.WarnLevel:
		msg = color.YellowString(msg)
	case log.ErrorLevel:
		msg = color.RedString(msg)
	}

	entry.Message = msg
	return f.textFormatter.Format(entry)
}

func initLogs() {
	log.SetFormatter(&customFormatter{
		textFormatter: log.TextFormatter{
			DisableTimestamp: true,
		},
	})
}

// setupDebugFlag parses if the --debug flag was given and will (probably) eventually
// parse a --quiet flag too, and do other cli action specific setup
func setupDebugFlag(c *cli.Context) {
	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}
}

var debugModeFlag = cli.BoolFlag{
	Name:    "debug",
	Usage:   "Show debug logs",
	EnvVars: []string{"DEBUG"},
}
