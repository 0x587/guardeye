package listener

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"

	"github.com/0x587/guardeye/api/internal/logic/es/parser"
)

func NewResListener() *ResListener {
	return &ResListener{
		qe: &QueryEntry{},
	}
}

type EnvInjector interface {
	GetMsg() string
}

type ResListener struct {
	*parser.BaseGuardQueryListener
	qe *QueryEntry
}

type QueryEntry struct {
	ses []*SourceEntry
	res []*ResultEntry
}

type SourceEntry struct {
	Node struct {
		Any bool
		Nid string
	}
	Providers []struct {
		Any   bool
		PType string
		PArgs []string
	}
}

func (l *ResListener) EnterSource(ctx *parser.SourceContext) {
	s := &SourceEntry{}
	node := ctx.Node()
	if node == nil {
		return
	}
	if node.STAR() != nil {
		s.Node.Any = true
	} else {
		s.Node.Nid = node.GetNodeId().GetText()
	}
	for _, p := range ctx.AllProvider() {
		pl := struct {
			Any   bool
			PType string
			PArgs []string
		}{}
		if p.STAR() != nil {
			pl.Any = true
		} else {
			pl.PType = p.GetProviderType().GetText()
			pl.PArgs = lo.Map(p.AllProviderArg(), func(a parser.IProviderArgContext, _ int) string {
				return a.GetText()
			})
		}
		s.Providers = append(s.Providers, pl)
	}
	l.qe.ses = append(l.qe.ses, s)
}

type ValueEntry struct {
	Vf func(EnvInjector) (any, error)
}

type ResultEntry struct {
	Value *ValueEntry
	Alias string
}

func (l *ResListener) EnterResultColumn(ctx *parser.ResultColumnContext) {
	re := &ResultEntry{}
	re.Value = &ValueEntry{Vf: l.makeValueEntryFunc(ctx.ValueExpr())}
	alias := ctx.ResultAlias()
	if alias != nil {
		re.Alias = alias.GetText()
	}
	l.qe.res = append(l.qe.res, re)
}

func (l *ResListener) makeValueEntryFunc(ctx parser.IValueExprContext) func(EnvInjector) (any, error) {
	switch {
	case ctx.NUMERIC_LITERAL() != nil:
		return func(_ EnvInjector) (any, error) {
			return strconv.ParseInt(ctx.NUMERIC_LITERAL().GetText(), 10, 64)
		}
	case ctx.STRING_LITERAL() != nil:
		return func(_ EnvInjector) (any, error) {
			s := ctx.STRING_LITERAL().GetText()
			s = s[1 : len(s)-1]
			return s, nil
		}
	case ctx.IDENTIFIER() != nil:
		return func(_ EnvInjector) (any, error) {
			return ctx.IDENTIFIER().GetText(), nil
		}
	case ctx.BuildinSource() != nil:
		return func(mi EnvInjector) (any, error) {
			switch ctx.BuildinSource().GetText() {
			case "$msg":
				return mi.GetMsg(), nil
			default:
				// TODO
				return ctx.BuildinSource().GetText(), nil
			}
		}
	case ctx.DOT() != nil:
		return func(mi EnvInjector) (any, error) {
			express := ctx.AllValueExpr()
			if len(express) != 2 {
				return nil, errors.Errorf("len(exprs) %d != 2", len(express))
			}
			leftValue, err := l.makeValueEntryFunc(express[0])(mi)
			if err != nil {
				return nil, err
			}
			leftValueAsMap, ok := leftValue.(map[string]any)
			if !ok {
				return nil, errors.Errorf("want got a map, but got %T", leftValue)
			}
			rightValue, err := l.makeValueEntryFunc(express[1])(mi)
			if err != nil {
				return nil, err
			}
			rightValueAsString, ok := rightValue.(string)
			if !ok {
				return nil, errors.Errorf("want got a string, but got %T", rightValue)
			}
			v, ok := leftValueAsMap[rightValueAsString]
			if !ok {
				return nil, errors.Errorf("key %s not found", rightValueAsString)
			}
			return v, nil
		}
	case ctx.OPEN_BRA() != nil:
		return func(mi EnvInjector) (any, error) {
			express := ctx.AllValueExpr()
			if len(express) != 2 {
				return nil, errors.Errorf("len(exprs) %d != 2", len(express))
			}
			leftValue, err := l.makeValueEntryFunc(express[0])(mi)
			if err != nil {
				return nil, err
			}
			leftValueAsArr, ok := leftValue.([]any)
			if !ok {
				return nil, errors.Errorf("want got a array, but got %T", leftValue)
			}
			rightValue, err := l.makeValueEntryFunc(express[1])(mi)
			if err != nil {
				return nil, err
			}
			var index int64
			rightValueAsInt, ok1 := rightValue.(int64)
			if ok1 {
				index = rightValueAsInt
			}
			rightValueAsDouble, ok2 := rightValue.(float64)
			if ok2 {
				index = int64(rightValueAsDouble)
			}
			if !(ok1 && ok2) {
				return nil, errors.Errorf("want got a int64, but got %T", rightValue)
			}
			if index >= int64(len(leftValueAsArr)) {
				return nil, errors.Errorf("index out of range")
			}
			if index < 0 {
				return nil, errors.Errorf("index must be positive")
			}
			return leftValueAsArr[index], nil
		}
	case ctx.BuildinFunction() != nil:
		args := ctx.AllValueExpr()
		var unmarshalFun func([]byte, any) error
		switch ctx.BuildinFunction().GetText() {
		case "json":
			unmarshalFun = json.Unmarshal
		case "yaml":
			unmarshalFun = yaml.Unmarshal
		}
		return func(mi EnvInjector) (any, error) {
			if len(args) != 1 {
				return nil, errors.Errorf("len(args) %d != 1", len(args))
			}
			var obj any
			value, err := l.makeValueEntryFunc(args[0])(mi)
			if err != nil {
				return nil, err
			}
			valueAsString, ok := value.(string)
			if !ok {
				return nil, errors.Errorf("want got a string, but got %T", value)
			}
			if err = unmarshalFun([]byte(valueAsString), &obj); err != nil {
				return nil, errors.Wrapf(err, "fail in func %s", ctx.BuildinFunction().GetText())
			}
			return obj, nil
		}
	case ctx.OPEN_PAR() != nil:
		return l.makeValueEntryFunc(ctx.ValueExpr(0))
	}
	return nil
}
