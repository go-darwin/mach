// Copyright 2022 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

package mach

import (
	"unsafe"

	"github.com/go-darwin/sys"
)

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

//go:cgo_import_dynamic mach_make_memory_entry _mach_make_memory_entry "/usr/lib/libSystem.B.dylib"

var mach_make_memory_entry_trampoline_addr uintptr

// MachMakeMemoryEntry allow pagers to create named entries that point to un-mapped abstract memory object.
func MachMakeMemoryEntry(targetTask VMMap, size *MemoryObjectSize, offset MemoryObjectOffset, permission VMProt, objectHandle *MemEntryNamePort, parentHandle MemEntryNamePort) error {
	_, _, errno := sys.Ccall9(mach_make_memory_entry_trampoline_addr, uintptr(targetTask), uintptr(unsafe.Pointer(&size)), uintptr(offset), uintptr(permission), uintptr(unsafe.Pointer(&objectHandle)), uintptr(parentHandle), 0, 0, 0)
	if errno != 0 {
		return sys.KernReturn(errno)
	}

	return nil
}
