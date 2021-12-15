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

package docgen

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gear/internal/logging"
	"gear/pkg/ast/definitions"
	"gear/pkg/parser"
)

type DocGenerator struct {
	Dir       string
	DocBuffer string
	CurAST    *definitions.Program
}

func NewDocGenerator(dir string) DocGenerator {
	return DocGenerator{
		Dir: dir,
	}
}

func (g *DocGenerator) Gen(filename string, code string) {
	g.CurAST = parser.GetParserResult(filename, &code)
	g.DocBuffer = g.ModuleDocGen(filename, g.CurAST)
	if g.CurAST.TopStmts != nil {
		for _, topStmt := range g.CurAST.TopStmts {
			if topStmt.FuncDefStmt != nil {
				g.DocBuffer += g.FunctionDocGen(topStmt.FuncDefStmt)
				body := topStmt.FuncDefStmt.FuncBody
				for _, stmt := range body {
					fmt.Println(EvaluateTokens(stmt.Tokens))
				}
			}
		}
	}
}

func (g *DocGenerator) FileGen() {
	os.Mkdir("docs", 0777)
	for _, file := range g.ListDir() {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			logging.PrintGeneralError(err.Error(), "")
			os.Exit(1)
		}
		code := string(bytes)
		g.Gen(file, code)
		readmeFilename := strings.Replace(file, "/", "_", -1)
		readmeFilename = strings.TrimSuffix(readmeFilename, ".hk") + ".md"
		err = ioutil.WriteFile("docs/"+readmeFilename, []byte(g.DocBuffer), 0644)
		fmt.Println(readmeFilename)
		if err != nil {
			logging.PrintGeneralError(err.Error(), "")
			os.Exit(1)
		}
	}
}
