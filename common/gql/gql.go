package gql

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/0x587/guardeye/common/gql/listener"
	"github.com/0x587/guardeye/common/gql/parser"
)

func ParseQuery(query string) (parser.IParseContext, []string) {
	input := antlr.NewInputStream(query)
	lexer := parser.NewGuardQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGuardQueryParser(stream)
	el := listener.NewErrListener()
	p.AddErrorListener(antlr.NewProxyErrorListener([]antlr.ErrorListener{
		el,
		antlr.NewDiagnosticErrorListener(false),
	}))
	return p.Parse(), el.GetErrs()
}

func ScheduleQuery(tree parser.IParseContext) *listener.Schedule {
	s := listener.NewScheduler()
	antlr.NewParseTreeWalker().Walk(s.Listener, tree)
	return s.GetSchedule()
}

func NewInjector(v string) *Injector {
	return &Injector{msg: v}
}

type Injector struct {
	msg string
}

func (i *Injector) GetMsg() string {
	return i.msg
}
