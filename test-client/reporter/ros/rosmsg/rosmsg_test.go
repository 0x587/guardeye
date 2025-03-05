package rosmsg

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"
)

//go:embed ros_std_msg.json
var rosStdMsgStr string

var example = `
int32[] unbounded_integer_array
int32[5] five_integers_array
int32[<=5] up_to_five_integers_array

string string_of_unbounded_size
string<=10 up_to_ten_characters_string

string[<=5] up_to_five_unbounded_strings
string<=10[] unbounded_array_of_strings_up_to_ten_characters_each
string<=10[<=5] up_to_five_strings_up_to_ten_characters_each
custom_type custom_type_field
`

func TestParseMsg(t *testing.T) {
	type args struct {
		s string
	}
	var tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Example",
			args: args{
				s: example,
			},
			wantErr: false,
		},
	}
	rosStdMsg := make(map[string]string)
	_ = json.Unmarshal([]byte(rosStdMsgStr), &rosStdMsg)
	for name, define := range rosStdMsg {
		tests = append(tests, struct {
			name    string
			args    args
			wantErr bool
		}{name: name, args: args{s: define}, wantErr: false})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "msg") {
				_, err := ParseMsg(tt.args.s)
				if (err != nil) != tt.wantErr {
					t.Errorf("\n------DEFINE------\n%s------------------\nParse() error = %v, wantErr %v", tt.args.s, err, tt.wantErr)
					return
				}
			} else if strings.Contains(tt.name, "srv") {
				_, err := ParseSrv(tt.args.s)
				if (err != nil) != tt.wantErr {
					t.Errorf("\n------DEFINE------\n%s------------------\nParse() error = %v, wantErr %v", tt.args.s, err, tt.wantErr)
					return
				}
			}
		})
	}
}
