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
 * op2.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// binary operation
type Op2 uint8

const (
	AND = Op2(AND3)
	ADD = Op2(ADD3)
	ADC = Op2(ADC3) // add with carry
	MUL = Op2(MUL3)
	SHL = Op2(SHL3) // shift left
	SHR = Op2(SHR3) // shift right
	OR  = Op2(OR3)
	XOR = Op2(XOR3)
	SUB = Op2(SUB3)
	SBB = Op2(SBB3) // subtract with borrow

	DIV = Op2(DIV3)
	REM = Op2(REM3)

	MOV Op2 = 0x2B // implemented as OR3 xzr,src,dst

	NEG2 Op2 = 0xFD
	NOT2 Op2 = 0xFE
	CAST Op2 = 0xFF // TODO pick a value. sign extend, zero extend or narrow.

/*
	CMP  Op2 = ?? // compare, set flags
	XCHG Op2 = ?? // exchange
*/
)

var op2Name = map[Op2]string{
	ADD:  "ADD",
	AND:  "AND",
	ADC:  "ADC",
	MUL:  "MUL",
	SHL:  "SHL",
	SHR:  "SHR",
	OR:   "OR",
	XOR:  "XOR",
	SUB:  "SUB",
	SBB:  "SBB",
	DIV:  "DIV",
	REM:  "REM",
	MOV:  "MOV",
	CAST: "CAST",
	NEG2: "NEG2",
	NOT2: "NOT2",
	/*
		CMP:  "CMP",
		XCHG: "XCHG",
	*/
}

func (op Op2) String() string {
	s, ok := op2Name[op]
	if !ok {
		s = fmt.Sprintf("Op2(%d)", int(op))
	}
	return s
}

// ============================================================================
func (asm *Asm) Op2(op Op2, src Arg, dst Arg) *Asm {
	if op == CAST {
		if SizeOf(src) != SizeOf(dst) {
			return asm.Cast(src, dst)
		}
		op = MOV
	}
	if op == MOV {
		return asm.Mov(src, dst)
	}
	// dst OP= src
	//    translates to
	// dst = dst OP src
	//    note the argument order
	return asm.Op3(Op3(op), dst, src, dst)
}