package es

import (
	"reflect"
	"testing"

	"github.com/antlr4-go/antlr/v4"

	"github.com/0x587/guardeye/api/internal/logic/es/listener"
)

func Test_parseQuery(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name     string
		args     args
		wantErrs []string
	}{
		{
			name: "1",
			args: args{
				query: `
select 
	json($msg) as value1, 
	json($msg, 'key2') as value2,
	json($msg, 'key3').key[1] as value3,
	json($msg, 'key4').key[json($msg,'index')] as value4,
	json($msg, 'key5').(json(3)[json(6)])[json(4)] as value5
from 
	node * provider *, 
	node nid1 provider *,
	node nid2 provider mqtt, file('/log')
where 
	json($msg, 'key3').key[1] > 1 and
	json($msg, 'key3').key[1] == 1 and
	(
		json($msg, 'key3').key[1] <> 1 or
		json($msg, 'key3').key[1] >= 1
	);`,
			},
			wantErrs: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotErrs := parseQuery(tt.args.query); !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("parseQuery() = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_scheduleQuery(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		args   args
		wanRes []any
	}{
		{
			name: "1",
			args: args{
				query: `
select 
	'{"a":123}' as value1,
	json('{"a":123}') as value2,
	json('{"a":123}').a as value3,
	json('{"a":123}').('a') as value4,
	json('{"a":[1,2,3,4]}').a[2] as value5,
	json('{"a":[1,2,3,4]}').a[json('{"b":2}').b] as value6
from 
	node * provider *;
`,
			},
			wanRes: []any{
				`{"a":123}`,
				map[string]any{"a": 123},
				123,
				123,
				3,
				3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := listener.NewResListener()
			tree, errs := parseQuery(tt.args.query)
			if len(errs) != 0 {
				t.Errorf("parseQuery() = %v", errs)
			}
			antlr.NewParseTreeWalker().Walk(l, tree)
			scheduleQuery(tree)
			//for index, re := range l.qe.res {
			//	v, err := re.value.vf()
			//	if err != nil {
			//		t.Errorf("evalQuery() = %v", err)
			//	}
			//	if fmt.Sprintf("%v", tt.wanRes[index]) != fmt.Sprintf("%v", v) {
			//		t.Errorf("want %#v, got %#v", tt.wanRes[index], v)
			//	}
			//}
		})
	}
	//antlr.NewParseTreeWalker().Walk(&listener{}, tree)
}
