// Copyright 2022 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

package machtesting

/*
#include <mach/mach.h>
#include <mach/mach_interface.h>
*/
import "C"
import "github.com/go-darwin/mach"

func TaskSelfTrap() mach.MachPortName {
	return mach.MachPortName(C.task_self_trap())
}

func MachThreadSelf() mach.ThreadPort {
	return mach.ThreadPort(C.mach_thread_self())
}

func MachHostSelf() mach.HostNamePort {
	return mach.HostNamePort(C.mach_host_self())
}
