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

type ArrayType struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// [int]
	// [[int]]
	// [([int])]
	Type *Type `"[" @@ "]"`
}

// type TupleType struct {
// 	// (int, float, int64, str, byte)
// 	// (int, (int, float))
// 	Types []*Type `"(" (@@ ("," @@)*)? ")"`
// }

type Pointer struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// *int, *[(int, float)], *[*int]
	Type *Type `"*" @@`
}

type Type struct {
	Pos    lexer.Position
	EndPos lexer.Position
	Tokens []lexer.Token

	// Type of every instance in hike can be:

	// Primary type (type token), for example:
	// str, bool, int64, float32, byte, ...
	Type string `@Type`

	// A name (declarated class), for example: Person
	Name *Name `| @@`

	// Array type, for example:
	// [str], [byte], [[bool]], ...
	ArrayType *ArrayType `| @@`

	// Removed!
	// Tuple type, example:
	// (str, bool, int, float), (str, str), ...
	// TupleType *TupleType `| @@`

	// Pointer, example:
	// *int, *[(int, float)], *[*int]
	Pointer *Pointer `| @@`

	// Combined:
	// ([int], float, [(str, byte)])
}
