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

// Entry point of Abstract Syntax Tree (Root)
type Program struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Current module name
	Module string `"module" @Id ";"`

	// All module imports
	ImportStmts []*ImportStmt `@@*`

	Externs []*Extern `@@*`

	// Statements like function declaration, global variable declaration,
	// but not while loops, function calls
	// (they can't be outside of the function)
	TopStmts []*TopStmt `@@*`
}

// All module imports
type ImportStmt struct {
	// Statement of importing some module, for example:
	// import "./hello_world";
	Module string `"import" @Str ";"`
}

// Statements like function declaration, global variable declaration,
// but not while loops, function calls (they can't be outside of the
// function), and they all must be an the Root (ast.Program).
type TopStmt struct {
	// Check out `ast_variable.go`
	GlobalVarDefStmt *GlobalVarDefStmt `@@ ";"`

	// Check out `ast_function.go`
	ExportFunc  *ExportFunc  `| @@`
	FuncDefStmt *FuncDefStmt `| @@`

	// Check out `ast_class.go`
	ClassDefStmt *ClassDefStmt `| @@`
}
