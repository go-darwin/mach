// Copyright 2019 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

#include "go_asm.h"
#include "textflag.h"

// func machReplyPort() (ret uint32)
TEXT ·machReplyPort(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+26), AX // mach_reply_port
	SYSCALL
	MOVL  AX, ret+0(FP)
	POPQ  BP
	RET

// func threadSelfTrap() (ret uint32)
TEXT ·threadSelfTrap(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+27), AX // thread_self_trap
	SYSCALL
	MOVL  AX, ret+0(FP)
	POPQ  BP
	RET

// func taskSelfTrap() (ret uint32)
TEXT ·taskSelfTrap(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+28), AX // task_self_trap
	SYSCALL
	MOVL  AX, ret+0(FP)
	POPQ  BP
	RET

// func hostSelfTrap() (ret uint32)
TEXT ·hostSelfTrap(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+29), AX // host_self_trap
	SYSCALL
	MOVL  AX, ret+0(FP)
	POPQ  BP
	RET

// func machMsgTrap(msg *MachMsgHeader, option MachMsgOption, sendSize, rcvSize MachMsgSize, rcvName MachPortName, timeout MachMsgTimeout, notify MachPortName) (ret int32)
TEXT ·machMsgTrap(SB), NOSPLIT, $0
	MOVQ  msg+0(FP), DI       // arg 1 msg
	MOVL  option+8(FP), SI    // arg 2 option
	MOVL  sendSize+12(FP), DX // arg 3 sendSize
	MOVL  rcvSize+16(FP), R10 // arg 4 rcvSize
	MOVL  rcvName+20(FP), R8  // arg 5 rcvName
	MOVL  timeout+24(FP), R9  // arg 6 timeout
	MOVL  notify+28(FP), R11  // arg 7 notify
	PUSHQ R11                 // seventh arg, on stack
	MOVL  $(0x1000000+31), AX // mach_msg_trap
	SYSCALL
	POPQ  R11
	MOVL  AX, ret+32(FP)
	RET

// func machMsgOverwriteTrap(sendMsg *MachMsgHeader, option MachMsgOption, sendSize, receiveLimit MachMsgSize, receiveName MachPort, timeout MachMsgTimeout, notify MachPort, receiveMsg *MachMsgHeader, receiveMsgSize MachMsgSize) (ret int32)
TEXT ·machMsgOverwriteTrap(SB), NOSPLIT, $0
	MOVQ  msg+0(FP), DI            // arg 1 msg
	MOVL  option+8(FP), SI         // arg 2 option
	MOVL  sendSize+12(FP), DX      // arg 3 sendSize
	MOVL  receiveLimit+16(FP), R10 // arg 4 rcvSize
	MOVL  receiveName+20(FP), R8   // arg 5 rcvName
	MOVL  timeout+24(FP), R9       // arg 6 timeout
	MOVL  notify+28(FP), R11       // arg 7 notify
	PUSHQ R11                      // seventh arg, on stack
	MOVL  $(0x1000000+32), AX      // mach_msg_overwrite
	SYSCALL
	POPQ  R11
	MOVL  AX, ret+32(FP)
	RET

// func threadGetSpecialReplyPort() (ret uint32)
TEXT ·threadGetSpecialReplyPort(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+50), AX // thread_get_special_reply_port
	SYSCALL
	MOVL  AX, ret+0(FP)
	POPQ  BP
	RET

// func pfzExit() (ret sys.KernReturn)
TEXT ·pfzExit(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+58), AX // pfz_exit
	SYSCALL
	MOVL  AX, ret+0(FP)
	POPQ  BP
	RET

// func swtchPri() (ret bool)
TEXT ·swtchPri(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+59), AX // swtch_pri
	SYSCALL
	SETEQ ret+0(FP)
	POPQ  BP
	RET

// func swtch() (ret bool)
TEXT ·swtch(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+60), AX // swtch
	SYSCALL
	SETEQ ret+0(FP)
	POPQ  BP
	RET

// func mkTimerCreateTrap() (ret uint32)
TEXT ·mkTimerCreateTrap(SB), NOSPLIT, $0
	PUSHQ BP                  // make a frame, keep stack aligned
	MOVQ  SP, BP
	MOVL  $(0x1000000+91), AX // mk_timer_create_trap
	SYSCALL
	MOVL  AX, ret+0(FP)
	POPQ  BP
	RET
