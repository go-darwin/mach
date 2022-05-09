// Copyright 2022 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

package mach

import "github.com/go-darwin/sys"

//go:cgo_import_dynamic mach_vm_allocate mach_vm_allocate "/usr/lib/libSystem.B.dylib"

var mach_vm_allocate_trampoline_addr uintptr

// MachVMAllocate allocate a region of virtual memory.
func MachVMAllocate(targetTask VMMap, address MachVMAddress, size MachVMSize, flags int) error {
	_, _, errno := sys.Ccall6(mach_vm_allocate_trampoline_addr, uintptr(targetTask), uintptr(address), uintptr(size), uintptr(flags), 0, 0)
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}

//go:cgo_import_dynamic mach_vm_deallocate mach_vm_deallocate "/usr/lib/libSystem.B.dylib"

var mach_vm_deallocate_trampoline_addr uintptr

// MachVMDeallocate deallocate a region of virtual memory.
func MachVMDeallocate(target VMMap, address MachVMAddress, size MachVMSize) error {
	_, _, errno := sys.Ccall(mach_vm_allocate_trampoline_addr, uintptr(target), uintptr(address), uintptr(size))
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}
