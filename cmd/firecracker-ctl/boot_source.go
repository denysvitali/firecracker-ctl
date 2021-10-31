package main

import (
	"github.com/denysvitali/firecracker-ctl/pkg"
	"github.com/sirupsen/logrus"
)

type BootSourceCmd struct {
	KernelImage string `arg:"positional,required"`
	Initrd      string `arg:"--initrd,-i"`
	BootArgs    string `arg:"--boot-args,-b"`
}

func (c BootSourceCmd) Do(f *firecrackerctl.Client, logger *logrus.Logger) {
	err := f.BootSource(firecrackerctl.BootSourceRequest{
		KernelImagePath: c.KernelImage,
		InitrdPath:      c.Initrd,
		BootArgs:        c.BootArgs,
	})
	if err != nil {
		logger.Fatalf("unable to set boot source: %v", err)
	}
	logger.Infof("successfully set boot source")
}
