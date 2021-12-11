/*
 *   Copyright (c) 2021 Adi Salimgereev
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

	"gear/pkg/ast/repr"
)

type MeasureCmd struct {
	Input string
}

type ParseCmd struct {
	Input string
}

func (m *MeasureCmd) Run(globals *Globals) error {
	now := time.Now()
	bytes, err := ioutil.ReadFile(m.Input)
	if err != nil {
		e := logging.NewGeneralError(err.Error(), "")
		e.Show()
		os.Exit(1)
	}
	code := string(bytes)
	s := logging.NewSuccess(
		fmt.Sprint("File read in ",
			time.Since(now)))
	s.Print()
	now = time.Now()
	parser.GetParserResult(m.Input, &code)
	s = logging.NewSuccess(
		fmt.Sprint("File parsed in ",
			time.Since(now)))
	s.Print()
	return nil
}

func (p *ParseCmd) Run(globals *Globals) error {
	bytes, err := ioutil.ReadFile(p.Input)
	if err != nil {
		e := logging.NewGeneralError(err.Error(), "")
		e.Show()
		os.Exit(1)
	}
	code := string(bytes)
	r := parser.GetParserResult(p.Input, &code)
	repr.Println(r)
	return nil
}
