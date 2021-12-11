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

// Class declaration expressions (so far, no inheritance)
type ClassDefStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Name of the class, for example:
	// class Person {...}
	Name string `"class" @Id "{"`

	// Class fields, both private and public. All fields must be
	// declared before the methods of this class!
	ClassScalarDef []*ClassScalarDef `(@@ ";")*`

	// Class' methods definitions
	ClassFuncDefStmts []*ClassFuncDefStmt `@@* "}"`
}

// Definitions of class' attributes (public/private)
type ClassScalarDef struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// If the field is public, then the keyword "pub" is
	// placed before the declaration
	ClassPublicScalarDef *ScalarDef `"pub" @@`

	// If not, if the field is private, then nothing but a scalar
	// declaration of the class attribute is needed
	ClassPrivateScalarDef *ScalarDef `| @@`
}

// Class' methods definitions (public/private)
type ClassFuncDefStmt struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Separation into public and private methods, works the same
	// as with class attributes
	PrivateFuncDefStmt *FuncDefStmt `@@`
	PublicFuncDefStmt  *FuncDefStmt `| "pub" @@`
}

type InitializeClass struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// If you want to call class constructor, you just need to write
	// new keyword, and then class name with constructor arguments, example:
	// let person = new Person("adi", "salimgereev")
	FuncCall *Expr `"new" @@`
}
