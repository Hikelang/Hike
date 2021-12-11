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
package definitions

import "github.com/alecthomas/participle/v2/lexer"

// Function declaration Abstract Syntax Tree
type FuncDefStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Name of a function, for example:
	// foo = ...
	FuncName string `@Id "=" `

	Generics []*string `("<" (@Id ("," @Id)*)? ">")?`

	// Function parameters can be represented as scalar definitions, example:
	// foo = (a int, b int)
	FuncParameters []*ScalarDef `"(" (@@ ("," @@)*)? ")"`

	// The type returned by the function (may lack). Example:
	// foo = (a int, b int) int
	FuncReturnType *Type `@@?`

	// Function body:
	// foo = (a int, b int) int {
	//   return a + b;
	// }
	FuncBody []*Stmt `"{" @@* "}"`
}

// Anonymous functions are immediately created and executed, because they don't have a name,
// so there is only one Abstract Syntax Tree for them here.
type AnonymousFuncCall struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Parameters []*ScalarDef `"(" (@@ ("," @@)*)? ")"`
	ReturnType *Type        `@@?`
	FuncBody   []*Stmt      `"{" @@* "}"`
	Arguments  []*Expr      `"(" (@@ ("," @@)*)? ")"`
}

// Run coroutine, just need to add "co" keyword
type CoroutineCall struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	FuncCall          *Expr              `"co" @@`
	AnonymousFuncCall *AnonymousFuncCall `| "co" @@`
}

// Export function to another packages. Similar to `CoroutineCall`, just one
// keyword is added
type ExportFunc struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	FuncDefStmt *FuncDefStmt `"export" @@`
}

type TypeConversion struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Type *Type `@@ "("`
	Expr *Expr `@@ ")"`
}

type Extern struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name       string       `"extern" @Id`
	Parameters []*ScalarDef `"(" (@@ ("," @@)*)? ")"`
	Type       *Type        `@@? ";"`
}
