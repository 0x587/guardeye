package tokv

import (
	"reflect"
	"testing"
)

func TestYamlToKv(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		wantRes KV
	}{
		{
			"Test1",
			args{v: `
header:
  stamp:
    sec: 1731584501
    nanosec: 58992304
  frame_id: odom
`},
			map[string]string{
				"header.stamp.sec":     "1731584501",
				"header.stamp.nanosec": "58992304",
				"header.frame_id":      "odom",
			},
		},
		{
			"Test2",
			args{v: `
header:
  stamp:
    sec: 1731584501
    nanosec: 58992304
  frame_id: odom
covariance:
- 0.03
- 0.0
`},
			map[string]string{
				"covariance.0":         "0.03",
				"covariance.1":         "0",
				"header.stamp.sec":     "1731584501",
				"header.stamp.nanosec": "58992304",
				"header.frame_id":      "odom",
			},
		},
		{
			"Test3",
			args{v: `
headers:
- sec: 1731584501
  nanosec: 58992304
- sec: 1731584501
  nanosec: 58992304`},
			map[string]string{
				"headers.0.nanosec": "58992304",
				"headers.0.sec":     "1731584501",
				"headers.1.nanosec": "58992304",
				"headers.1.sec":     "1731584501",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := YamlToKv(tt.args.v); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("YamlToKv() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
