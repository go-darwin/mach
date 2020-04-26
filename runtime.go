// Copyright 2020 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,amd64

package mach

import (
	"unsafe"
)

//go:linkname funcPC runtime.funcPC
func funcPC(f interface{}) uintptr

//go:nosplit
//go:cgo_unsafe_args

// FuncPC returns the entry PC of the function f.
// It assumes that f is a func value. Otherwise the behavior is undefined.
// CAREFUL: In programs with plugins, funcPC can return different values
// for the same function (because there are actually multiple copies of
// the same function in the address space). To be safe, don't use the
// results of this function in any == expression. It is only safe to
// use the result as an address at which to start executing code.
func FuncPC(f interface{}) uintptr {
	return funcPC(f)
}

//go:nosplit
//go:linkname libcCall runtime.libcCall
func libcCall(fn, arg unsafe.Pointer) int32

//go:nosplit
//go:cgo_unsafe_args

// LibcCall calls a fn with arg as its argument. Return what fn returns.
// fn is the raw pc value of the entry point of the desired function.
// Switches to the system stack, if not already there.
// Preserves the calling point as the location where a profiler traceback will begin.
func LibcCall(fn, arg unsafe.Pointer) int32 {
	return libcCall(fn, arg)
}

//go:nosplit
//go:linkname syscall runtime.syscall
func syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr)

//go:nosplit
//go:cgo_unsafe_args

// Syscall calls a function in libc on behalf of the syscall package.
//
// syscall takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
func Syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	return syscall(fn, a1, a2, a3)
}

//go:nosplit
//go:linkname syscall6 runtime.syscall6
func syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr)

//go:nosplit
//go:cgo_unsafe_args

// Syscall6 calls a function in libc on behalf of the syscall package.
//
// syscall6 takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   a4    uintptr
//   a5    uintptr
//   a6    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall6 must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall6 expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
func Syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	return syscall6(fn, a1, a2, a3, a4, a5, a6)
}

//go:nosplit
//go:linkname syscall6X runtime.syscall6X
func syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr)

//go:nosplit
//go:cgo_unsafe_args

// Syscall6X calls a function in libc on behalf of the syscall package.
//
// syscall6X takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   a4    uintptr
//   a5    uintptr
//   a6    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall6X must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall6X is like syscall6 but expects a 64-bit result
// and tests for 64-bit -1 to decide there was an error.
func Syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	return syscall6X(fn, a1, a2, a3, a4, a5, a6)
}

//go:linkname syscallPtr runtime.syscallPtr
func syscallPtr(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr)

//go:nosplit

// SyscallPtr is like syscallX except that the libc function reports an
// error by returning NULL and setting errno.
func SyscallPtr(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	return syscallPtr(fn, a1, a2, a3)
}

//go:linkname rawSyscall runtime.rawSyscall
func rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr)

//go:nosplit
//go:cgo_unsafe_args

// RawSyscall calls a function in libc on behalf of the syscall package.
func RawSyscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	return rawSyscall(fn, a1, a2, a3)
}

//go:nosplit
//go:linkname rawSyscall6 runtime.rawSyscall6
func rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr)

//go:nosplit
//go:cgo_unsafe_args

// RawSyscall6 calls a function in libc on behalf of the syscall package.
func RawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	return rawSyscall6(fn, a1, a2, a3, a4, a5, a6)
}
