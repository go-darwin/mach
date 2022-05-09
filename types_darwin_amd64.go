// Copyright 2019 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

// copied and edit from https://github.com/apple-opensource/xnu/blob/7195.101.1/osfmk/mach/mach_traps.h
package mach

/*
#cgo CFLAGS: -mmacosx-version-min=10.6

#include <stdint.h>

#include <mach/mach.h>

#include <machine/endian.h>

#include <sys/cdefs.h>

// https://github.com/apple-oss-distributions/xnu/blob/xnu-8020.101.4/osfmk/mach/mach_traps.h#L363-L877

// Syscall data translations routines
//
// The kernel may support multiple userspace ABIs, and must use
// argument structures with elements large enough for any of them.

#if CONFIG_REQUIRES_U32_MUNGING
#define PAD_(t) (sizeof(uint64_t) <= sizeof(t) \
	        ? 0 : sizeof(uint64_t) - sizeof(t))
#define PAD_ARG_8
#else
#define PAD_(t) (sizeof(uint32_t) <= sizeof(t) \
	        ? 0 : sizeof(uint32_t) - sizeof(t))
#if __arm__ && (__BIGGEST_ALIGNMENT__ > 4)
#define PAD_ARG_8
#else
#define PAD_ARG_8 char arg8_pad_[sizeof(uint32_t)];
#endif
#endif

#if BYTE_ORDER == LITTLE_ENDIAN
#define PADL_(t)        0
#define PADR_(t)        PAD_(t)
#else
#define PADL_(t)        PAD_(t)
#define PADR_(t)        0
#endif

#define PAD_ARG_(arg_type, arg_name) \
  char arg_name##_l_[PADL_(arg_type)]; arg_type arg_name; char arg_name##_r_[PADR_(arg_type)]

// To support 32-bit clients as well as 64-bit clients, argument
// structures may need to be munged to repack the arguments. All
// active architectures do this inline in the code to dispatch Mach
// traps, without calling out to the BSD system call mungers.

struct kern_invalid_args {
	int32_t dummy;
};

struct mach_reply_port_args {
	int32_t dummy;
};

struct thread_get_special_reply_port_args {
	int32_t dummy;
};

struct thread_self_trap_args {
	int32_t dummy;
};

struct task_self_trap_args {
	int32_t dummy;
};

struct host_self_trap_args {
	int32_t dummy;
};

struct mach_msg_overwrite_trap_args {
	PAD_ARG_(user_addr_t, msg);
	PAD_ARG_(mach_msg_option_t, option);
	PAD_ARG_(mach_msg_size_t, send_size);
	PAD_ARG_(mach_msg_size_t, rcv_size);
	PAD_ARG_(mach_port_name_t, rcv_name);
	PAD_ARG_(mach_msg_timeout_t, timeout);
	PAD_ARG_(mach_msg_priority_t, priority);
	PAD_ARG_8
	    PAD_ARG_(user_addr_t, rcv_msg); // Unused on mach_msg_trap
};

struct semaphore_signal_trap_args {
	PAD_ARG_(mach_port_name_t, signal_name);
};

struct semaphore_signal_all_trap_args {
	PAD_ARG_(mach_port_name_t, signal_name);
};

struct semaphore_signal_thread_trap_args {
	PAD_ARG_(mach_port_name_t, signal_name);
	PAD_ARG_(mach_port_name_t, thread_name);
};

struct semaphore_wait_trap_args {
	PAD_ARG_(mach_port_name_t, wait_name);
};

struct semaphore_wait_signal_trap_args {
	PAD_ARG_(mach_port_name_t, wait_name);
	PAD_ARG_(mach_port_name_t, signal_name);
};

struct semaphore_timedwait_trap_args {
	PAD_ARG_(mach_port_name_t, wait_name);
	PAD_ARG_(unsigned int, sec);
	PAD_ARG_(clock_res_t, nsec);
};

struct semaphore_timedwait_signal_trap_args {
	PAD_ARG_(mach_port_name_t, wait_name);
	PAD_ARG_(mach_port_name_t, signal_name);
	PAD_ARG_(unsigned int, sec);
	PAD_ARG_(clock_res_t, nsec);
};

struct task_for_pid_args {
	PAD_ARG_(mach_port_name_t, target_tport);
	PAD_ARG_(int, pid);
	PAD_ARG_(user_addr_t, t);
};

struct task_name_for_pid_args {
	PAD_ARG_(mach_port_name_t, target_tport);
	PAD_ARG_(int, pid);
	PAD_ARG_(user_addr_t, t);
};

struct pid_for_task_args {
	PAD_ARG_(mach_port_name_t, t);
	PAD_ARG_(user_addr_t, pid);
};

struct debug_control_port_for_pid_args {
	PAD_ARG_(mach_port_name_t, target_tport);
	PAD_ARG_(int, pid);
	PAD_ARG_(user_addr_t, t);
};

struct macx_swapon_args {
	PAD_ARG_(uint64_t, filename);
	PAD_ARG_(int, flags);
	PAD_ARG_(int, size);
	PAD_ARG_(int, priority);
};

struct macx_swapoff_args {
	PAD_ARG_(uint64_t, filename);
	PAD_ARG_(int, flags);
};

struct macx_triggers_args {
	PAD_ARG_(int, hi_water);
	PAD_ARG_(int, low_water);
	PAD_ARG_(int, flags);
	PAD_ARG_(mach_port_t, alert_port);
};

struct macx_backing_store_suspend_args {
	PAD_ARG_(boolean_t, suspend);
};

struct macx_backing_store_recovery_args {
	PAD_ARG_(int, pid);
};

struct swtch_pri_args {
	PAD_ARG_(int, pri);
};

struct pfz_exit_args {
	int32_t dummy;
};

struct swtch_args {
	int32_t dummy;
};

struct clock_sleep_trap_args {
	PAD_ARG_(mach_port_name_t, clock_name);
	PAD_ARG_(sleep_type_t, sleep_type);
	PAD_ARG_(int, sleep_sec);
	PAD_ARG_(int, sleep_nsec);
	PAD_ARG_(user_addr_t, wakeup_time);
};

struct thread_switch_args {
	PAD_ARG_(mach_port_name_t, thread_name);
	PAD_ARG_(int, option);
	PAD_ARG_(mach_msg_timeout_t, option_time);
};

struct mach_timebase_info_trap_args {
	PAD_ARG_(user_addr_t, info);
};

struct mach_wait_until_trap_args {
	PAD_ARG_(uint64_t, deadline);
};

struct mk_timer_create_trap_args {
	int32_t dummy;
};

struct mk_timer_destroy_trap_args {
	PAD_ARG_(mach_port_name_t, name);
};

struct mk_timer_arm_trap_args {
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(uint64_t, expire_time);
};

struct mk_timer_arm_leeway_trap_args {
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(uint64_t, mk_timer_flags);
	PAD_ARG_(uint64_t, expire_time);
	PAD_ARG_(uint64_t, mk_leeway);
};

struct mk_timer_cancel_trap_args {
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(user_addr_t, result_time);
};

struct _kernelrpc_mach_vm_allocate_trap_args {
	PAD_ARG_(mach_port_name_t, target);     // 1 word
	PAD_ARG_(user_addr_t, addr);            // 1 word
	PAD_ARG_(mach_vm_size_t, size);         // 2 words
	PAD_ARG_(int, flags);                   // 1 word
};                                              // Total: 5

struct _kernelrpc_mach_vm_deallocate_args {
	PAD_ARG_(mach_port_name_t, target);     // 1 word
	PAD_ARG_(mach_vm_address_t, address);   // 2 words
	PAD_ARG_(mach_vm_size_t, size);         // 2 words
};                                              // Total: 5

struct task_dyld_process_info_notify_get_trap_args {
	PAD_ARG_(mach_vm_address_t, names_addr);     // 2 words
	PAD_ARG_(mach_vm_address_t, names_count_addr);  // 2 words
};                                               // Total: 4

struct _kernelrpc_mach_vm_protect_args {
	PAD_ARG_(mach_port_name_t, target);     // 1 word
	PAD_ARG_(mach_vm_address_t, address);   // 2 words
	PAD_ARG_(mach_vm_size_t, size);         // 2 words
	PAD_ARG_(boolean_t, set_maximum);       // 1 word
	PAD_ARG_(vm_prot_t, new_protection);    // 1 word
};                                              // Total: 7

struct _kernelrpc_mach_vm_map_trap_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(user_addr_t, addr);
	PAD_ARG_(mach_vm_size_t, size);
	PAD_ARG_(mach_vm_offset_t, mask);
	PAD_ARG_(int, flags);
	PAD_ARG_8
	    PAD_ARG_(vm_prot_t, cur_protection);
};

struct _kernelrpc_mach_vm_purgable_control_trap_args {
	PAD_ARG_(mach_port_name_t, target);     // 1 word
	PAD_ARG_(mach_vm_offset_t, address);    // 2 words
	PAD_ARG_(vm_purgable_t, control);       // 1 word
	PAD_ARG_(user_addr_t, state);           // 1 word
};                                              // Total: 5

struct _kernelrpc_mach_port_allocate_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_right_t, right);
	PAD_ARG_(user_addr_t, name);
};

struct _kernelrpc_mach_port_deallocate_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
};

struct _kernelrpc_mach_port_mod_refs_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(mach_port_right_t, right);
	PAD_ARG_(mach_port_delta_t, delta);
};

struct _kernelrpc_mach_port_move_member_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, member);
	PAD_ARG_(mach_port_name_t, after);
};

struct _kernelrpc_mach_port_insert_right_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(mach_port_name_t, poly);
	PAD_ARG_(mach_msg_type_name_t, polyPoly);
};

struct _kernelrpc_mach_port_get_attributes_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(mach_port_flavor_t, flavor);
	PAD_ARG_(user_addr_t, info);
	PAD_ARG_(user_addr_t, count);
};

struct _kernelrpc_mach_port_insert_member_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(mach_port_name_t, pset);
};

struct _kernelrpc_mach_port_extract_member_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(mach_port_name_t, pset);
};

struct _kernelrpc_mach_port_construct_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(user_addr_t, options);
	PAD_ARG_(uint64_t, context);
	PAD_ARG_(user_addr_t, name);
};

struct _kernelrpc_mach_port_destruct_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(mach_port_delta_t, srdelta);
	PAD_ARG_(uint64_t, guard);
};

struct _kernelrpc_mach_port_guard_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(uint64_t, guard);
	PAD_ARG_(boolean_t, strict);
};

struct _kernelrpc_mach_port_unguard_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(uint64_t, guard);
};

struct mach_generate_activity_id_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(int, count);
	PAD_ARG_(user_addr_t, activity_id);
};

// Voucher trap interfaces

struct host_create_mach_voucher_args {
	PAD_ARG_(mach_port_name_t, host);
	PAD_ARG_(mach_voucher_attr_raw_recipe_array_t, recipes);
	PAD_ARG_(int, recipes_size);
	PAD_ARG_(user_addr_t, voucher);
};

struct mach_voucher_extract_attr_recipe_args {
	PAD_ARG_(mach_port_name_t, voucher_name);
	PAD_ARG_(mach_voucher_attr_key_t, key);
	PAD_ARG_(mach_voucher_attr_raw_recipe_t, recipe);
	PAD_ARG_(user_addr_t, recipe_size);
};

struct _kernelrpc_mach_port_type_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_right_t, name);
	PAD_ARG_(user_addr_t, ptype);
};

struct _kernelrpc_mach_port_request_notification_args {
	PAD_ARG_(mach_port_name_t, target);
	PAD_ARG_(mach_port_name_t, name);
	PAD_ARG_(mach_msg_id_t, msgid);
	PAD_ARG_(mach_port_mscount_t, sync);
	PAD_ARG_(mach_port_name_t, notify);
	PAD_ARG_(mach_msg_type_name_t, notifyPoly);
	PAD_ARG_(user_addr_t, previous);
};

// not published to LP64 clients yet
struct iokit_user_client_trap_args {
	PAD_ARG_(void *, userClientRef);
	PAD_ARG_(uint32_t, index);
	PAD_ARG_(void *, p1);
	PAD_ARG_(void *, p2);
	PAD_ARG_(void *, p3);
	PAD_ARG_(void *, p4);
	PAD_ARG_(void *, p5);
	PAD_ARG_8
	    PAD_ARG_(void *, p6);
};
*/
import "C"

