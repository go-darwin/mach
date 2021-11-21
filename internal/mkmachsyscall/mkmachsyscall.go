// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

// Command mkmachsyscall generates mach syscalls assembly trampolines to call libSystem routines from Go.
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-clang/clang-v13/clang"
)

func init() {
	spew.Config = spew.ConfigState{
		Indent:                  "  ",
		SortKeys:                true, // maps should be spewed in a deterministic order
		DisableMethods:          false,
		DisablePointerMethods:   false,
		DisablePointerAddresses: false, // don't spew the addresses of pointers
		DisableCapacities:       false, // don't spew capacities of collections
		ContinueOnMethod:        true,  // recursion should continue once a custom error or Stringer interface is invoked
		SpewKeys:                false, // if unable to sort map keys then spew keys to strings and sort those
		MaxDepth:                4,     // maximum number of levels to descend into nested data structures.
	}
}

const ptrsize = 8 // Pointer size. All supported platforms are 64-bit.

var xnuVersion string

func main() {
	flag.StringVar(&xnuVersion, "version", "", "xun release version")
	flag.Parse()

	if err := mkMachSyscall(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const (
	syscallSWFilename = "syscall_sw.c"
	syscallSWPath     = "osfmk/kern/" + syscallSWFilename

	machTrapsFilename = "mach_traps.h"
	machTrapsPath     = "osfmk/mach/" + machTrapsFilename
)

func mkMachSyscall() error {
	syscallSW, err := xnuSource(xnuVersion, syscallSWPath)
	if err != nil {
		return fmt.Errorf("failed to get %s version %s: %w", syscallSWPath, xnuVersion, err)
	}
	_ = syscallSW
	machTraps, err := xnuSource(xnuVersion, machTrapsPath)
	if err != nil {
		return fmt.Errorf("failed to get %s version %s: %w", machTrapsPath, xnuVersion, err)
	}
	// remove PRIVATE define if statement
	machTraps = bytes.ReplaceAll(machTraps, []byte("#ifdef  PRIVATE"), nil)
	machTraps = bytes.ReplaceAll(machTraps, []byte("#endif  /* PRIVATE */"), nil)

	idx := clang.NewIndex(1, 0)
	defer idx.Dispose()

	// log.Printf("parse %s file\n", syscallSWFilename)
	// parseCfile(idx, syscallSWFilename, []clang.UnsavedFile{clang.NewUnsavedFile(syscallSWFilename, string(syscallSW))})
	fmt.Printf("parse %s file\n", machTrapsFilename)
	parseCfile(idx, machTrapsFilename, []clang.UnsavedFile{clang.NewUnsavedFile(machTrapsFilename, string(machTraps))})

	// in = string(in1) + string(in2)
	// if err := writeASMFile(in, fmt.Sprintf("zsyscall_darwin_%s.s", arch), "go1.13"); err != nil {
	// 	return fmt.Errorf("failed to writeASMFile: %w", err)
	// }

	return nil
}

const clangFlags = uint32(clang.TranslationUnit_DetailedPreprocessingRecord |
	clang.TranslationUnit_Incomplete |
	clang.TranslationUnit_PrecompiledPreamble |
	clang.TranslationUnit_CacheCompletionResults |
	clang.TranslationUnit_ForSerialization |
	clang.TranslationUnit_CXXChainedPCH |
	clang.TranslationUnit_SkipFunctionBodies |
	clang.TranslationUnit_IncludeBriefCommentsInCodeCompletion |
	clang.TranslationUnit_CreatePreambleOnFirstParse |
	clang.TranslationUnit_KeepGoing)

func parseCfile(idx clang.Index, filename string, unsavedFiles []clang.UnsavedFile) {
	tu := idx.ParseTranslationUnit(filename, nil, unsavedFiles, clangFlags)
	defer tu.Dispose()

	cursor := tu.TranslationUnitCursor()

	var visit clang.CursorVisitor
	visit = func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		if cursor.IsNull() {
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {
		case clang.Cursor_FunctionDecl:
			fmt.Printf("%s %s(", cursor.ResultType().Spelling(), cursor.Spelling())

			numArgs := cursor.NumArguments()
			for i := int32(0); i < numArgs; i++ {
				fmt.Printf("%s %s", cursor.Argument(uint32(i)).Type().Spelling(), cursor.Argument(uint32(i)).Spelling())
				if i+1 < numArgs {
					fmt.Printf(", ")
				}
			}
			fmt.Println(")")

			return clang.ChildVisit_Recurse

		case clang.Cursor_ParmDecl:
			return clang.ChildVisit_Recurse

		default:
			return clang.ChildVisit_Recurse
		}
	}

	cursor.Visit(visit)
}

func xnuSource(version, path string) ([]byte, error) {
	const xunURI = "https://raw.githubusercontent.com/apple-opensource/xnu/%s/%s"

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf(xunURI, version, path), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to cerate http request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not fetch %s file source: %w", filepath.Base(path), err)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read resp.Body: %w", err)
	}

	return buf, nil
}

func writeASMFile(in string, fileName string, buildTags string) error {
	trampolines := map[string]bool{}

	var out bytes.Buffer

	p(&out, "// go run machsyscall_darwin_amd64.go %s\n", strings.Join(os.Args[1:], " "))
	p(&out, "// Code generated by the command above; DO NOT EDIT.\n")
	p(&out, "\n")
	p(&out, "//go:build %s\n", buildTags)
	p(&out, "// +build %s\n", buildTags)
	p(&out, "\n")
	p(&out, "#include \"textflag.h\"\n")
	for _, line := range strings.Split(in, "\n") {
		const prefix = "var "
		const suffix = "_trampoline_addr uintptr"

		if !strings.HasPrefix(line, prefix) || !strings.HasSuffix(line, suffix) {
			continue
		}

		fn := strings.TrimSuffix(strings.TrimPrefix(line, prefix), suffix)
		if !trampolines[fn] {
			trampolines[fn] = true
			p(&out, "\nTEXT %s_trampoline<>(SB),NOSPLIT,$0-0\n", fn)
			p(&out, "\tJMP\t%s(SB)\n\n", fn)
			p(&out, "GLOBL\t·%s_trampoline_addr(SB), RODATA, $%d\n", fn, ptrsize)
			p(&out, "DATA\t·%[1]s_trampoline_addr(SB)/%[2]d, $%[1]s_trampoline<>(SB)\n", fn, ptrsize)
		}
	}
	if err := os.WriteFile(fileName, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("can't write %s: %w", fileName, err)
	}

	return nil
}

func p(w io.Writer, format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
