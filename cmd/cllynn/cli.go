package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// CliErrorWriter is used to write logs from urfave/cli's NewExitError
type CliErrorWriter struct{}

// citation: https://github.com/sirupsen/logrus/issues/436#issuecomment-345468176
func (w CliErrorWriter) Write(b []byte) (int, error) {
	n := len(b)
	if n > 0 && b[n-1] == '\n' {
		b = b[:n-1]
	}
	errorString := fmt.Sprintf("%s: %s", appName, string(b))
	log.Error(errorString)
	return n, nil
}
