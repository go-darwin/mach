// Copyright 2019 The mach Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,amd64

package mach

import (
	"unsafe"
)

//go:nosplit
func machReplyPort() uintptr

// MachReplyPort allocate a port for the caller.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// mach_trap:26:mach_reply_port
func MachReplyPort() uintptr {
	return machReplyPort()
}

//go:nosplit
func threadSelfTrap() uintptr

// ThreadSelfTrap give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// mach_trap:27:thread_self_trap
func ThreadSelfTrap() uintptr {
	return threadSelfTrap()
}

//go:nosplit
func taskSelfTrap() uintptr

// TaskSelfTrap give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// mach_trap:28:task_self_trap
func TaskSelfTrap() uintptr {
	return taskSelfTrap()
}

//go:nosplit
func hostSelfTrap() uintptr

// HostSelfTrap give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// mach_trap:29:host_self_trap
func HostSelfTrap() uintptr {
	return hostSelfTrap()
}

//go:nosplit
func machMsgTrap(unsafe.Pointer, uint32, uint32, uint32, uint32, uint32, uint32) uint32

// MachMsgTrap possibly send a message; possibly receive a message.
//
// Returns the all of mach_msg_send and mach_msg_receive error codes.
//
// mach_trap:31:mach_msg_trap
func MachMsgTrap(msg unsafe.Pointer, opt, ssize, rsize, rname, to, not uint32) uint32 {
	return machMsgTrap(msg, opt, ssize, rsize, rname, to, not)
}
