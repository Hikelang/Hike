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

package docgen

import (
	"fmt"
	"os"
	"strings"

	"gear/pkg/ast/definitions"
	"gear/pkg/docparser"
)

func WordsToString(words []string) string {
	str := ""
	for _, word := range words {
		str += word + " "
	}
	return str
}

func (g *DocGenerator) FunctionDocGen(ast *definitions.FuncDefStmt) string {
	code := GetEverythingWithoutComments(ast.Tokens)
	doc := GetComment(ast.Tokens)
	commentAST := &docparser.Comment{}
	err := docparser.Parser.ParseString("", doc, commentAST)
	if err != nil {
		fmt.Println("Failed to parse this comment:", doc)
		os.Exit(1)
	}
	generatedDoc := ""
	generatedDoc += WordsToString(commentAST.FuncComment.Comment)
	if commentAST.FuncComment.Brief != nil {
		generatedDoc += "\n`@brief` " +
			WordsToString(commentAST.FuncComment.Brief)
	}
	return "### Function `" + ast.FuncName + "`:\n" +
		generatedDoc + "\n```\n" +
		strings.TrimFunc(code, func(c rune) bool {
			return c == '\r' || c == '\n'
		}) + "\n```"
}
