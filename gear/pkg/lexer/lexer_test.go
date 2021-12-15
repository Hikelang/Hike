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

package lexer

import (
	"fmt"
	"os"
	"testing"

	"gear/internal/logging"
	"github.com/alecthomas/participle/v2/lexer"
)

type Token struct {
	Type    lexer.TokenType
	Literal string
}

var TestIndex uint = 0

func LexerTest(t *testing.T, input string, excepted []Token) {
	tokens := GetTokens("", input)
	if len(tokens) != len(excepted) {
		logging.PrintGeneralError(fmt.Sprint("Excepted:", excepted, " Got: ", tokens), "If you are developer and you are implementing some feature and this test failed, try to rewrite your code.\nBut if you are a user of hike programming language, you can send this error message into issues tab.")
		os.Exit(1)
	}
	for indx, token := range tokens {
		if token.Type != excepted[indx].Type {
			logging.PrintGeneralError(fmt.Sprint("Excepted:", excepted, " Got: ", tokens), "If you are developer and you are implementing some feature and this test failed, try to rewrite your code.\nBut if you are a user of hike programming language, you can send this error message into issues tab.")
			os.Exit(1)
		}
		if token.Value != excepted[indx].Literal {
			logging.PrintGeneralError(fmt.Sprint("Excepted:", excepted, " Got: ", tokens), "If you are developer and you are implementing some feature and this test failed, try to rewrite your code.\nBut if you are a user of hike programming language, you can send this error message into issues tab.")
			os.Exit(1)
		}
	}
	TestIndex++
	fmt.Println(fmt.Sprint("Test ", TestIndex, " passed: ", input))
}

func TestEOF(t *testing.T) {
	LexerTest(t, "", []Token{})
}

func TestWhitespace(t *testing.T) {
	LexerTest(t, " ", []Token{{Type: -3, Literal: " "}})
	LexerTest(t, "\n", []Token{{Type: -3, Literal: "\n"}})
	LexerTest(t, "\t", []Token{{Type: -3, Literal: "\t"}})
}

func TestDigits(t *testing.T) {
	LexerTest(t, "1 2 3 4 5 6 7 8 9 0", []Token{{Type: -8, Literal: "1"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "2"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "3"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "4"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "5"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "6"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "7"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "8"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "9"},
		{Type: -3, Literal: " "},
		{Type: -8, Literal: "0"}})
}

func TestInteger(t *testing.T) {
	LexerTest(t, "1234567890", []Token{{Type: -8, Literal: "1234567890"}})
}

func TestFloatPointNumber(t *testing.T) {
	LexerTest(t, "1234567890.1234567890", []Token{{Type: -8, Literal: "1234567890.1234567890"}})
	// LexerTest(t, ".1234567890", []Token{{Type: -8, Literal: ".1234567890"}})
}

func TestNumberSystemsLiterals(t *testing.T) {
	LexerTest(t, "0x0123456789abcdef", []Token{{Type: -9, Literal: "0x0123456789abcdef"}})
	LexerTest(t, "0X0123456789abcdef", []Token{{Type: -9, Literal: "0X0123456789abcdef"}})
	LexerTest(t, "0X0123456789ABCDEF", []Token{{Type: -9, Literal: "0X0123456789ABCDEF"}})
	LexerTest(t, "0b01", []Token{{Type: -10, Literal: "0b01"}})
	LexerTest(t, "0B01", []Token{{Type: -10, Literal: "0B01"}})
	LexerTest(t, "0o01234567", []Token{{Type: -11, Literal: "0o01234567"}})
	LexerTest(t, "0O01234567", []Token{{Type: -11, Literal: "0O01234567"}})
}

func TestStringLiteral(t *testing.T) {
	LexerTest(t, "\"hello\"", []Token{{Type: -13, Literal: "\"hello\""}})
	LexerTest(t, "\"hello world\"", []Token{{Type: -13, Literal: "\"hello world\""}})
	LexerTest(t, "\"hello\" \"world\"", []Token{{Type: -13, Literal: "\"hello\""}, {Type: -3, Literal: " "}, {Type: -13, Literal: "\"world\""}})
}

func TestBoolLiteral(t *testing.T) {
	LexerTest(t, "true", []Token{{Type: -5, Literal: "true"}})
	LexerTest(t, "false", []Token{{Type: -5, Literal: "false"}})
}

func TestIdentifier(t *testing.T) {
	LexerTest(t, "some_cool_identifier", []Token{{Type: -6, Literal: "some_cool_identifier"}})
}

func TestComment(t *testing.T) {
	LexerTest(t, "// comment", []Token{{Type: -2, Literal: "// comment"}})
}
