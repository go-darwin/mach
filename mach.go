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
//
//go:nosplit
func ReplyPort() uint32 {
	return machReplyPort()
}

//go:noescape
//go:nosplit
func threadSelfTrap() (ret uint32)

// ThreadSelfTrap give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//go:nosplit
func ThreadSelfTrap() uint32 {
	return threadSelfTrap()
}

// MachThreadSelf give the caller send rights for his own thread port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: Use ThreadSelfTrap instead of.
//
//go:nosplit
func MachThreadSelf() uint32 {
	return threadSelfTrap()
}

//go:noescape
//go:nosplit
func taskSelfTrap() (ret uint32)

// TaskSelfTrap give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//go:nosplit
func TaskSelfTrap() uint32 {
	return taskSelfTrap()
}

// MachTaskSelf give the caller send rights for his own task port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: Use TaskSelfTrap instead of.
//
//go:nosplit
func MachTaskSelf() uint32 {
	return taskSelfTrap()
}

//go:noescape
//go:nosplit
func hostSelfTrap() (ret uint32)

// HostSelfTrap give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
//go:nosplit
func HostSelfTrap() uint32 {
	return hostSelfTrap()
}

// MachHostSelf give the caller send rights for his own host port.
//
// Returns the MACH_PORT_NULL if there are any resource failures or other errors.
//
// Deprecated: Use HostSelfTrap instead of.
//
//go:nosplit
func MachHostSelf() uint32 {
	return hostSelfTrap()
}

//go:noescape
//go:nosplit
func machMsgTrap(msg *MachMsgHeader, option MachMsgOption, sendSize, rcvSize MachMsgSize, rcvName MachPort, timeout MachMsgTimeout, notify MachPort) (ret int32)

// MsgTrap possibly send a message; possibly receive a message.
//
// Returns the all of mach_msg_send and mach_msg_receive error codes.
//
//go:nosplit
func MsgTrap(msg *MachMsgHeader, option MachMsgOption, sendSize, rcvSize MachMsgSize, rsize, rcvName MachPort, timeout MachMsgTimeout, notify MachPort) int32 {
	return machMsgTrap(msg, option, sendSize, rcvSize, rcvName, timeout, notify)
}

//go:noescape
//go:nosplit
func machMsgOverwriteTrap(sendMsg *MachMsgHeader, option MachMsgOption, sendSize, receiveLimit MachMsgSize, receiveName MachPort, timeout MachMsgTimeout, notify MachPort, receiveMsg *MachMsgHeader, receiveMsgSize MachMsgSize) (ret int32)

// MsgOverwriteTrap send and/or receive messages with possible overwrite.
//
//go:nosplit
func MsgOverwriteTrap(sendMsg *MachMsgHeader, option MachMsgOption, sendSize, receiveLimit MachMsgSize, receiveName MachPort, timeout MachMsgTimeout, notify MachPort, receiveMsg *MachMsgHeader, receiveMsgSize MachMsgSize) int32 {
	return machMsgOverwriteTrap(sendMsg, option, sendSize, receiveLimit, receiveName, timeout, notify, receiveMsg, receiveMsgSize)
}

//go:noescape
//go:nosplit
func threadGetSpecialReplyPort() (ret uint32)

// ThreadGetSpecialReplyPort allocate a special reply port for the calling thread.
//
// Returns the
//
//	mach_port_name_t
//
// send right & receive right for special reply port.
//
//	MACH_PORT_NULL
//
// if there are any resource failures.
//
// or other errors.
//
//go:nosplit
func ThreadGetSpecialReplyPort() (ret uint32) {
	return threadGetSpecialReplyPort()
}

//go:noescape
//go:nosplit
func pfzExit() (ret sys.KernReturn)

// PfzExit called from commpage to take a delayed preemption when exiting
// the "Preemption Free Zone" (PFZ).
//
//go:nosplit
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
//go:nosplit
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
//go:nosplit
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

