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
package definitions

import "github.com/alecthomas/participle/v2/lexer"

type Stmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	IfElseStmt *IfElseStmt `@@`
	// Not implemented normally, ...yet
	// ForStmt         *ForStmt         `| @@`
	WhileStmt  *WhileStmt  `| @@`
	ForEach    *ForEach    `| @@`
	ReturnStmt *ReturnStmt `| @@ ";"`
	VarDefStmt *VarDefStmt `| @@ ";"`
	Expr       *Expr       `| @@ ";"`
	OrEqual    *OrEqual    `| @@ ";"`
	AndEqual   *AndEqual   `| @@ ";"`
	PlusEqual  *PlusEqual  `| @@ ";"`
	MinusEqual *MinusEqual `| @@ ";"`
	StarEqual  *StarEqual  `| @@ ";"`
	SlashEqual *SlashEqual `| @@ ";"`
}

// TODO: Implement ForStmt without problems
type ForStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	VarDefStmt *VarDefStmt `"for" @@ ";"`
	Expr       *Expr       `@@ ";"`
	Counter    *Expr       `@@ "{"`
	Stmts      []*Stmt     `@@* "}"`
}

type ForEach struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name  string  `"for" @Id "in"`
	Expr  *Expr   `@@ "{"`
	Stmts []*Stmt `@@* "}"`
}

// If-Else Statement expression is required for sequential conditional checks
// and as part of three compound expressions:
type IfElseStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// if *Expr {...}
	IfStmt *IfStmt `@@`

	// else if *Expr {...}
	// else if *Expr {...}
	// ...
	// (not required)
	ElseIfStmts []*ElseIfStmt `(@@*)?`

	// else {...}
	// (not required)
	ElseStmt *ElseStmt `@@?`
}

type ElseIfStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Expr  *Expr   `"else" "if" @@`
	Stmts []*Stmt `"{" @@* "}"`
}

type ElseStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Expr  *Expr   `"else" "{"`
	Stmts []*Stmt `@@* "}"`
}

type IfStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Expr  *Expr   `"if" @@ "{"`
	Stmts []*Stmt `@@* "}"`
}

// Similarly:
type WhileStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Expr  *Expr   `"while" @@ "{"`
	Stmts []*Stmt `@@* "}"`
}

type ReturnStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Expr *Expr `"return" @@`
}

// Please not change an order of this structures, this can cause a lot of bugs
type OrEqual struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name *VarName `@@ "|" "|" "="`
	Expr *Expr    `@@`
}

type AndEqual struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name *VarName `@@ "&" "&" "="`
	Expr *Expr    `@@`
}

type PlusEqual struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name *VarName `@@ "+" "="`
	Expr *Expr    `@@`
}

type MinusEqual struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name *VarName `@@ "-" "="`
	Expr *Expr    `@@`
}
type StarEqual struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name *VarName `@@ "*" "="`
	Expr *Expr    `@@`
}

type SlashEqual struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name *VarName `@@ "/" "="`
	Expr *Expr    `@@`
}
