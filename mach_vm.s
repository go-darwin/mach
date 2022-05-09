// Copyright 2022 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

#include "textflag.h"

GLOBL ·mach_vm_allocate_trampoline_addr(SB), RODATA, $8
DATA ·mach_vm_allocate_trampoline_addr(SB)/8, $mach_vm_allocate_trampoline<>(SB)

TEXT mach_vm_allocate_trampoline<>(SB), NOSPLIT, $0-0
	JMP mach_vm_allocate(SB)

GLOBL ·mach_vm_deallocate_trampoline_addr(SB), RODATA, $8
DATA ·mach_vm_deallocate_trampoline_addr(SB)/8, $mach_vm_deallocate_trampoline<>(SB)

TEXT mach_vm_deallocate_trampoline<>(SB), NOSPLIT, $0-0
	JMP mach_vm_deallocate(SB)
