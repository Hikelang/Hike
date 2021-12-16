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

// Error not related to a syntax error, execution error,
// etc. It means that, the error is not related to the file,
// but, for example, to incorrectly passed arguments.
type GeneralError struct {
	Message string
	Note    string
}

func NewGeneralError(message string, note string) GeneralError {
	return GeneralError{
		Message: message,
		Note:    note,
	}
}

func (g *GeneralError) Show() {
	r := color.New(color.Bold, color.FgRed)
	r.Printf("# %s\n", g.Message)
	if g.Note != "" {
		y := color.New(color.Bold, color.FgYellow)
		y.Printf(" # 💡: %s\n", g.Note)
	}
}

// Note: please use this function if message length is low
func PrintGeneralError(message string, note string) {
	g := NewGeneralError(message, note)
	g.Show()
}
