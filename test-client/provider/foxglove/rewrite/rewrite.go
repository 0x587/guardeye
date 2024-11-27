package rewrite

type FoxgloveDefine struct {
	Name        string       `json:"name"`
	Definitions []Definition `json:"definitions"`
}

type Definition struct {
	Type        string `json:"type"`
	IsArray     bool   `json:"isArray"`
	Name        string `json:"name"`
	IsComplex   bool   `json:"isComplex"`
	ArrayLength int    `json:"arrayLength"`
	IsConstant  bool   `json:"isConstant"`
	Value       any    `json:"value"`
	ValueText   string `json:"valueText"`
}

func parse(schema string) ([]FoxgloveDefine, error) {

}
