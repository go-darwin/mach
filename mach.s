// Copyright 2019 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

#include "textflag.h"

// func machReplyPort() (ret uint32)
TEXT ·machReplyPort(SB), NOSPLIT, $8
	MOVL $(0x1000000+26), AX // 26: mach_reply_port
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func threadSelfTrap() (ret uint32)
TEXT ·threadSelfTrap(SB), NOSPLIT, $8
	MOVL $(0x1000000+27), AX // 27: thread_self_trap
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func taskSelfTrap() (ret uint32)
TEXT ·taskSelfTrap(SB), NOSPLIT, $8
	MOVL $(0x1000000+28), AX // 28: task_self_trap
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func hostSelfTrap() (ret uint32)
TEXT ·hostSelfTrap(SB), NOSPLIT, $8
	MOVL $(0x1000000+29), AX // 29: host_self_trap
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func machMsgTrap(msg unsafe.Pointer, opt, ssize, rsize, rname, to, not uint32) (ret uint32)
TEXT ·machMsgTrap(SB), NOSPLIT, $40
	MOVL  opt+8(FP), SI       // arg 2 opt
	MOVL  ssize+0(SP), DX     // arg 3 ssize
	MOVL  rsize+0(SP), R10    // arg 4 rsize
	MOVL  rname+0(SP), R8     // arg 5 rname
	MOVL  to+0(SP), R9        // arg 6 to
	MOVL  not+0(SP), R11      // arg 7 not
	PUSHQ R11                 // seventh arg, on stack
	MOVQ  msg+0(FP), DI       // arg 1 msg
	MOVL  $(0x1000000+31), AX // 31: mach_msg_trap
	SYSCALL
	POPQ  R11
	MOVL  AX, ret+32(FP)
	RET

// func threadGetSpecialReplyPort() (ret MachPortName)
TEXT ·threadGetSpecialReplyPort(SB), NOSPLIT, $8
	MOVL $(0x1000000+50), AX // 50: thread_get_special_reply_port
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func pfzExit() (ret sys.KernReturn)
TEXT ·pfzExit(SB), NOSPLIT, $8
	MOVL $(0x1000000+58), AX // 58: pfz_exit
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func swtchPri() (ret bool)
TEXT ·swtchPri(SB), NOSPLIT, $8
	MOVL $(0x1000000+59), AX // 59: swtch_pri
	SYSCALL
	SETEQ ret+0(FP)
	RET

// func swtch() (ret bool)
TEXT ·swtch(SB), NOSPLIT, $8
	MOVL $(0x1000000+60), AX // 60: swtch
	SYSCALL
	SETEQ ret+0(FP)
	RET

// func mkTimerCreateTrap() (ret MachPortName)
TEXT ·mkTimerCreateTrap(SB), NOSPLIT, $8
	MOVL $(0x1000000+91), AX // 91: mk_timer_create_trap
	SYSCALL
	MOVL AX, ret+0(FP)
	RET
