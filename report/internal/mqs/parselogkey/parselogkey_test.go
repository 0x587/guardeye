package parselogkey

import (
	"slices"
	"testing"
)

func Test_parseYaml(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test1",
			args{msg: `
header:
  stamp:
    sec: 1731584501
    nanosec: 58992304
  frame_id: odom
`},
			[]string{"header.stamp.sec", "header.stamp.nanosec", "header.frame_id"},
		},
		{
			"Test2",
			args{msg: `
header:
  stamp:
    sec: 1731584501
    nanosec: 58992304
  frame_id: odom
covariance:
- 0.03
- 0.0
`},
			[]string{"covariance[0]", "covariance[1]", "header.stamp.sec", "header.stamp.nanosec", "header.frame_id"},
		},
		{
			"Test3",
			args{msg: `
headers:
- sec: 1731584501
  nanosec: 58992304
- sec: 1731584501
  nanosec: 58992304
`},
			[]string{"headers[0].nanosec", "headers[0].sec", "headers[1].nanosec", "headers[1].sec"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseYaml(tt.args.msg)
			slices.Sort(got)
			slices.Sort(tt.want)
			if !slices.Equal(got, tt.want) {
				t.Errorf("parseYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}
