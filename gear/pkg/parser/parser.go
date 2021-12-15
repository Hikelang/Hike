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

package parser

import (
	"os"

	"gear/internal/logging"
	"gear/pkg/ast/definitions"
	hikeLexer "gear/pkg/lexer"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
)

var Parser = participle.MustBuild(&definitions.Program{},
	participle.Lexer(hikeLexer.Lex),
	participle.Elide("Doc", "WS"),
	participle.UseLookahead(2),
)

func PrintParserOutput(filename string, code *string) {
	definitions := &definitions.Program{}
	defer func() {
		repr.Println(definitions)
	}()
	err := Parser.ParseString("", *code, definitions)
	if err != nil {
		codeError := logging.NewCodeError(filename, code, lexer.Position{}, err.Error(), "")
		codeError.Show()
		os.Exit(1)
	}
}

func GetParserResult(filename string, code *string) *definitions.Program {
	definitions := &definitions.Program{}
	err := Parser.ParseString("", *code, definitions)
	if err != nil {
		codeError := logging.NewCodeError(filename, code, lexer.Position{}, err.Error(), "")
		codeError.Show()
		os.Exit(1)
	}
	return definitions
}
