package main

import (
	"github.com/denysvitali/firecracker-ctl/pkg"
	"github.com/sirupsen/logrus"
	"strings"
)

type ActionCmd struct {
	Action string `arg:"positional,required"`
}

const (
	actStart string = "start"
	actCtrlAltDel string = "ctrl-alt-del"
	actFlushMetrics string = "flush-metrics"
)

var validActions = []string{
	actStart,
	actCtrlAltDel,
	actFlushMetrics,
}

func (c ActionCmd) Do(f *firecrackerctl.Client, logger *logrus.Logger) {
	var err error
	switch strings.ToLower(c.Action) {
	case actStart:
		err = f.Start()
	case actCtrlAltDel:
		err = f.SendCtrlAltDel()
	case actFlushMetrics:
		err = f.FlushMetrics()
	default:
		logger.Fatalf("invalid action: please provide any of %s",
			strings.Join(validActions, ", "),
		)
	}

	if err != nil {
		logger.Fatalf("action failed: %v", err)
	}

	logger.Infof("action succeeded!")
}
