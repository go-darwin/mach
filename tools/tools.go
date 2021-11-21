// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc && tools
// +build darwin,amd64,gc,tools

package tools

import (
	_ "github.com/go-darwin/tools/cmd/asmvet"
	_ "github.com/klauspost/asmfmt/cmd/asmfmt"
	_ "gotest.tools/gotestsum"
)
