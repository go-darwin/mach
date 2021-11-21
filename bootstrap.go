// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

// This functions corresponds to "bootstrap.h".

package mach

import (
	"unsafe"

	"github.com/go-darwin/sys"
)

//go:cgo_import_dynamic bootstrap_check_in bootstrap_check_in "/usr/lib/libSystem.B.dylib"

var bootstrap_check_in_trampoline_addr uintptr

func BootstrapCheckIn(bp MachPort, serviceName string, sp MachPort) error {
	_, _, errno := sys.Syscall(bootstrap_check_in_trampoline_addr, uintptr(bp), uintptr(unsafe.Pointer(&serviceName)), uintptr(sp))
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic bootstrap_create_server bootstrap_create_server "/usr/lib/libSystem.B.dylib"

var bootstrap_create_server_trampoline_addr uintptr

func BootstrapCreateServer(bp MachPort, serverCmd string, serverUID uint32, onDemand bool, serverPort MachPort) error {
	_, _, errno := sys.Syscall9(bootstrap_create_server_trampoline_addr, uintptr(bp), uintptr(unsafe.Pointer(&serverCmd)), uintptr(serverUID), uintptr(unsafe.Pointer(&onDemand)), uintptr(serverPort), 0, 0, 0, 0)
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic bootstrap_create_service bootstrap_create_service "/usr/lib/libSystem.B.dylib"

var bootstrap_create_service_trampoline_addr uintptr

func BootstrapCreateService(bp MachPort, serviceName string, sp MachPort) error {
	_, _, errno := sys.Syscall(bootstrap_create_service_trampoline_addr, uintptr(bp), uintptr(unsafe.Pointer(&serviceName)), uintptr(sp))
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic bootstrap_look_up bootstrap_look_up "/usr/lib/libSystem.B.dylib"

var bootstrap_look_up_trampoline_addr uintptr

func BootstrapLookUp(bp MachPort, serviceName string, sp MachPort) error {
	_, _, errno := sys.Syscall(bootstrap_look_up_trampoline_addr, uintptr(bp), uintptr(unsafe.Pointer(&serviceName)), uintptr(sp))
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic bootstrap_parent bootstrap_parent "/usr/lib/libSystem.B.dylib"

var bootstrap_parent_trampoline_addr uintptr

func BootstrapParent(bp MachPort, parentPort MachPort) error {
	_, _, errno := sys.Syscall(bootstrap_parent_trampoline_addr, uintptr(bp), uintptr(parentPort), 0)
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic bootstrap_register bootstrap_register "/usr/lib/libSystem.B.dylib"

var bootstrap_register_trampoline_addr uintptr

func BootstrapRegister(bp MachPort, serviceName string, sp MachPort) error {
	_, _, errno := sys.Syscall(bootstrap_register_trampoline_addr, uintptr(bp), uintptr(unsafe.Pointer(&serviceName)), uintptr(sp))
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic bootstrap_status bootstrap_status "/usr/lib/libSystem.B.dylib"

var bootstrap_status_trampoline_addr uintptr

func BootstrapStatus(bp MachPort, serviceName string, serviceActive BootstrapStatusT) error {
	_, _, errno := sys.Syscall(bootstrap_status_trampoline_addr, uintptr(bp), uintptr(unsafe.Pointer(&serviceName)), uintptr(serviceActive))
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic bootstrap_strerror bootstrap_strerror "/usr/lib/libSystem.B.dylib"

var bootstrap_strerror_trampoline_addr uintptr

func BootstrapStrerror(r sys.KernReturn) string {
	s, _, errno := sys.Syscall(bootstrap_strerror_trampoline_addr, uintptr(r), 0, 0)
	if errno != 0 {
		return ""
	}

	return sys.GoString((*byte)(unsafe.Pointer(&s)))
}

//go:cgo_import_dynamic bootstrap_subset bootstrap_subset "/usr/lib/libSystem.B.dylib"

var bootstrap_subset_trampoline_addr uintptr

func BootstrapSubset(bp MachPort, requestorPort MachPort, subsetPort MachPort) error {
	_, _, errno := sys.Syscall(bootstrap_subset_trampoline_addr, uintptr(bp), uintptr(requestorPort), uintptr(subsetPort))
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic bootstrap_unprivileged bootstrap_unprivileged "/usr/lib/libSystem.B.dylib"

var bootstrap_unprivileged_trampoline_addr uintptr

func BootstrapUnprivileged(bp MachPort, unprivPort MachPort) error {
	_, _, errno := sys.Syscall(bootstrap_unprivileged_trampoline_addr, uintptr(bp), uintptr(unprivPort), 0)
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}
