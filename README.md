<!--
Copyright 2014 The Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE fil
-->

rate-limit
==========

[![Build Status](https://travis-ci.org/synful/rate-limit.svg)](https://travis-ci.org/synful/rate-limit)

A simple rate limiting command-line utility. Stdin is piped to stdout, and the rate at which data is copied is limited.

Usage:
```
Usage of rate-limit:
  -rate="1MB": rate limit in data per second; units are (B, KB, KiB, MB, MiB, GB, GiB)
```

Installation: `go get github.com/joshlf13/rate-limit`