type MachPortName = C.mach_port_name_t

type MachMsgReturn = C.mach_msg_return_t

type MachMsgHeader = C.mach_msg_header_t

type MachMsgOption = C.mach_msg_option_t

type MachMsgSize = C.mach_msg_size_t

type MachMsgTimeout = C.mach_msg_timeout_t

type MachMsgPriority = C.mach_msg_priority_t

type ClockRes = C.clock_res_t

type SleepType = C.sleep_type_t

type MachTimespec = C.mach_timespec_t

type MachVMOffset = C.mach_vm_offset_t

type MachVMSize = C.mach_vm_size_t

type MachPortNameArray = C.mach_port_name_array_t

type Natural = C.natural_t

type Boolean = C.boolean_t

type VMProt = C.vm_prot_t

type VMPurgable = C.vm_purgable_t

type MachPortRight = C.mach_port_right_t

type MachPortDelta = C.mach_port_delta_t

type MachMsgTypeName = C.mach_msg_type_name_t

type MachPortFlavor = C.mach_port_flavor_t

type MachMsgTypeNumber = C.mach_msg_type_number_t

type MachPort = C.mach_port_t

type MachVoucherAttrRawRecipeArray = C.mach_voucher_attr_raw_recipe_array_t

type MachVoucherAttrKey = C.mach_voucher_attr_key_t

