// Copyright 2019 The go-python Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux && !android) || dragonfly || openbsd || freebsd
// +build linux,!android dragonfly openbsd freebsd

package main

const (
	libExt       = ".so"
	extraGccArgs = ""
)
