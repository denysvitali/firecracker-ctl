package main

import (
	"github.com/denysvitali/firecracker-ctl/pkg"
	"github.com/sirupsen/logrus"
)

type DriveCmd struct {
	DriveId      string `arg:"-i,--id,required"`
	IsReadOnly   bool   `arg:"-r,--read-only,required"`
	IsRootDevice bool   `arg:"-R,--root-device,required"`
	PathOnHost   string `arg:"-p,--path,required"`

	// Optional
	PartUuid string `arg:"-u,--part-uuid"`
	// TODO: add RateLimiter
}

func (c DriveCmd) Do(f *firecrackerctl.Client, logger *logrus.Logger) {
	err := f.Drives(firecrackerctl.DrivesRequest{
		DriveId:      c.DriveId,
		IsReadOnly:   c.IsReadOnly,
		IsRootDevice: c.IsRootDevice,
		PathOnHost:   c.PathOnHost,
		PartUuid:     c.PartUuid,
	})
	if err != nil {
		logger.Fatalf("unable to set boot source: %v", err)
	}
	logger.Infof("successfully set %s", c.DriveId)
}