type MachVoucherAttrRawRecipe = C.mach_voucher_attr_raw_recipe_t

type IPCSpace = C.ipc_space_t

type KernInvalidArgs = C.struct_kern_invalid_args

type MachReplyPortArgs = C.struct_mach_reply_port_args

type ThreadGetSpecialReplyPortArgs = C.struct_thread_get_special_reply_port_args

type ThreadSelfTrapArgs = C.struct_thread_self_trap_args

type TaskSelfTrapArgs = C.struct_task_self_trap_args

type HostSelfTrapArgs = C.struct_host_self_trap_args

type MachMsgOverwriteTrapArgs = C.struct_mach_msg_overwrite_trap_args

type SemaphoreSignalTrapArgs = C.struct_semaphore_signal_trap_args

type SemaphoreSignalAllTrapArgs = C.struct_semaphore_signal_all_trap_args

type SemaphoreSignalThreadTrapArgs = C.struct_semaphore_signal_thread_trap_args

type SemaphoreWaitTrapArgs = C.struct_semaphore_wait_trap_args

type SemaphoreWaitSignalTrapArgs = C.struct_semaphore_wait_signal_trap_args

type SemaphoreTimedwaitTrapArgs = C.struct_semaphore_timedwait_trap_args

type SemaphoreTimedwaitSignalTrapArgs = C.struct_semaphore_timedwait_signal_trap_args

