// Copyright 2019 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,amd64

package mach

import (
	"unsafe"
)

//go:nosplit
func machReplyPort() (ret uint32)

// ReplyPort allocate a port for the caller.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//  mach_trap:26: mach_reply_port
func ReplyPort() uint32 {
	return machReplyPort()
}

//go:nosplit
func threadSelfTrap() (ret uint32)

// ThreadSelfTrap give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//  mach_trap:27: thread_self_trap
func ThreadSelfTrap() uint32 {
	return threadSelfTrap()
}

// MachThreadSelf give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//  mach_trap:27: thread_self_trap
//
// Deprecated: use ThreadSelfTrap instead of.
func MachThreadSelf() uint32 {
	return threadSelfTrap()
}

//go:nosplit
func taskSelfTrap() (ret uint32)

// TaskSelfTrap give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//  mach_trap:28: task_self_trap
func TaskSelfTrap() uint32 {
	return taskSelfTrap()
}

// MachTaskSelf give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//  mach_trap:28: task_self_trap
//
// Deprecated: use TaskSelfTrap instead of.
func MachTaskSelf() uint32 {
	return taskSelfTrap()
}

//go:nosplit
func hostSelfTrap() (ret uint32)

// HostSelfTrap give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//  mach_trap:29: host_self_trap
func HostSelfTrap() uint32 {
	return hostSelfTrap()
}

// MachHostSelf give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//  mach_trap:29: host_self_trap
//
// Deprecated: use HostSelfTrap instead of.
func MachHostSelf() uint32 {
	return hostSelfTrap()
}

//go:nosplit
func machMsgTrap(msg unsafe.Pointer, opt, ssize, rsize, rname, to, not uint32) (ret uint32)

// MsgTrap possibly send a message; possibly receive a message.
//
// Returns the all of mach_msg_send and mach_msg_receive error codes.
//
//  mach_trap:31: mach_msg_trap
func MsgTrap(msg unsafe.Pointer, opt, ssize, rsize, rname, to, not uint32) uint32 {
	return machMsgTrap(msg, opt, ssize, rsize, rname, to, not)
}
