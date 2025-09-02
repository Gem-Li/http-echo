// Copyright(C) 2025 Baidu Inc. All Rights Reserved.
// Author: Bo li (libo57@baidu.com)
// Date: 2025/09/02
// Description buildinfo.go:

package buildinfo

import (
	"flag"
	"fmt"
	"os"
)

var (
	version = flag.Bool("v", false, "print version and exit")
	Version string
)

func InitBuildInfo() {
	if *version {
		printVersion()
		os.Exit(0)
	}
}

func init() {
	oldUsage := flag.Usage
	flag.Usage = func() {
		oldUsage()
		printVersion()
	}
}

func printVersion() {
	fmt.Fprintf(flag.CommandLine.Output(), "Version: %s\n", Version)
}
