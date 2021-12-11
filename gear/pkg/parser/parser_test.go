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

package parser

import (
	"fmt"
	"strings"
	"testing"

	"gear/pkg/ast/definitions"
)

func parserTest(t *testing.T, code string) {
	ast := &definitions.Program{}
	err := Parser.ParseString("", code, ast)
	if err != nil {
		fmt.Println("Failed to parse this code:", code)
		t.Fatal(err)
	}
	t.Log("Parsed: ", strings.Replace(code, "\n", "\\n", -1))
}

func TestModule(t *testing.T) {
	codes := []string{
		`module main;`,
		`module main; import "some_example_import";`,
		`module main; global a = 3;`,
		`module main;
main = () {}`,
	}

	for _, code := range codes {
		parserTest(t, code)
	}
}

func TestPrimaryLiteral(t *testing.T) {
	literals := []string{
		`1234567890`,
		`1234567890.1234567890`,
		`"string"`,
		`{1, 2, 3}`,
		`dict[str:str]{"key": "value"}`,
	}

	for _, code := range literals {
		parserTest(t, `module main; main = () {`+code+`;}`)
	}
}

func TestFunction(t *testing.T) {
	codes := []string{
		`foo = () {}`,
		`foo = () {print(1 + 2);}`,
		`foo = (a int, b int, c ...int) {}`,
		`foo = (a int, b int, c ...int) int {}`,
		`foo = <T, M>(a T, b M, c ...M) int {}`,
	}

	for _, code := range codes {
		parserTest(t, `module main;`+code)
	}
}

func TestFunctionCall(t *testing.T) {
	codes := []string{
		`foo();`,
		`print(foo());`,
		`foo(1, 2, 3, 4, 5);`,
		`print(foo(1, 2, 3, 4));`,
		`co print(1, 2, 3, 4);`,
		`co print(foo(1, 2, 3, 4));`,
	}

	for _, code := range codes {
		parserTest(t, `module main; main = () {`+code+`}`)
	}
}

func TestAnonymousFunction(t *testing.T) {
	codes := []string{
		`() {} ();`,
		`(a int, b int) {return a + b;} (1, 2);`,
		`co (a int, b int) {return a + b;} (1, 2);`,
	}

	for _, code := range codes {
		parserTest(t, `module main; main = () {`+code+`}`)
	}
}

func TestType(t *testing.T) {
	types := []string{
		`int`, `int8`, `int16`, `int32`, `int64`,
		`float`, `float32`, `float64`,
		`complex`, `complex64`, `complex128`,
		`string`, `byte`,

		`[int]`,
		`[[int]]`,

		`*int`,
		`[*int]`,
	}

	for _, code := range types {
		parserTest(t, `module main;main = () {let var `+code+` ;}`)
	}
}

func TestExtern(t *testing.T) {
	externals := []string{
		`extern printf();`,
		`extern printf(a str);`,
		`extern printf(a str, b ...str) str;`,
	}

	for _, code := range externals {
		parserTest(t, `module main;`+code)
	}
}

func TestClass(t *testing.T) {
	classDefs := []string{
		`class Person {}`,
		`class Person {
	pub name str;
}`,
		`class Person {
	pub name str;
	age int;
}`,
		`class Person {
	pub name str;
	age int;

	pub constructor = () {
		this.name = name;
		this.age = 23; 
	}

	pub sayMyName = () {
		print(this.name);
	}
}`,
		`class Person {
	pub name str;
	age int;

	pub constructor = () {
		this.name = name;
		this.age = 23;
	}

	changeAge = () {
		this.age = 32;
	}

	pub sayAge = () {
		this.changeAge();
		print(this.age);
	}
}`,
	}

	for _, code := range classDefs {
		parserTest(t, `module main;`+code)
	}
}

func TestClassConstructorCall(t *testing.T) {
	codes := []string{
		`new Person()`,
		`new Person(1, 2)`,
		`new Person(new PassportData("adi", "salimgereev"))`,
	}

	for _, code := range codes {
		parserTest(t, `module main; main=(){`+code+`;}`)
	}
}

func TestForLoop(t *testing.T) {
	codes := []string{
		`for i in range(1, 10, 2) {
			print(i);
		}`,
		`for i in someArray {
			print(i);
		}`,
		`for i in {1, 2, 3} {
			print(i);
		}`,
	}

	for _, code := range codes {
		parserTest(t, `module main; main=(){`+code+`}`)
	}
}

func TestWhileLoop(t *testing.T) {
	codes := []string{
		`let n = 10;
while n <= 20 {
	print(n);
	++n;	
}`,
		`while true {}`,
		`while n <= 20 {
	while n > 3 {
		++n;
	}			
}`,
	}

	for _, code := range codes {
		parserTest(t, `module main; main=(){`+code+`}`)
	}
}

func TestBinaryExpr(t *testing.T) {
	codes := []string{
		`1`,
		`1+1`,
		`1*1`,
		`1+1*1`,
		`(1+1)*1`,
		`(1+1)*(a=2)`,
		`2&&(1+1)*(a=2)`,
		`2&&(1+1)*(a=2||3)`,
		`2&&(1+1<=2)*(a=(2||3==3))`,
	}

	for _, code := range codes {
		parserTest(t, `module main; main=(){`+code+`;}`)
	}
}

func TestOpEqual(t *testing.T) {
	codes := []string{
		`a += 3`,
		`a -= 3`,
		`a /= 3`,
		`a *= 3`,
		`a &&= 3`,
		`a ||= 3`,
	}

	for _, code := range codes {
		parserTest(t, `module main; main=(){`+code+`;}`)
	}
}

func TestNamespace(t *testing.T) {
	codes := []string{
		`a.a`,
		`a.a()`,
		`a().a`,
		`a.a.a`,
		`a.a().a`,
		`a.a().a()`,
		`a().a().a()`,
		`a().a.a()`,
	}

	for _, code := range codes {
		parserTest(t, `module main; main=(){`+code+`;}`)
	}
}
