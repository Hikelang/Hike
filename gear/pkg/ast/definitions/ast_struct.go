package definitions

type StructDef struct {
	Name string `"struct" @Id "{"`
	ScalarDefs *ScalarDef `"(" (@@ ("," @@)*)? ")" "}"`
}

