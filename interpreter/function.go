/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * function.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	r "reflect"
	// "strings"
)

func packValues(val0 r.Value, vals []r.Value) []r.Value {
	if len(vals) == 0 && val0 != None {
		vals = []r.Value{val0}
	}
	return vals
}

func unpackValues(vals []r.Value) (r.Value, []r.Value) {
	val0 := None
	if len(vals) > 0 {
		val0 = vals[0]
	}
	return val0, vals
}

func (env *Env) evalDeclNamedFunction(node *ast.FuncDecl) (r.Value, []r.Value) {
	name := node.Name.Name
	if name == temporaryFunctionName {
		// do *NOT* use env.evalBlock(), because it would create all bindings
		// in its block scope -> they are lost after env.evalBlock() returns
		return env.evalStatements(node.Body.List)
	}

	fun, t := env.evalDeclFunction(node, node.Type, node.Body)
	ret := env.defineVar(name, t, fun)
	return ret, nil
}

func (env *Env) evalDeclFunction(nodeForReceiver *ast.FuncDecl, funcType *ast.FuncType, body *ast.BlockStmt) (r.Value, r.Type) {
	var ret r.Value
	isMacro := false

	if nodeForReceiver != nil && nodeForReceiver.Recv != nil {
		recvList := nodeForReceiver.Recv.List
		if recvList != nil && len(recvList) == 0 {
			isMacro = true
		} else {
			// TODO implement receiver
			env.Errorf("unimplemented: method declarations (i.e. functions with receiver): %v", nodeForReceiver)
			return ret, nil
		}
	}
	t, argNames, resultNames := env.evalTypeFunction(funcType)
	tret := t
	funcName := nodeForReceiver.Name.Name
	if isMacro {
		funcName = "macro " + funcName
	} else {
		funcName = "func " + funcName
	}

	closure := func(args []r.Value) (results []r.Value) {
		return env.evalFuncCall(funcName, body, t, argNames, args, resultNames)
	}
	if isMacro {
		// env.Debugf("defined macro %v, type %v, args (%v), returns (%v)", nodeForReceiver.Name.Name, t, strings.Join(argNames, ", "), strings.Join(resultNames, ", "))
		ret = r.ValueOf(Macro{Closure: closure, ArgNum: len(argNames)})
		tret = TypeOf(ret) // do NOT change t, is needed by the closure above
	} else {
		ret = r.MakeFunc(t, closure)
	}
	return ret, tret
}

// eval an interpreted function
func (env *Env) evalFuncCall(envName string, body *ast.BlockStmt, t r.Type, argNames []string, args []r.Value, resultNames []string) (results []r.Value) {
	if t.Kind() != r.Func {
		return env.PackErrorf("call of non-function type %v", t)
	}
	env = NewEnv(env, envName)
	defer func() {
		if rec := recover(); rec != nil {
			if ret, ok := rec.(Return); ok {
				results = env.convertFuncCallResults(t, ret.Results, true)
			} else {
				panic(rec)
			}
		}
	}()

	for i, resultName := range resultNames {
		env.defineVar(resultName, t.Out(i), r.Zero(t.Out(i)))
	}
	for i, argName := range argNames {
		env.defineVar(argName, t.In(i), args[i])
	}
	// not env.evalBlock(): in Go, the function arguments and body are in the same scope
	rets := packValues(env.evalStatements(body.List))
	results = env.convertFuncCallResults(t, rets, false)
	return results
}

func (env *Env) convertFuncCallResults(t r.Type, rets []r.Value, warn bool) []r.Value {
	retsN := len(rets)
	expectedN := t.NumOut()
	if retsN < expectedN {
		if warn {
			env.Warnf("not enough return values: expected %d, found %d: %v", expectedN, retsN, rets)
		}
		tmp := make([]r.Value, expectedN)
		copy(tmp, rets)
		rets = tmp
	} else if retsN > expectedN {
		if warn {
			env.Warnf("too many return values: expected %d, found %d: %v", expectedN, retsN, rets)
		}
		rets = rets[:expectedN]
	}
	for i := range rets {
		rets[i] = rets[i].Convert(t.Out(i))
	}
	return rets
}