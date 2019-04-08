// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli"
)

var noCache bool

func initCmd(ctx *cli.Context) error {
	noCache = ctx.Bool("no-cache")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "appleopensource"
	app.Usage = "An opensource.apple.com resource management tool."
	app.Version = fmt.Sprintf("%s (%s)", version, gitCommit)
	app.Before = initCmd

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug mode",
		},
		cli.BoolFlag{
			Name:  "no-cache, n",
			Usage: "Disable cache",
		},
	}

	app.Commands = []cli.Command{
		cacheCommand,
		fetchCommand,
		listCommand,
		releaseCommand,
		versionsCommand,
	}

	cli.ErrWriter = &fatalWriter{cli.ErrWriter}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

type fatalWriter struct {
	cliErrWriter io.Writer
}

func (f *fatalWriter) Write(b []byte) (n int, err error) {
	return f.cliErrWriter.Write(b)
}