type TaskForPidArgs = C.struct_task_for_pid_args

type TaskNameForPidArgs = C.struct_task_name_for_pid_args

type PidForTaskArgs = C.struct_pid_for_task_args

type DebugControlPortForPidArgs = C.struct_debug_control_port_for_pid_args

type MacxSwaponArgs = C.struct_macx_swapon_args

type MacxSwapoffArgs = C.struct_macx_swapoff_args

type MacxTriggersArgs = C.struct_macx_triggers_args

type MacxBackingStoreSuspendArgs = C.struct_macx_backing_store_suspend_args

type MacxBackingStoreRecoveryArgs = C.struct_macx_backing_store_recovery_args

type SwtchPriArgs = C.struct_swtch_pri_args

type PfzExitArgs = C.struct_pfz_exit_args

type SwtchArgs = C.struct_swtch_args

type ClockSleepTrapArgs = C.struct_clock_sleep_trap_args

type ThreadSwitchArgs = C.struct_thread_switch_args

type MachTimebaseInfoTrapArgs = C.struct_mach_timebase_info_trap_args

type MachWaitUntilTrapArgs = C.struct_mach_wait_until_trap_args

type MkTimerCreateTrapArgs = C.struct_mk_timer_create_trap_args

type MkTimerDestroyTrapArgs = C.struct_mk_timer_destroy_trap_args

