package listener

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"

	"github.com/0x587/guardeye/api/internal/logic/es/parser"
	"github.com/0x587/guardeye/common/utils"
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

type (
	QueryEntry struct {
		ses []*SourceEntry
		res []*ResultEntry
		te  *TimeEntry
	}
	TimeEntry struct {
		From utils.ErrOr[time.Time]
		To   utils.ErrOr[time.Time]
	}
	SourceEntry struct {
		Node struct {
			Any bool
			Nid string
		}
		Providers []struct {
			Any   bool
			PType string
			PArgs []string
		}
		NeedKeys []string
	}
	SourceDependEntry struct {
		NeedKeys []string
	}
	ValueEntry struct {
		Vf           func(EnvInjector) (any, error)
		SourceDepend SourceDependEntry
		IsLiteral    bool
	}
	ResultEntry struct {
		Value *ValueEntry
		Alias string
	}
)

func (l *ResListener) EnterTimeStmt(ctx *parser.TimeStmtContext) {
	aCtx := ctx.AbsTimeStmt()
	if aCtx != nil {
		t1Str := aCtx.STRING_LITERAL(0).GetText()
		t1Str = strings.Replace(t1Str, "'", "", -1)
		t1, err := time.Parse("2006-01-02 15:04:05-07", t1Str)
		l.qe.te = &TimeEntry{
			From: utils.NewErrOr(t1, err),
		}
		if aCtx.TO_() != nil {
			t2Str := aCtx.STRING_LITERAL(1).GetText()
			t2Str = strings.Replace(t2Str, "'", "", -1)
			t2, err := time.Parse("2006-01-02 15:04:05-07", t2Str)
			l.qe.te.To = utils.NewErrOr(t2, errors.Wrap(err, "parse time fail"))
		} else {
			l.qe.te.To = utils.NewErrOr(time.Now(), nil)
		}
		return
	}
	rCtx := ctx.RelatTimeStmt()
	n := time.Now().In(lo.Must(time.LoadLocation("UTC")))
	t1, err := parseRelatTime(n, rCtx.NUMERIC_LITERAL(0), rCtx.TimeUnit(0))
	l.qe.te = &TimeEntry{
		From: utils.NewErrOr(t1, err),
	}
	if rCtx.TO_() != nil {
		t2, err := parseRelatTime(n, rCtx.NUMERIC_LITERAL(0), rCtx.TimeUnit(0))
		l.qe.te.To = utils.NewErrOr(t2, err)
	} else {
		l.qe.te.To = utils.NewErrOr(time.Now(), nil)
	}
}

