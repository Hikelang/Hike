/*
 *   Copyright (c) 2021 Hike authors
 *   All rights reserved.

 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *   of this software and associated documentation files (the "Software"), to deal
 *   in the Software without restriction, including without limitation the rights
 *   to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *   copies of the Software, and to permit persons to whom the Software is
 *   furnished to do so, subject to the following conditions:

 *   The above copyright notice and this permission notice shall be included in all
 *   copies or substantial portions of the Software.

 *   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *   IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *   FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *   AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *   LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *   OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *   SOFTWARE.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gear/internal/logging"
	"gear/pkg/parser"
)

type BuildCmd struct {
	Input string
}

func (b *BuildCmd) Run(globals *Globals) error {
	now := time.Now()
	bytes, err := ioutil.ReadFile(b.Input)
	if err != nil {
		e := logging.NewGeneralError(err.Error(), "")
		e.Show()
		os.Exit(1)
	}
	code := string(bytes)
	s := logging.NewSuccess(
		fmt.Sprintf("File read in %d microseconds",
			time.Since(now).Microseconds()))
	s.Print()
	now = time.Now()
	r := parser.GetParserResult(b.Input, &code)
	s = logging.NewSuccess(
		fmt.Sprintf("File parsed in %d microseconds",
			time.Since(now).Microseconds()))
	s.Print()
	if r.Module != "main" {
		e := logging.NewGlobalError(
			b.Input,
			fmt.Sprintf("found `%s` module in compilation target.", r.Module),
			"you must change module name into `main`.",
		)
		e.Show()
	} else {
		s := logging.NewSuccess("Found `main` package")
		s.Print()
	}

	mainFunctionExists := false
	for _, stmt := range r.TopStmts {
		if stmt.FuncDefStmt != nil {
			if stmt.FuncDefStmt.FuncName == "main" {
				if stmt.FuncDefStmt.FuncReturnType.Type != "int" {
					e := logging.NewCodeError(b.Input, &code, stmt.FuncDefStmt.FuncReturnType.Pos, "`main` func. must return `int` type.", "try to change `main () ...` into `main () int`")
					e.Show()
				}
				mainFunctionExists = true
			}
		}
	}

	if !mainFunctionExists {
		e := logging.NewGlobalError(
			b.Input,
			"`main () int` not found in compilation target.",
			"you must add `main () int` func. into this file.",
		)
		e.Show()
	} else {
		s := logging.NewSuccess("Found `main` function")
		s.Print()
	}

	return nil
}
