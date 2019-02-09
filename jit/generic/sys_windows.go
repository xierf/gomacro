// +build windows

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * sys_windows.go
 *
 *  Created on May 25, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var PAGESIZE = windows.Getpagesize()

type memarea struct {
	addr, size uintptr
}

func (asm *Asm) mmap() memarea {
	if len(asm.code) == 0 {
		errorf("mmap(): code is empty")
	}
	asm.Epilogue()
	if VERBOSE {
		fmt.Printf("asm: %#v\n", asm.code)
	}
	size := uintptr((len(asm.code) + PAGESIZE - 1) &^ (PAGESIZE - 1))
	mem, err := windows.VirtualAlloc(0, size, windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if err != nil {
		errorf("sys/windows.VirtualAlloc failed: %v", err)
	}
	memcpy(mem, uintptr(unsafe.Pointer(&asm.code[0])), size)
	var old uint32
	err = windows.VirtualProtect(mem, size, windows.PAGE_EXECUTE_READ, &old)
	if err != nil {
		windows.VirtualFree(mem, 0, windows.MEM_RELEASE)
		errorf("sys/windows.VirtualProtect failed: %v", err)
	}
	return memarea{mem, size}
}

// memory copy. a bit slow, but avoids depending on CGO
func memcpy(dst uintptr, src uintptr, size uintptr) {
	for i := uintptr(0); i < size; i++ {
		*(*uint8)(unsafe.Pointer(dst + i)) = *(*uint8)(unsafe.Pointer(src + i))
	}
}

func munmap(mem memarea) {
	if mem.addr != 0 {
		windows.VirtualFree(area.addr, 0, windows.MEM_RELEASE)
	}
}