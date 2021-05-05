// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build tools
// +build tools

// Package tools manages tools using during development.
package tools

import (
	_ "github.com/klauspost/asmfmt/cmd/asmfmt"
)
