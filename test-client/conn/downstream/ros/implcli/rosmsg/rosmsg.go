package rosmsg

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"

	parser2 "github.com/0x587/guardeye/test-client/conn/downstream/ros/implcli/rosmsg/parser"
)

func ParseMsg(s string) (parser2.IMsgStatContext, error) {
	input := antlr.NewInputStream(s)
	lexer := parser2.NewRosmsgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser2.NewRosmsgParser(stream)
	el := newErrListener()
	p.AddErrorListener(antlr.NewProxyErrorListener([]antlr.ErrorListener{
		el,
	}))
	res := p.MsgStat()
	if el.GetErrs() == nil {
		return res, nil
	}
	return nil, errors.Join(el.GetErrs()...)
}

func ParseSrv(s string) (parser2.ISrvStatContext, error) {
	input := antlr.NewInputStream(s)
	lexer := parser2.NewRosmsgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser2.NewRosmsgParser(stream)
	el := newErrListener()
	p.AddErrorListener(antlr.NewProxyErrorListener([]antlr.ErrorListener{
		el,
	}))
	res := p.SrvStat()
	if el.GetErrs() == nil {
		return res, nil
	}
	return nil, errors.Join(el.GetErrs()...)
}