// /* 10 */ "_kernelrpc_mach_vm_allocate_trap",
// /* 11 */ "_kernelrpc_mach_vm_purgable_control_trap",
// /* 12 */ "_kernelrpc_mach_vm_deallocate_trap",
// /* 13 */ "task_dyld_process_info_notify_get_trap",
// /* 14 */ "_kernelrpc_mach_vm_protect_trap",
// /* 15 */ "_kernelrpc_mach_vm_map_trap",
// /* 16 */ "_kernelrpc_mach_port_allocate_trap",
// /* 18 */ "_kernelrpc_mach_port_deallocate_trap",
// /* 19 */ "_kernelrpc_mach_port_mod_refs_trap",
// /* 20 */ "_kernelrpc_mach_port_move_member_trap",
// /* 21 */ "_kernelrpc_mach_port_insert_right_trap",
// /* 22 */ "_kernelrpc_mach_port_insert_member_trap",
// /* 23 */ "_kernelrpc_mach_port_extract_member_trap",
// /* 24 */ "_kernelrpc_mach_port_construct_trap",
// /* 25 */ "_kernelrpc_mach_port_destruct_trap",
// /* 26 */ "mach_reply_port",
// /* 27 */ "thread_self_trap",
// /* 28 */ "task_self_trap",
// /* 29 */ "host_self_trap",
// /* 31 */ "mach_msg_trap",
// /* 32 */ "mach_msg_overwrite_trap",
// /* 33 */ "semaphore_signal_trap",
// /* 34 */ "semaphore_signal_all_trap",
// /* 35 */ "semaphore_signal_thread_trap",
// /* 36 */ "semaphore_wait_trap",
// /* 37 */ "semaphore_wait_signal_trap",
// /* 38 */ "semaphore_timedwait_trap",
// /* 39 */ "semaphore_timedwait_signal_trap",
// /* 40 */ "_kernelrpc_mach_port_get_attributes_trap",
// /* 41 */ "_kernelrpc_mach_port_guard_trap",
// /* 42 */ "_kernelrpc_mach_port_unguard_trap",
// /* 43 */ "mach_generate_activity_id",
// /* 44 */ "task_name_for_pid",
// /* 45 */ "task_for_pid",
// /* 46 */ "pid_for_task",
// /* 48 */ "macx_swapon",
// /* 49 */ "macx_swapoff",
// /* 50 */ "thread_get_special_reply_port",
// /* 51 */ "macx_triggers",
// /* 52 */ "macx_backing_store_suspend",
// /* 53 */ "macx_backing_store_recovery",
// /* 58 */ "pfz_exit",
// /* 59 */ "swtch_pri",
// /* 60 */ "swtch",
// /* 61 */ "thread_switch",
// /* 62 */ "clock_sleep_trap",
// /* 70 */ "host_create_mach_voucher_trap",
// /* 72 */ "mach_voucher_extract_attr_recipe_trap",
// /* 76 */ "_kernelrpc_mach_port_type_trap",
// /* 77 */ "_kernelrpc_mach_port_request_notification_trap",
// /* 89 */ "mach_timebase_info_trap",
// /* 90 */ "mach_wait_until_trap",
// /* 91 */ "mk_timer_create_trap",
// /* 92 */ "mk_timer_destroy_trap",
// /* 93 */ "mk_timer_arm_trap",
// /* 94 */ "mk_timer_cancel_trap",
// /* 95 */ "mk_timer_arm_leeway_trap",
// /* 96 */ "debug_control_port_for_pid",
// /* 100 */ "iokit_user_client_trap",

