package docgen

import (
	"testing"
)

func TestFuncDefDocGen(t *testing.T) {
	g := NewDocGenerator(".")
	g.Gen("", `module main;

// @brief some function
foo = () {}

// @brief some another function
foo2 = () {
	print("hello");
	print(foo());
}`)
}
