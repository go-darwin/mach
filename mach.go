// SPDX-FileCopyrightText: Copyright 2019 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package mach

import (
	"unsafe"

	"github.com/go-darwin/sys"
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

//go:nosplit
func pfzExit() (ret sys.KernReturn)

// PfzExit called from commpage to take a delayed preemption when exiting
// the "Preemption Free Zone" (PFZ).
//
//  mach_trap:58: pfz_exit
func PfzExit() (ret sys.KernReturn) {
	return pfzExit()
}

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

//go:nosplit
func mkTimerCreateTrap() (ret MachPortName)

// MkTimerCreateTrap makes timer_create trap.
//
//  mach_trap:91: mk_timer_create_trap
func MkTimerCreateTrap() (ret MachPortName) {
	return mkTimerCreateTrap()
}