func parseRelatTime(now time.Time, n antlr.TerminalNode, u parser.ITimeUnitContext) (time.Time, error) {
	parseInt, err := strconv.ParseInt(n.GetText(), 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	switch u.GetText() {
	default:
		return time.Time{}, errors.Errorf("unknown time unit %s", u.GetText())
	case "m":
		return now.Add(time.Duration(-parseInt) * time.Minute), nil
	case "h":
		return now.Add(time.Duration(-parseInt) * time.Hour), nil
	case "d":
		return now.Add(time.Duration(-parseInt) * 24 * time.Hour), nil
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

func (l *ResListener) EnterResultColumn(ctx *parser.ResultColumnContext) {
	re := &ResultEntry{}
	re.Value = l.makeValueEntry(ctx.ValueExpr())
	alias := ctx.ResultAlias()
	if alias != nil {
		re.Alias = alias.GetText()
	}
	l.qe.res = append(l.qe.res, re)
}

func (l *ResListener) makeValueEntry(ctx parser.IValueExprContext) *ValueEntry {
	switch {
	case ctx.NUMERIC_LITERAL() != nil:
		return &ValueEntry{
			Vf: func(_ EnvInjector) (any, error) {
				return strconv.ParseInt(ctx.NUMERIC_LITERAL().GetText(), 10, 64)
			},
			IsLiteral: true,
		}
	case ctx.STRING_LITERAL() != nil:
		return &ValueEntry{
			Vf: func(_ EnvInjector) (any, error) {
				s := ctx.STRING_LITERAL().GetText()
				s = s[1 : len(s)-1]
				return s, nil
			},
			IsLiteral: true,
		}
	case ctx.IDENTIFIER() != nil:
		return &ValueEntry{
			Vf: func(_ EnvInjector) (any, error) {
				return ctx.IDENTIFIER().GetText(), nil
			},
			IsLiteral: true,
		}
	case ctx.BuildinSource() != nil:
		return &ValueEntry{
			Vf: func(mi EnvInjector) (any, error) {
				switch ctx.BuildinSource().GetText() {
				case "$msg":
					return mi.GetMsg(), nil
				default:
					// TODO
					panic("implement me")
				}
			},
		}
	case ctx.DOT() != nil:
		var errInParse error
		express := ctx.AllValueExpr()
		if len(express) != 2 {
			errInParse = errors.Errorf("len(exprs) %d != 2", len(express))
		}
		leftValueEntry := l.makeValueEntry(express[0])
		rightValueEntry := l.makeValueEntry(express[1])
		var rightValueAsString string
		var ok bool
		// 字面量可直接求值
		if rightValueEntry.IsLiteral {
			rightValue, err := l.makeValueEntry(express[1]).Vf(nil)
			if err != nil {
				errInParse = err
			}
			rightValueAsString, ok = rightValue.(string)
			if !ok {
				errInParse = errors.Errorf("want got a string, but got %T", rightValue)
			}
		}
		res := &ValueEntry{
			Vf: func(mi EnvInjector) (any, error) {
				if errInParse != nil {
					return nil, errInParse
				}
				leftValue, err := leftValueEntry.Vf(mi)
				if err != nil {
					return nil, err
				}
				leftValueAsMap, ok := leftValue.(map[string]any)
				if !ok {
					return nil, errors.Errorf("want got a map, but got %T", leftValue)
				}
				rightValue, err := l.makeValueEntry(express[1]).Vf(mi)
				if err != nil {
					return nil, err
				}
				if rightValueAsString == "" {
					rightValueAsString, ok = rightValue.(string)
					if !ok {
						return nil, errors.Errorf("want got a string, but got %T", rightValue)
					}
				}
				v, ok := leftValueAsMap[rightValueAsString]
				if !ok {
					return nil, errors.Errorf("key %s not found", rightValueAsString)
				}
				return v, nil
			},
		}
		res.SourceDepend.NeedKeys = append(res.SourceDepend.NeedKeys, leftValueEntry.SourceDepend.NeedKeys...)
		res.SourceDepend.NeedKeys = append(res.SourceDepend.NeedKeys, rightValueAsString)
		return res
	case ctx.OPEN_BRA() != nil:
		return &ValueEntry{
			Vf: func(mi EnvInjector) (any, error) {
				express := ctx.AllValueExpr()
				if len(express) != 2 {
					return nil, errors.Errorf("len(exprs) %d != 2", len(express))
				}
				leftValue, err := l.makeValueEntry(express[0]).Vf(mi)
				if err != nil {
					return nil, err
				}
				leftValueAsArr, ok := leftValue.([]any)
				if !ok {
					return nil, errors.Errorf("want got a array, but got %T", leftValue)
				}
				rightValue, err := l.makeValueEntry(express[1]).Vf(mi)
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
			},
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
		return &ValueEntry{
			Vf: func(mi EnvInjector) (any, error) {
				if len(args) != 1 {
					return nil, errors.Errorf("len(args) %d != 1", len(args))
				}
				var obj any
				value, err := l.makeValueEntry(args[0]).Vf(mi)
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
			},
		}
	case ctx.OPEN_PAR() != nil:
		return l.makeValueEntry(ctx.ValueExpr(0))
	}
	return nil
}
