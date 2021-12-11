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

type VarDefStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Name of a variable (not *Name, but identifier token literal)
	Name string `"let" @Id`

	// Fixed type of a variable (Not required)
	Type *Type `@@?`

	// Value (Not required)
	Expr *Expr `("=" @@)?`
}

type ConstVarDefStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Name of a variable (not *Name, but identifier token literal)
	Name string `"const" @Id`

	// Fixed type of a variable
	Type *Type `@@?`

	// Value
	Expr *Expr `"=" @@`
}

// Similarly with previous Nodes, but just with "global" keyword:
type GlobalVarDefStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name  string `"global" @Id`
	Type  *Type  `@@? "="`
	Value *Expr  `@@`
}

type ScalarDef struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	/* only one function's argument at the input */
	Id   string `@Id`
	Type *Type  `@@`

	/* bigger amount of arguments (from 2 to ...) */
	// Should not be used at class attribute declaration !!! `ast_class.go`
	AId   string `| @Id`
	AType *Type  `"." "." "." @@`
}
