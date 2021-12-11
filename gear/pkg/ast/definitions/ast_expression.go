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

// Array literal (comma-concatenated list of expressions => array)
type ArrayLiteral struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Exprs []*Expr `"{" (@@ ("," @@)*)? "}"`
}

// Why this structure is not named as ArrayIndex? The point
// is that the dictionary also has an index, or rather it has a key, and
// to get values by key from the dictionary, you need to use the same syntax
// as when getting an element in a list.

// It means, that this structure is needed to describe the syntax for getting
// an element or values in a list or dictionary, respectively
type Index struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Name *VarName `@@ "["`
	Expr *Expr    `@@ "]"`
}

type Expr struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Entry point of all binary operations
	Assignment *Assignment `@@`
}

// I wouldn't advise you, to write Expr "operator" Expr when adding a new operator
// here. This will raise a `github.com/alecthomas/participle` library error.
// It's all about the occurrence of left recursion. Therefore, all structures
// for binary operators are written in such a strange way here. But there is a clear
// division into the order of those in execution! For example, expressions with a
// multiplication sign are parsed first, and only then expressions with a plus sign (that is,
// multiplication is performed first of addition).

// Binary operations:
type Assignment struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Equality *Equality `@@`
	Op       string    `( @"="`
	Next     *Equality `  @@ )?`
}

type Equality struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Comparison *Comparison `@@`
	Op         string      `[ @( "!" "=" | "=" "=" )`
	Next       *Equality   `  @@ ]`
}

type Comparison struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	LogicalExpr *LogicalExpr `@@`
	Op          string       `[ @( ">" "=" | ">" | "<" "=" | "<" )`
	Next        *Comparison  `  @@ ]`
}

type LogicalExpr struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Addition *Addition    `@@`
	Op       string       `[ @( "&" "&" | "|" "|" )`
	Next     *LogicalExpr `  @@ ]`
}

type Addition struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Multiplication *Multiplication `@@`
	Op             string          `[ @( "-" | "+" )`
	Next           *Addition       `  @@ ]`
}

type Multiplication struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	NamespaceSeperator *NamespaceSeperator `@@`
	Op                 string              `[ @( "/" | "*" )`
	Next               *Multiplication     `  @@ ]`
}

type NamespaceSeperator struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Unary          *Unary              `@@`
	Op             string              `[ @( ".")`
	Next           *NamespaceSeperator `@@ ]`
	FuncParameters []*Expr             `("(" (@@ ("," @@)*)? ")")?`
}

type Unary struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	Op      string   `  ( @( "!" | "-" | "+" "+" | "-" "-" | "*" | "&" )`
	Unary   *Unary   `    @@ )`
	Primary *Primary `| @@`
}

// `Primary`. This is a main part of an expression. As we can see in Expr structure
// definition:
// type Expr struct {
// 	Assignment *Assignment `@@`
//  |_ Equality
//     |_ Comparison
//        |_ ....
//           |_ Primary <- this is an expression literal (value, that we operate with)

type NameOrFuncCall struct {
	Pos            lexer.Position
	EndPos         lexer.Position
	Tokens         []lexer.Token
	Id             *string `@Id`
	FuncParameters []*Expr `("(" (@@ ("," @@)*)? ")")?`
}

type Primary struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Simple literals
	String    string `@Str`
	Number    string `| @Num`
	HexNumber string `| @Hex`
	BinNumber string `| @Bin`
	OctNumber string `| @Oct`
	Bool      string `| @Bool`

	// More complex literals
	// More here `ast_dictionary.go`
	Dictionary *Dictionary `| @@`

	// More here `ast_class.go`
	InitializeClass *InitializeClass `| @@`

	// More here `ast_function.go`
	CoroutineCall     *CoroutineCall     `| @@`
	AnonymousFuncCall *AnonymousFuncCall `| @@`
	NameOrFuncCall    *NameOrFuncCall    `| @@`
	TypeConversion    *TypeConversion    `| @@`

	//                                                  /\
	// Other structures can be found higher in the code ||
	ArrayIndex   *Index        `| @@`
	ArrayLiteral *ArrayLiteral `| @@`
	Expr         *Expr         `| "(" @@ ")"`
}
