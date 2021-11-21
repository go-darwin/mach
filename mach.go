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
func machReplyPort() (ret MachPortName)

// ReplyPort allocate a port for the caller.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
func ReplyPort() MachPortName {
	return machReplyPort()
}

//go:noescape
//go:nosplit
func threadSelfTrap() (ret MachPortName)

// ThreadSelfTrap give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
func ThreadSelfTrap() MachPortName {
	return threadSelfTrap()
}

// MachThreadSelf give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: use ThreadSelfTrap instead of.
func MachThreadSelf() MachPortName {
	return threadSelfTrap()
}

//go:noescape
//go:nosplit
func taskSelfTrap() (ret MachPortName)

// TaskSelfTrap give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
func TaskSelfTrap() MachPortName {
	return taskSelfTrap()
}

// MachTaskSelf give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: use TaskSelfTrap instead of.
func MachTaskSelf() MachPortName {
	return taskSelfTrap()
}

//go:noescape
//go:nosplit
func hostSelfTrap() (ret MachPortName)

// HostSelfTrap give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
func HostSelfTrap() MachPortName {
	return hostSelfTrap()
}

// MachHostSelf give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: use HostSelfTrap instead of.
func MachHostSelf() MachPortName {
	return hostSelfTrap()
}

//go:noescape
//go:nosplit
func machMsgTrap(msg *MachMsgHeader, option MachMsgOption, sendSize, rcvSize MachMsgSize, rcvName MachPortName, timeout MachMsgTimeout, notify MachPortName) (ret MachMsgReturn)

// MsgTrap possibly send a message; possibly receive a message.
//
// Returns the all of mach_msg_send and mach_msg_receive error codes.
func MsgTrap(msg *MachMsgHeader, option MachMsgOption, sendSize, rcvSize MachMsgSize, rsize, rcvName MachPortName, timeout MachMsgTimeout, notify MachPortName) MachMsgReturn {
	return machMsgTrap(msg, option, sendSize, rcvSize, rcvName, timeout, notify)
}

//go:noescape
//go:nosplit
func threadGetSpecialReplyPort() (ret MachPortName)

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
//
//  mach_trap:50: thread_get_special_reply_port
func ThreadGetSpecialReplyPort() (ret MachPortName) {
	return threadGetSpecialReplyPort()
}

//go:noescape
//go:nosplit
func pfzExit() (ret sys.KernReturn)

// PfzExit called from commpage to take a delayed preemption when exiting
// the "Preemption Free Zone" (PFZ).
//
//  mach_trap:58: pfz_exit
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
//
//  mach_trap:59: swtch_pri
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
//
//  mach_trap:60: swtch
func Swtch() (ret bool) {
	return swtch()
}

//go:noescape
//go:nosplit
func mkTimerCreateTrap() (ret MachPortName)

// MkTimerCreateTrap makes timer_create trap.
//
//  mach_trap:91: mk_timer_create_trap
func MkTimerCreateTrap() (ret MachPortName) {
	return mkTimerCreateTrap()
}