// 10: kern_return_t _kernelrpc_mach_vm_allocate_trap(mach_port_name_t target, mach_vm_offset_t * addr, mach_vm_size_t size, int flags)
// 11: kern_return_t _kernelrpc_mach_vm_purgable_control_trap(mach_port_name_t target, mach_vm_offset_t address, vm_purgable_t control, int * state)
// 12: kern_return_t _kernelrpc_mach_vm_deallocate_trap(mach_port_name_t target, mach_vm_address_t address, mach_vm_size_t size)
// 13: kern_return_t task_dyld_process_info_notify_get(mach_port_name_array_t names_addr, natural_t * names_count_addr)
// 14: kern_return_t _kernelrpc_mach_vm_protect_trap(mach_port_name_t target, mach_vm_address_t address, mach_vm_size_t size, boolean_t set_maximum, vm_prot_t new_protection)
// 15: kern_return_t _kernelrpc_mach_vm_map_trap(mach_port_name_t target, mach_vm_offset_t * address, mach_vm_size_t size, mach_vm_offset_t mask, int flags, vm_prot_t cur_protection)
// 16: kern_return_t _kernelrpc_mach_port_allocate_trap(mach_port_name_t target, mach_port_right_t right, mach_port_name_t * name)
// 18: kern_return_t _kernelrpc_mach_port_deallocate_trap(mach_port_name_t target, mach_port_name_t name)
// 19: kern_return_t _kernelrpc_mach_port_mod_refs_trap(mach_port_name_t target, mach_port_name_t name, mach_port_right_t right, mach_port_delta_t delta)
// 20: kern_return_t _kernelrpc_mach_port_move_member_trap(mach_port_name_t target, mach_port_name_t member, mach_port_name_t after)
// 21: kern_return_t _kernelrpc_mach_port_insert_right_trap(mach_port_name_t target, mach_port_name_t name, mach_port_name_t poly, mach_msg_type_name_t polyPoly)
// 22: kern_return_t _kernelrpc_mach_port_insert_member_trap(mach_port_name_t target, mach_port_name_t name, mach_port_name_t pset)
// 23: kern_return_t _kernelrpc_mach_port_extract_member_trap(mach_port_name_t target, mach_port_name_t name, mach_port_name_t pset)
// 24: kern_return_t _kernelrpc_mach_port_construct_trap(mach_port_name_t target, mach_port_options_t * options, uint64_t context, mach_port_name_t * name)
// 25: kern_return_t _kernelrpc_mach_port_destruct_trap(mach_port_name_t target, mach_port_name_t name, mach_port_delta_t srdelta, uint64_t guard)
// 32: mach_msg_return_t mach_msg_overwrite_trap(mach_msg_header_t * msg, mach_msg_option_t option, mach_msg_size_t send_size, mach_msg_size_t rcv_size, mach_port_name_t rcv_name, mach_msg_timeout_t timeout, mach_msg_priority_t priority, mach_msg_header_t * rcv_msg, mach_msg_size_t rcv_limit)
// 33: kern_return_t semaphore_signal_trap(mach_port_name_t signal_name)
// 34: kern_return_t semaphore_signal_all_trap(mach_port_name_t signal_name)
// 35: kern_return_t semaphore_signal_thread_trap(mach_port_name_t signal_name, mach_port_name_t thread_name)
// 36: kern_return_t semaphore_wait_trap(mach_port_name_t wait_name)
// 37: kern_return_t semaphore_wait_signal_trap(mach_port_name_t wait_name, mach_port_name_t signal_name)
// 38: kern_return_t semaphore_timedwait_trap(mach_port_name_t wait_name, unsigned int sec, clock_res_t nsec)
// 39: kern_return_t semaphore_timedwait_signal_trap(mach_port_name_t wait_name, mach_port_name_t signal_name, unsigned int sec, clock_res_t nsec)
// 40: kern_return_t _kernelrpc_mach_port_get_attributes_trap(mach_port_name_t target, mach_port_name_t name, mach_port_flavor_t flavor, mach_port_info_t port_info_out, mach_msg_type_number_t * port_info_outCnt)
// 41: kern_return_t _kernelrpc_mach_port_guard_trap(mach_port_name_t target, mach_port_name_t name, uint64_t guard, boolean_t strict)
// 42: kern_return_t _kernelrpc_mach_port_unguard_trap(mach_port_name_t target, mach_port_name_t name, uint64_t guard)
// 43: kern_return_t mach_generate_activity_id(mach_port_name_t target, int count, uint64_t * activity_id)
// 44: kern_return_t task_name_for_pid(mach_port_name_t target_tport, int pid, mach_port_name_t * tn)
// 45: kern_return_t task_for_pid(mach_port_name_t target_tport, int pid, mach_port_name_t * t)
// 46: kern_return_t pid_for_task(mach_port_name_t t, int * x)
// 48: kern_return_t macx_swapon(uint64_t filename, int flags, int size, int priority)
// 49: kern_return_t macx_swapoff(uint64_t filename, int flags)
// 51: kern_return_t macx_triggers(int hi_water, int low_water, int flags, mach_port_t alert_port)
// 52: kern_return_t macx_backing_store_suspend(boolean_t suspend)
// 53: kern_return_t macx_backing_store_recovery(int pid)
// 61: kern_return_t thread_switch(mach_port_name_t thread_name, int option, mach_msg_timeout_t option_time)
// 62: kern_return_t clock_sleep_trap(mach_port_name_t clock_name, sleep_type_t sleep_type, int sleep_sec, int sleep_nsec, mach_timespec_t * wakeup_time)
// 70: kern_return_t host_create_mach_voucher_trap(mach_port_name_t host, mach_voucher_attr_raw_recipe_array_t recipes, int recipes_size, mach_port_name_t * voucher)
// 72: kern_return_t mach_voucher_extract_attr_recipe_trap(mach_port_name_t voucher_name, mach_voucher_attr_key_t key, mach_voucher_attr_raw_recipe_t recipe, mach_msg_type_number_t * recipe_size)
// 76: kern_return_t _kernelrpc_mach_port_type_trap(ipc_space_t task, mach_port_name_t name, mach_port_type_t * ptype)
// 77: kern_return_t _kernelrpc_mach_port_request_notification_trap(ipc_space_t task, mach_port_name_t name, mach_msg_id_t msgid, mach_port_mscount_t sync, mach_port_name_t notify, mach_msg_type_name_t notifyPoly, mach_port_name_t * previous)
// 96: kern_return_t debug_control_port_for_pid(mach_port_name_t target_tport, int pid, mach_port_name_t * t)
// XXX: __uint16_t _OSSwapInt16(__uint16_t _data)
// XXX: __uint32_t _OSSwapInt32(__uint32_t _data)
// XXX: __uint64_t _OSSwapInt64(__uint64_t _data)
// XXX: kern_return_t mach_voucher_deallocate(mach_port_name_t voucher)
// XXX: kern_return_t vm_stats(void * info, unsigned int * count)
// XXX: mach_msg_return_t mach_msg(mach_msg_header_t * msg, mach_msg_option_t option, mach_msg_size_t send_size, mach_msg_size_t rcv_size, mach_port_name_t rcv_name, mach_msg_timeout_t timeout, mach_port_name_t notify)
// XXX: mach_msg_return_t mach_msg_overwrite(mach_msg_header_t * msg, mach_msg_option_t option, mach_msg_size_t send_size, mach_msg_size_t rcv_size, mach_port_name_t rcv_name, mach_msg_timeout_t timeout, mach_port_name_t notify, mach_msg_header_t * rcv_msg, mach_msg_size_t rcv_limit)
