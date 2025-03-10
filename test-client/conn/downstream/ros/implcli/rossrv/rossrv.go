package rossrv

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"

	parser2 "github.com/0x587/guardeye/test-client/conn/downstream/ros/implcli/rossrv/parser"
)

func Parse(s string) (parser2.IRootContext, error) {
	input := antlr.NewInputStream(s)
	lexer := parser2.NewRossrvLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser2.NewRossrvParser(stream)
	el := newErrListener()
	p.AddErrorListener(antlr.NewProxyErrorListener([]antlr.ErrorListener{
		el,
	}))
	res := p.Root()
	if el.GetErrs() == nil {
		return res, nil
	}
	return nil, errors.Join(el.GetErrs()...)
}
