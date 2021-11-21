// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

#include "textflag.h"

GLOBL ·bootstrap_check_in_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_check_in_trampoline_addr(SB)/8, $bootstrap_check_in_trampoline<>(SB)

TEXT bootstrap_check_in_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_check_in(SB)

GLOBL ·bootstrap_create_server_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_create_server_trampoline_addr(SB)/8, $bootstrap_create_server_trampoline<>(SB)

TEXT bootstrap_create_server_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_create_server(SB)

GLOBL ·bootstrap_create_service_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_create_service_trampoline_addr(SB)/8, $bootstrap_create_service_trampoline<>(SB)

TEXT bootstrap_create_service_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_create_service(SB)

GLOBL ·bootstrap_look_up_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_look_up_trampoline_addr(SB)/8, $bootstrap_look_up_trampoline<>(SB)

TEXT bootstrap_look_up_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_look_up(SB)

GLOBL ·bootstrap_parent_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_parent_trampoline_addr(SB)/8, $bootstrap_parent_trampoline<>(SB)

TEXT bootstrap_parent_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_parent(SB)

GLOBL ·bootstrap_register_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_register_trampoline_addr(SB)/8, $bootstrap_register_trampoline<>(SB)

TEXT bootstrap_register_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_register(SB)

GLOBL ·bootstrap_status_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_status_trampoline_addr(SB)/8, $bootstrap_status_trampoline<>(SB)

TEXT bootstrap_status_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_status(SB)

GLOBL ·bootstrap_strerror_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_strerror_trampoline_addr(SB)/8, $bootstrap_strerror_trampoline<>(SB)

TEXT bootstrap_strerror_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_strerror(SB)

GLOBL ·bootstrap_subset_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_subset_trampoline_addr(SB)/8, $bootstrap_subset_trampoline<>(SB)

TEXT bootstrap_subset_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_subset(SB)

GLOBL ·bootstrap_unprivileged_trampoline_addr(SB), RODATA, $8
DATA ·bootstrap_unprivileged_trampoline_addr(SB)/8, $bootstrap_unprivileged_trampoline<>(SB)

TEXT bootstrap_unprivileged_trampoline<>(SB), NOSPLIT, $0-0
	JMP bootstrap_unprivileged(SB)
