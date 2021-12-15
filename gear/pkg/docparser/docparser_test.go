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

package docparser

import (
	"fmt"
	"strings"
	"testing"
)

func parserTest(t *testing.T, doc string) {
	ast := &Comment{}
	err := Parser.ParseString("", doc, ast)
	if err != nil {
		fmt.Println("Failed to parse this code:", doc)
		t.Fatal(err)
	}
	t.Log("Parsed: ", strings.Replace(doc, "\n", "\\n", -1))
}

func TestSimpleComment(t *testing.T) {
	parserTest(t, `hello`)
	parserTest(t, `hello world`)
	parserTest(t, `hello world
@brief some brief`)
	parserTest(t, `hello world
@brief some brief
@description some description`)
	parserTest(t, `hello world
@brief some brief
@description some description
@param a some cool argument
@param b some another cool argument`)
	parserTest(t, `hello world
@brief some brief
@description some description
@param a some cool argument
@param b some another cool argument
@return some cool description of return content`)
}
