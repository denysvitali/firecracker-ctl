package main

import (
	"github.com/alexflint/go-arg"
	"github.com/denysvitali/firecracker-ctl/pkg"
	"github.com/sirupsen/logrus"
)

var args struct {
	SockUrl    string         `arg:"--sock,-s,env:FC_SOCK,required"`
	BootSource *BootSourceCmd `arg:"subcommand:boot-source"`
	Drive      *DriveCmd      `arg:"subcommand:drive"`
	Action     *ActionCmd     `arg:"subcommand:action"`
}

func main() {
	logger := logrus.New()
	arg.MustParse(&args)
	c := firecrackerctl.New(args.SockUrl)

	if args.BootSource != nil {
		args.BootSource.Do(&c, logger)
		return
	}

	if args.Drive != nil {
		args.Drive.Do(&c, logger)
		return
	}

	if args.Action != nil {
		args.Action.Do(&c, logger)
		return
	}

	logger.Fatalf("subcommand not specified")
}
