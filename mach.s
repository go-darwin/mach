// Copyright 2019 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,amd64

#include "textflag.h"

// func machReplyPort() uintptr
TEXT ·machReplyPort(SB), NOSPLIT, $8
	MOVL $(0x1000000+26), AX // 26: mach_reply_port
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func threadSelfTrap() uintptr
// func machThreadSelf() uintptr
TEXT ·threadSelfTrap(SB), NOSPLIT, $8
	MOVL $(0x1000000+27), AX // 27: thread_self_trap
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func taskSelfTrap() uintptr
// func machTaskSelf() uintptr
TEXT ·taskSelfTrap(SB), NOSPLIT, $8
	MOVL $(0x1000000+28), AX // 28: task_self_trap
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func hostSelfTrap() uintptr
// func machHostSelf() uintptr
TEXT ·hostSelfTrap(SB), NOSPLIT, $8
	MOVL $(0x1000000+29), AX // 29: host_self_trap
	SYSCALL
	MOVL AX, ret+0(FP)
	RET

// func machMsgTrap(unsafe.Pointer, uint32, uint32, uint32, uint32, uint32, uint32) uint32
TEXT ·mach_msg_trap(SB), NOSPLIT, $8
	MOVQ  8(SP), DI
	MOVL  16(SP), SI
	MOVL  20(SP), DX
	MOVL  24(SP), R10
	MOVL  28(SP), R8
	MOVL  32(SP), R9
	MOVL  36(SP), R11
	PUSHQ R11                 // seventh arg, on stack
	MOVL  $(0x1000000+31), AX // 31: mach_msg_trap
	SYSCALL
	POPQ  R11
	MOVL  AX, ret+0(FP)
	RET
