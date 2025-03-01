package rosmsg

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"

	"github.com/0x587/guardeye/common/rosmsg/parser"
)

func ParseMsg(s string) (parser.IMsg_statContext, error) {
	input := antlr.NewInputStream(s)
	lexer := parser.NewRosmsgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRosmsgParser(stream)
	el := newErrListener()
	p.AddErrorListener(antlr.NewProxyErrorListener([]antlr.ErrorListener{
		el,
	}))
	res := p.Msg_stat()
	if el.GetErrs() == nil {
		return res, nil
	}
	return nil, errors.Join(el.GetErrs()...)
}

func ParseSrv(s string) (parser.ISrv_statContext, error) {
	input := antlr.NewInputStream(s)
	lexer := parser.NewRosmsgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRosmsgParser(stream)
	el := newErrListener()
	p.AddErrorListener(antlr.NewProxyErrorListener([]antlr.ErrorListener{
		el,
	}))
	res := p.Srv_stat()
	if el.GetErrs() == nil {
		return res, nil
	}
	return nil, errors.Join(el.GetErrs()...)
}
