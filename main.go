// Copyright 2014 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	rate_lib "github.com/synful/rate"
)

var (
	rate = flag.String("rate", "1MB", "rate limit in data per second; units are (B, KB, KiB, MB, MiB, GB, GiB)")
)

var units = map[string]int{
	"B":   1,
	"KB":  1000,
	"KiB": 1024,
	"MB":  1000 * 1000,
	"MiB": 1024 * 1024,
	"GB":  1000 * 1000 * 1000,
	"GiB": 1024 * 1024 * 1024,
}

const (
	ERR_USAGE = 2 + iota
)

func main() {
	flag.Parse()

	str := *rate
	var rate float64
	var unit string
	n, err := fmt.Sscanf(str, "%f%s", &rate, &unit)
	size, ok := units[unit]
	if n != 2 || err != nil || !ok {
		flag.PrintDefaults()
		os.Exit(ERR_USAGE)
	}
	io.Copy(os.Stdout, rate_lib.NewLimitReader(os.Stdin, uint64(float64(size)*rate)))
}
