// Copyright 2022 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

package machtesting

/*
#include <mach/mach.h>
*/
import "C"

type MachPortName = uint32

type ThreadPort = uint32

type HostNamePort = uint32

func TaskSelfTrap() MachPortName {
	return MachPortName(C.task_self_trap())
}

func MachThreadSelf() ThreadPort {
	return ThreadPort(C.mach_thread_self())
}

func MachHostSelf() HostNamePort {
	return HostNamePort(C.mach_host_self())
}