type MkTimerArmTrapArgs = C.struct_mk_timer_arm_trap_args

type MkTimerArmLeewayTrapArgs = C.struct_mk_timer_arm_leeway_trap_args

type MkTimerCancelTrapArgs = C.struct_mk_timer_cancel_trap_args

type KernelrpcMachVmAllocateTrapArgs = C.struct__kernelrpc_mach_vm_allocate_trap_args

type KernelrpcMachVmDeallocateArgs = C.struct__kernelrpc_mach_vm_deallocate_args

type TaskDyldProcessInfoNotifyGetTrapArgs = C.struct_task_dyld_process_info_notify_get_trap_args

type KernelrpcMachVmProtectArgs = C.struct__kernelrpc_mach_vm_protect_args

type KernelrpcMachVmMapTrapArgs = C.struct__kernelrpc_mach_vm_map_trap_args

type KernelrpcMachVmPurgableControlTrapArgs = C.struct__kernelrpc_mach_vm_purgable_control_trap_args

type KernelrpcMachPortAllocateArgs = C.struct__kernelrpc_mach_port_allocate_args

type KernelrpcMachPortDeallocateArgs = C.struct__kernelrpc_mach_port_deallocate_args

type KernelrpcMachPortModRefsArgs = C.struct__kernelrpc_mach_port_mod_refs_args

type KernelrpcMachPortMoveMemberArgs = C.struct__kernelrpc_mach_port_move_member_args

type KernelrpcMachPortInsertRightArgs = C.struct__kernelrpc_mach_port_insert_right_args

type KernelrpcMachPortGetAttributesArgs = C.struct__kernelrpc_mach_port_get_attributes_args

type KernelrpcMachPortInsertMemberArgs = C.struct__kernelrpc_mach_port_insert_member_args

type KernelrpcMachPortExtractMemberArgs = C.struct__kernelrpc_mach_port_extract_member_args

type KernelrpcMachPortConstructArgs = C.struct__kernelrpc_mach_port_construct_args

type KernelrpcMachPortDestructArgs = C.struct__kernelrpc_mach_port_destruct_args

type KernelrpcMachPortGuardArgs = C.struct__kernelrpc_mach_port_guard_args

type KernelrpcMachPortUnguardArgs = C.struct__kernelrpc_mach_port_unguard_args

type MachGenerateActivityIdArgs = C.struct_mach_generate_activity_id_args

//
// Voucher trap interfaces
//

type HostCreateMachVoucherArgs = C.struct_host_create_mach_voucher_args

type MachVoucherExtractAttrRecipeArgs = C.struct_mach_voucher_extract_attr_recipe_args

type KernelrpcMachPortTypeArgs = C.struct__kernelrpc_mach_port_type_args

type KernelrpcMachPortRequestNotificationArgs = C.struct__kernelrpc_mach_port_request_notification_args

// not published to LP64 clients yet
type IokitUserClientTrapArgs = C.struct_iokit_user_client_trap_args

type ErrorType = C.mach_error_t
