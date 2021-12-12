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

package logging

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/fatih/color"
)

// Error related to a specific part of code inside a specific file.
type CodeError struct {
	Filename string
	Code     *string
	Pos      lexer.Position
	Message  string
	Note     string
}

func NewCodeError(filename string, code *string, pos lexer.Position, message string, note string) CodeError {
	return CodeError{
		Filename: filename,
		Code:     code,
		Pos:      pos,
		Message:  message,
		Note:     note,
	}
}

func (c *CodeError) Show() {
	r := color.New(color.Bold, color.FgRed)
	r.Printf("# %s: %d:%d: %s\n", c.Filename, c.Pos.Line, c.Pos.Column, c.Message)

	if c.Note != "" {
		y := color.New(color.Bold, color.FgYellow)
		y.Printf("#> note: %s\n", c.Note)
	}
}
