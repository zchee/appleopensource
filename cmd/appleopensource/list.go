// Copyright 2016 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	cli "github.com/alecthomas/kingpin"
	"github.com/zchee/appleopensource"
)

var (
	listTarballs = cmdList.Flag("tarballs", "List the tarballs resources.").Short('t').Bool()
	listSource   = cmdList.Flag("source", "List the source resources.").Short('s').Bool()
	listNoCache  = cmdList.Flag("no-cache", "Disable the cache.").Short('n').Bool()
)

// index return the opensource.apple.com project index, and caches the HTML DOM tree into cacheDir.
func index(typ string) ([]byte, error) {
	cachedir := cacheDir()
	fname := filepath.Join(cachedir, fmt.Sprintf("%s.html", typ))
	if isExist(fname) && !*listNoCache {
		return ioutil.ReadFile(fname)
	}

	if err := os.MkdirAll(cachedir, 0775); err != nil {
		return nil, err
	}

	buf, err := appleopensource.IndexProject(typ)
	if err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(fname, buf, 0664); err != nil {
		return nil, err
	}

	return buf, nil
}

func runList(ctx *cli.ParseContext) error {
	mode := appleopensource.TypeTarballs
	switch {
	case *listSource:
		mode = appleopensource.TypeSource
	case *listTarballs:
		// nothing to do
	}

	index, err := index(mode.String())
	if err != nil {
		log.Fatal(err)
	}

	list, err := appleopensource.ListProject(index)

	var buf bytes.Buffer
	for _, b := range list {
		buf.WriteString(b.Name + "\n")
	}
	// TODO(zchee): trim last new line
	buf.Truncate(buf.Len() - 1)

	fmt.Printf(buf.String())

	return nil
}