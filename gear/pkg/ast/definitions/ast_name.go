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

// In Hike programming language, there is a clear difference between the function name and
// the variable enclosed in namespaces:

// Functions naming looks like this:
type Name struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Namespace is separated with a name via "::" syntax.
	Namespace string `@Id ":" ":"`
	Name      *Name  `@@`

	// At the end of recursive parsing it will all become an usual identifier. (name => id)
	Id *string `| @Id`
}

// But variables naming looks like this:
type VarName struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Namespace is separated via `.` syntax.
	Namespace string `@Id "."`
	Name      *Name  `@@`

	Id *string `| @Id`
}

// So it means, that:
// some_namespace::foo()   -   this naming is used for functions
// some_namespace.foo      -   & this is used for variables
