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

package logging

import (
	"github.com/fatih/color"
)

// Error that is related to a file, but not specific line of code
// for example: lack of `main` module or `main` function.
type GlobalError struct {
	Filename string
	Message  string
	Note     string
}

func NewGlobalError(filename string, message string, note string) GlobalError {
	return GlobalError{
		Filename: filename,
		Message:  message,
		Note:     note,
	}
}

func (g *GlobalError) Show() {
	r := color.New(color.Bold, color.FgRed)
	r.Printf("# %s: %s\n", g.Filename, g.Message)
	y := color.New(color.Bold, color.FgYellow)
	y.Printf("# 💡: %s\n", g.Note)
}

// Note: please use this function if message and note length is low
func PrintGlobalError(filename string, message string, note string) {
	g := NewGlobalError(filename, message, note)
	g.Show()
}
