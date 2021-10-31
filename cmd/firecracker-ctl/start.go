package main

import (
	"github.com/denysvitali/firecracker-ctl/pkg"
	"github.com/sirupsen/logrus"
)

type StartCmd struct {

}

func (c StartCmd) Do(f *firecrackerctl.Client, logger *logrus.Logger) {
	err := f.Start()
	if err != nil {
		logger.Fatalf("unable to set boot source: %v", err)
	}
	logger.Infof("successfully started")
}
