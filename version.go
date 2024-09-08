// Package version records versioning information about this module.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	version = "v0.1.0-0-g00000000"
	rc      = "0"
)

func showVersion() {
	re := regexp.MustCompile(`v?(.*)-(\d+)-g([a-f0-9]{7,8})`)
	matches := re.FindStringSubmatch(version)
	if len(matches) != 4 {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version)
		return
	}

	v := "v" + matches[1]

	if n, err := strconv.Atoi(rc); err == nil && n > 0 {
		v += "-rc." + rc
	}

	fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), v)
}
