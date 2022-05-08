// Copyright 2022 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

package machtesting

/*
#include <mach/mach.h>
*/
import "C"

func TaskSelfTrap() uint32 {
	ret := C.task_self_trap()
	return uint32(ret)
}

func MachThreadSelf() uint32 {
	ret := C.mach_thread_self()
	return uint32(ret)
}
