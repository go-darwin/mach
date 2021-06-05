// SPDX-FileCopyrightText: Copyright 2019 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

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
	MOVQ  msg+0(FP), DI
	MOVL  opt+8(FP), SI
	MOVL  ssize+0(SP), DX
	MOVL  rsize+0(SP), R10
	MOVL  rname+0(SP), R8
	MOVL  to+0(SP), R9
	MOVL  not+0(SP), R11
	PUSHQ R11                 // seventh arg, on stack
	MOVL  $(0x1000000+31), AX // 31: mach_msg_trap
	SYSCALL
	POPQ  R11
	MOVL  AX, ret+32(FP)
	RET
