// Copyright 2019 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

package mach

import (
	"github.com/go-darwin/sys"
)

//go:noescape
//go:nosplit
func machReplyPort() (ret uint32)

// ReplyPort allocate a port for the caller.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
func ReplyPort() uint32 {
	return machReplyPort()
}

//go:noescape
//go:nosplit
func threadSelfTrap() (ret uint32)

// ThreadSelfTrap give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
func ThreadSelfTrap() uint32 {
	return threadSelfTrap()
}

// MachThreadSelf give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: use ThreadSelfTrap instead of.
func MachThreadSelf() uint32 {
	return threadSelfTrap()
}

//go:noescape
//go:nosplit
func taskSelfTrap() (ret uint32)

// TaskSelfTrap give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
func TaskSelfTrap() uint32 {
	return taskSelfTrap()
}

// MachTaskSelf give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: use TaskSelfTrap instead of.
func MachTaskSelf() uint32 {
	return taskSelfTrap()
}

//go:noescape
//go:nosplit
func hostSelfTrap() (ret uint32)

// HostSelfTrap give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
func HostSelfTrap() uint32 {
	return hostSelfTrap()
}

// MachHostSelf give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: use HostSelfTrap instead of.
func MachHostSelf() uint32 {
	return hostSelfTrap()
}

//go:noescape
//go:nosplit
func machMsgTrap(msg *MachMsgHeader, option MachMsgOption, sendSize, rcvSize MachMsgSize, rcvName MachPortName, timeout MachMsgTimeout, notify MachPortName) (ret int32)

// MsgTrap possibly send a message; possibly receive a message.
//
// Returns the all of mach_msg_send and mach_msg_receive error codes.
func MsgTrap(msg *MachMsgHeader, option MachMsgOption, sendSize, rcvSize MachMsgSize, rsize, rcvName MachPortName, timeout MachMsgTimeout, notify MachPortName) int32 {
	return machMsgTrap(msg, option, sendSize, rcvSize, rcvName, timeout, notify)
}

//go:noescape
//go:nosplit
func threadGetSpecialReplyPort() (ret uint32)

// ThreadGetSpecialReplyPort allocate a special reply port for the calling thread.
//
// Returns the
//
//  mach_port_name_t
// send right & receive right for special reply port.
//  MACH_PORT_NULL
// if there are any resource failures.
//
// or other errors.
func ThreadGetSpecialReplyPort() (ret uint32) {
	return threadGetSpecialReplyPort()
}

//go:noescape
//go:nosplit
func pfzExit() (ret sys.KernReturn)

// PfzExit called from commpage to take a delayed preemption when exiting
// the "Preemption Free Zone" (PFZ).
func PfzExit() (ret sys.KernReturn) {
	return pfzExit()
}

//go:noescape
//go:nosplit
func swtchPri() (ret bool)

// SwtchPri attempt to context switch (logic in
// thread_block no-ops the context switch if nothing would happen).
//
// A boolean is returned that indicates whether there is anything
// else runnable.
//
// That's no excuse to spin, though.
func SwtchPri() (ret bool) {
	return swtchPri()
}

//go:noescape
//go:nosplit
func swtch() (ret bool)

// Swtch attempt to context switch (logic in
// thread_block no-ops the context switch if nothing would happen).
//
// A boolean is returned that indicates whether there is anything
// else runnable.
//
// That's no excuse to spin, though.
func Swtch() (ret bool) {
	return swtch()
}

//go:noescape
//go:nosplit
func mkTimerCreateTrap() (ret uint32)

// MkTimerCreateTrap makes timer_create trap.
func MkTimerCreateTrap() (ret uint32) {
	return mkTimerCreateTrap()
}
