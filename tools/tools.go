// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc && tools
// +build darwin,amd64,gc,tools

package tools

import (
	_ "github.com/klauspost/asmfmt/cmd/asmfmt"
	_ "go-darwin.dev/tools/cmd/asmvet"
	_ "gotest.tools/gotestsum"
)
