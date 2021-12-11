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

package lexer

import (
	"os"

	"gear/internal/logging"
	"github.com/alecthomas/participle/v2/lexer"
)

// Regexp definitions for all tokens, that lexer can handle.
var Lex = lexer.MustSimple([]lexer.Rule{
	{Name: "Doc", Pattern: `//.*|/\*.*?\*/`, Action: nil},
	{Name: "WS", Pattern: `\s+`, Action: nil},

	{Name: "Type", Pattern: `\b(int|byte|str|float|int8|int16|int32|int64|float32|float64|complex64|complex128|bool)\b`, Action: nil},
	{Name: "Bool", Pattern: `\b(true|false)\b`, Action: nil},
	{Name: "Id", Pattern: `\b([a-zA-Z_][a-zA-Z0-9_]*)\b`, Action: nil},
	{Name: "Punct", Pattern: `[-,\|().*/+%{};&!=:<>]|\[|\]`, Action: nil},
	{Name: "Num", Pattern: `\b(([0-9]*[.])?[0-9]+)\b`, Action: nil},
	{Name: "Hex", Pattern: `\b(0[xX][0-9a-fA-F]+)\b`, Action: nil},
	{Name: "Bin", Pattern: `\b(0[bB][0-1]+)\b`, Action: nil},
	{Name: "Oct", Pattern: `\b(0[oO][0-7]+)\b`, Action: nil},
	{Name: "Keyword", Pattern: `\b(extern|new|dict|if|else|global|return|module|import|co|class)\b`, Action: nil},
	{Name: "Str", Pattern: `("[^"\\]*(\\.[^"\\]*)*"|'[^'\\]*(\\.[^'\\]*)*')`, Action: nil},
})

func GetTokens(filename string, code string) []lexer.Token {
	tokens := make([]lexer.Token, 0)
	statefulLexer, err := Lex.LexString(filename, code)
	if err != nil {
		e := logging.NewGeneralError(err.Error(), "")
		e.Show()
		os.Exit(1)
	}
	for {
		token, err := statefulLexer.Next()
		if token.EOF() {
			break
		}
		if err != nil {
			e := logging.NewGeneralError(err.Error(), "")
			e.Show()
			os.Exit(1)
		}
		tokens = append(tokens, token)
	}
	return tokens
}
