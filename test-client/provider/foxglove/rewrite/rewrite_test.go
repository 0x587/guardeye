package rewrite

import (
	"testing"
)

func TestParseSchemas(t *testing.T) {
	tests := []struct {
		name     string
		schema   string
		expected []FoxgloveDefine
	}{
		{
			// Test case for resolving unqualified names
			name: "resolves unqualified names",
			schema: `
				Point[] points
				===========
				MSG: geometry_msgs/Point
				float64 x
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   true,
							IsComplex: true,
							Name:      "points",
							Type:      "geometry_msgs/Point",
						},
					},
				},
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "x",
							Type:      "float64",
						},
					},
					Name: "geometry_msgs/Point",
				},
			},
		},
		{
			// Test case for normalizing aliases
			name:   "normalizes aliases",
			schema: "char x\nbyte y",
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "x",
							Type:      "uint8",
						},
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "y",
							Type:      "int8",
						},
					},
				},
			},
		},
		{
			// Test case for handling array with complex types and unqualified names
			name: "handles array with complex types and unqualified names",
			schema: `
		Point[] points
		===========
		MSG: geometry_msgs/Point
		float64 x
	`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   true,
							IsComplex: true,
							Name:      "points",
							Type:      "geometry_msgs/Point",
						},
					},
				},
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "x",
							Type:      "float64",
						},
					},
					Name: "geometry_msgs/Point",
				},
			},
		},
		{
			// Test case for handling message types and their complex definitions
			name: "handles message types with complex definitions",
			schema: `
		geometry_msgs/Point[] points
		===========
		float64 x
		float64 y
	`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   true,
							IsComplex: true,
							Name:      "points",
							Type:      "geometry_msgs/Point",
						},
					},
				},
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "x",
							Type:      "float64",
						},
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "y",
							Type:      "float64",
						},
					},
					Name: "geometry_msgs/Point",
				},
			},
		},
		{
			// Test case for handling aliases like "char" and "byte"
			name:   "handles aliases for char and byte",
			schema: "char x\nbyte y",
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "x",
							Type:      "uint8",
						},
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "y",
							Type:      "int8",
						},
					},
				},
			},
		},
		{
			// Test case for handling unqualified message types
			name:   "handles unqualified message types",
			schema: "float64 x",
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "x",
							Type:      "float64",
						},
					},
				},
			},
		},
		{
			// Test case for resolving unqualified names
			name: "resolves unqualified names",
			schema: `
				Point[] points
				===========
				MSG: geometry_msgs/Point
				float64 x
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   true,
							IsComplex: true,
							Name:      "points",
							Type:      "geometry_msgs/Point",
						},
					},
				},
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "x",
							Type:      "float64",
						},
					},
					Name: "geometry_msgs/Point",
				},
			},
		},
		{
			// Test case for normalizing aliases
			name:   "normalizes aliases",
			schema: "char x\nbyte y",
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "x",
							Type:      "uint8",
						},
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "y",
							Type:      "int8",
						},
					},
				},
			},
		},
		{
			// Test case for ignoring comment lines
			name: "ignores comment lines",
			schema: `
				# your first name goes here
				string firstName

				# last name here
				### foo bar baz?
				string lastName
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "firstName",
							Type:      "string",
						},
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "lastName",
							Type:      "string",
						},
					},
				},
			},
		},
		{
			// Test case for parsing variable length arrays
			name: "parses variable length arrays",
			schema: `
				"string[] names"
				"int32[] names"
				"int64[] names"
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   true,
							IsComplex: false,
							Name:      "names",
							Type:      "string",
						},
					},
				},
				{
					Definitions: []Definition{
						{
							IsArray:   true,
							IsComplex: false,
							Name:      "names",
							Type:      "int32",
						},
					},
				},
				{
					Definitions: []Definition{
						{
							IsArray:   true,
							IsComplex: false,
							Name:      "names",
							Type:      "int64",
						},
					},
				},
			},
		},
		{
			// Test case for parsing fixed length array
			name:   "parses fixed length string array",
			schema: "string[3] names",
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							ArrayLength: 3,
							IsArray:     true,
							IsComplex:   false,
							Name:        "names",
							Type:        "string",
						},
					},
				},
			},
		},
		{
			// Test case for parsing nested complex types
			name: "parses nested complex types",
			schema: `
				string username
				Account account
				===========
				MSG: custom_type/Account
				string name
				uint16 id
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "username",
							Type:      "string",
						},
						{
							IsArray:   false,
							IsComplex: true,
							Name:      "account",
							Type:      "custom_type/Account",
						},
					},
				},
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "name",
							Type:      "string",
						},
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "id",
							Type:      "uint16",
						},
					},
					Name: "custom_type/Account",
				},
			},
		},
		{
			// Test case for parsing constants
			name: "returns constants",
			schema: `
				uint32 foo = 55
				int32 bar=-11 # Comment # another comment
				float32 baz= \t -32.25
				bool someBoolean = 0
				string fooStr = Foo    ${""}
				string EMPTY1 =  ${""}
				string EMPTY2 =
				string HASH = #
				string EXAMPLE="#comments" are ignored, and leading and trailing whitespace removed
				uint64 SMOOTH_MOVE_START = 0000000000000001 # e.g. kobuki_msgs/VersionInfo
				int64 LARGE_VALUE = -9223372036854775807
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							Name:       "foo",
							Type:       "uint32",
							IsConstant: true,
							Value:      55,
							ValueText:  "55",
						},
						{
							Name:       "bar",
							Type:       "int32",
							IsConstant: true,
							Value:      -11,
							ValueText:  "-11",
						},
						{
							Name:       "baz",
							Type:       "float32",
							IsConstant: true,
							Value:      -32.25,
							ValueText:  "-32.25",
						},
						{
							Name:       "someBoolean",
							Type:       "bool",
							IsConstant: true,
							Value:      false,
							ValueText:  "0",
						},
						{
							Name:       "fooStr",
							Type:       "string",
							IsConstant: true,
							Value:      "Foo",
							ValueText:  "Foo",
						},
						{
							Name:       "EMPTY1",
							Type:       "string",
							IsConstant: true,
							Value:      "",
							ValueText:  "",
						},
						{
							Name:       "EMPTY2",
							Type:       "string",
							IsConstant: true,
							Value:      "",
							ValueText:  "",
						},
						{
							Name:       "HASH",
							Type:       "string",
							IsConstant: true,
							Value:      "#",
							ValueText:  "#",
						},
						{
							Name:       "EXAMPLE",
							Type:       "string",
							IsConstant: true,
							Value:      "#comments\" are ignored, and leading and trailing whitespace removed",
							ValueText:  "#comments\" are ignored, and leading and trailing whitespace removed",
						},
						{
							Name:       "SMOOTH_MOVE_START",
							Type:       "uint64",
							IsConstant: true,
							Value:      1,
							ValueText:  "0000000000000001",
						},
						{
							Name:       "LARGE_VALUE",
							Type:       "int64",
							IsConstant: true,
							Value:      -9223372036854775807,
							ValueText:  "-9223372036854775807",
						},
					},
				},
			},
		},
		{
			// Test case for parsing python boolean values
			name: "works with python boolean values",
			schema: `
				bool Alive=True
				bool Dead=False
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							Name:       "Alive",
							Type:       "bool",
							IsConstant: true,
							Value:      true,
							ValueText:  "True",
						},
						{
							Name:       "Dead",
							Type:       "bool",
							IsConstant: true,
							Value:      false,
							ValueText:  "False",
						},
					},
				},
			},
		},
		{
			// Test case for handling type names for fields
			name: "handles type names for fields",
			schema: `
				time time
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							Name:      "time",
							Type:      "time",
							IsArray:   false,
							IsComplex: false,
						},
					},
				},
			},
		},
		{
			name: "allows numbers in package names",
			schema: `
				abc1/Foo2 value0
				===========
				MSG: abc1/Foo2
				int32 data
			`,
			expected: []FoxgloveDefine{
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: true,
							Name:      "value0",
							Type:      "abc1/Foo2",
						},
					},
				},
				{
					Definitions: []Definition{
						{
							IsArray:   false,
							IsComplex: false,
							Name:      "data",
							Type:      "int32",
						},
					},
					Name: "abc1/Foo2",
				},
			},
		},
	}

	// Loop through each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			types, err := parse(tt.schema)
			if err != nil {
				t.Fatalf("Error parsing schema: %v", err)
			}

			if !equalFoxgloveDefines(types, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, types)
			}
		})
	}
}
